package jwt

import (
	"errors"

	"golang.org/x/net/context"

	"gitlab.com/7chip/little-bird/backend/micro/auth/oauth2"
	"google.golang.org/grpc/credentials"
)

var errUnauthenticated = errors.New("jwt: token missing from context")

// NewJWTCredentialsFromToken returns a grpc rpc credential
// using the provided JWT token. Does not validate the Token.
func NewJWTCredentialsFromToken(token string) credentials.PerRPCCredentials {
	return jwtToken(token)
}

// NewJWTCredentials returns a grpc credential that extracts
// the JWT token from the callers context object.
func NewJWTCredentials() credentials.PerRPCCredentials {
	return jwtAuthenticator{}
}

type jwtToken string

// GetRequestMetadata implements the `credentials.PerRPCCredentials`
// method `GetRequestMetadata`, by returning a simple map to be appended
// to GRPCs metadata.
func (j jwtToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": string(j),
	}, nil
}

// RequireTransportSecurity implements `credentials.PerRPCCredentials`
// method `RequireTransportSecurity`. Indicates if we want a secure
// transport.
func (j jwtToken) RequireTransportSecurity() bool {
	return false
}

type jwtAuthenticator struct{}

// GetRequestMetadata implements the `credentials.PerRPCCredentials`
// method `GetRequestMetadata`, by returning a simple map to be appended
// to GRPCs metadata.
func (j jwtAuthenticator) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	jwt, ok := oauth2.JWTFromContext(ctx)
	if !ok {
		return nil, errUnauthenticated
	}
	return map[string]string{
		"authorization": jwt,
	}, nil
}

// RequireTransportSecurity implements `credentials.PerRPCCredentials`
// method `RequireTransportSecurity`. Indicates if we want a secure
// transport.
func (j jwtAuthenticator) RequireTransportSecurity() bool {
	return false
}
