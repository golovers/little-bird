package oauth2

import (
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// Token prefixes
const (
	BearerPrefix = "bearer"
	APIKeyPrefix = "apikey"
)

var grpcErr = grpc.Errorf
var errUnauthenticated = grpcErr(codes.Unauthenticated, "Unauthenticated")
var authKey = "authorization"
var minTokenLength = 5

// JWTEncodeFunc used to encode new JWT tokens
type JWTEncodeFunc func(c jwt.Claims) ([]byte, error)

// EncodeJWT returns a `oauth2.JWTEncodeFunc` to encode a struct
// into as the payload to a JWT, signed with `secret`.
func EncodeJWT(secret []byte) func(jwt.Claims) ([]byte, error) {
	return func(c jwt.Claims) ([]byte, error) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
		ss, err := token.SignedString(secret)
		return []byte(ss), err
	}
}

// Verifier implementations should verify bearer tokens and
// return a JWT token to be used for core-services.
type Verifier interface {
	VerifyToken(ctx context.Context, token string) (*TokenInfo, error)
}

// VerifierMap represents a map of the token prefix with the matching Verifier
type VerifierMap map[string]Verifier

// Authenticator implements the `auth.Authenticator` interface,
// and will validate each incoming request for a valid JWT token.
type Authenticator struct {
	encodeFunc JWTEncodeFunc
	verifiers  VerifierMap
}

// NewAuthenticator returns a new OAuth2 token introspector implementing
// `auth.Authenticator`.
func NewAuthenticator(verifiers VerifierMap, jwtSecret []byte) *Authenticator {
	return &Authenticator{
		verifiers:  verifiers,
		encodeFunc: EncodeJWT(jwtSecret),
	}
}

// Authenticate validates the OAuth2 Access Token with the Authorization Server,
// and replace the token with a signed JWT. Implements `auth.Authenticator.Authenticate`.
// Returns an error with the `codes.Unauthenticated` GRPC error codes if
// validation fails.
func (a *Authenticator) Authenticate(ctx context.Context) (context.Context, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if len(md[authKey]) > 0 {
			if fields := strings.Fields(md[authKey][0]); len(fields) == 2 {
				for prefix, verifier := range a.verifiers {
					// Check if the bearer prefix is present.
					if strings.EqualFold(prefix, fields[0]) {
						token := string(fields[1])
						if len(token) >= minTokenLength {
							// VerifyToken will look the token up using a network connection or
							// internal cache.
							info, err := verifier.VerifyToken(ctx, token)
							if err == nil {
								// Create a new JWT from the TokenInfo and add it to the metadata.
								// Scope + Userinfo + username
								jwt, err := a.encodeFunc(info)
								if err != nil {
									return nil, err
								}
								ctx = newJWTContext(ctx, string(jwt))
								ctx = newContext(ctx, info)
								return ctx, nil
							}
						}
					}
				}
			}
		}
	}
	return ctx, errUnauthenticated
}
