package oauth2

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
)

// Errors
var (
	ErrTokenValidationFailed = errors.New("oauth2: token validation failed")
)

// TokenInfo is the current status of the token.
type TokenInfo struct {
	Active    bool   `json:"active,omitempty"`
	Scope     string `json:"scope,omitempty"`
	Username  string `json:"username,omitempty"`
	ClientID  string `json:"client_id,omitempty"`
	TokenType string `json:"token_type,omitempty"`

	// Once we have service-accounts in place, this should be removed.
	// Its up to each service to decide how they would like to handle
	// admin-callers.
	Admin bool `json:"admin,omitempty"`

	jwt.StandardClaims
}

// Options to configure the Verifier
type verifierOptions struct {
	tokenTypeHint string

	clientAuth   bool
	clientID     string
	clientSecret string
}

// VerifierOption configures how Verifier is used.
type VerifierOption func(*verifierOptions)

// WithTokenTypeHint will indicate the token type to the token introspection
// endpoint.
func WithTokenTypeHint(hint string) VerifierOption {
	return func(o *verifierOptions) {
		o.tokenTypeHint = hint
	}
}

// WithClientCredentials sets a client ID and client secret to use while
// verifying tokens.
func WithClientCredentials(id, secret string) VerifierOption {
	return func(o *verifierOptions) {
		o.clientAuth = true
		o.clientID = id
		o.clientSecret = secret
	}
}

// TokenVerifier implements Verifier and perform token validation
// against an OAuth2 Token Introspection endpoint using HTTP.
type TokenVerifier struct {
	client *http.Client
	url    string

	options verifierOptions
}

// NewVerifier creates a new Verifier for the provided
// token introspection URL.
func NewVerifier(introspectionURL string, opts ...VerifierOption) *TokenVerifier {
	opt := verifierOptions{
		tokenTypeHint: "",
	}
	for _, o := range opts {
		o(&opt)
	}
	v := &TokenVerifier{
		url:     introspectionURL,
		client:  &http.Client{},
		options: opt,
	}
	return v
}

// VerifyToken implements Verifier, and checks with a OAuth2 Token Introspection
// Endpoint if the token is valid.
func (t *TokenVerifier) VerifyToken(ctx context.Context, token string) (*TokenInfo, error) {
	data := url.Values{}
	data.Set("token", token)
	data.Set("token_type_hint", t.options.tokenTypeHint)
	req, err := http.NewRequest("POST", t.url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if t.options.clientAuth {
		req.SetBasicAuth(t.options.clientID, t.options.clientSecret)
	}
	resp, err := ctxhttp.Do(ctx, t.client, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, ErrTokenValidationFailed
	}
	var tokenInfo TokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		return nil, err
	}
	if !tokenInfo.Active {
		return nil, ErrTokenValidationFailed
	}
	return &tokenInfo, nil
}
