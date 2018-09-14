package oauth2

import (
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

type testVerifier struct {
	token string
}

func (t *testVerifier) VerifyToken(ctx context.Context, token string) (*TokenInfo, error) {
	t.token = token
	return nil, nil
}

func TestAuthenticator(t *testing.T) {
	tt := []struct {
		prefix string
		bearer string
		token  string
		output error
	}{
		{BearerPrefix, "token", "", errUnauthenticated},
		{BearerPrefix, "Bearer t", "", errUnauthenticated},
		{BearerPrefix, "Bearertoken", "", errUnauthenticated},
		{BearerPrefix, "Bearer long-token with space", "", errUnauthenticated},
		{BearerPrefix, "Bearer token", "token", nil},
		{BearerPrefix, "bearer token", "token", nil},
		{BearerPrefix, "BEARER token", "token", nil},
		{BearerPrefix, "Bearer long-token", "long-token", nil},
		{BearerPrefix, "bearer token", "token", nil},
		// Other prefix
		{APIKeyPrefix, "token", "", errUnauthenticated},
		{APIKeyPrefix, "ApiKey t", "", errUnauthenticated},
		{APIKeyPrefix, "ApiKeytoken", "", errUnauthenticated},
		{APIKeyPrefix, "apikey token", "token", nil},
		{APIKeyPrefix, "apiKey token", "token", nil},
		{APIKeyPrefix, "ApiKey token", "token", nil},
		{APIKeyPrefix, "ApiKey long-token", "long-token", nil},
		{APIKeyPrefix, "apikey token", "token", nil},
	}
	auth := Authenticator{}
	for _, tc := range tt {
		v := &testVerifier{}
		auth.verifiers = VerifierMap{tc.prefix: v}
		auth.encodeFunc = EncodeJWT([]byte("my-secret"))
		md := metadata.Pairs(authKey, tc.bearer)
		ctx := metadata.NewIncomingContext(context.Background(), md)
		if _, err := auth.Authenticate(ctx); err != tc.output {
			t.Errorf("auth.Authenticate(ctx(%s)) = %v; expected %v", tc.bearer, err, tc.output)
		}
		if v.token != tc.token {
			t.Errorf("Verifier got token '%s'; expected '%s'", v.token, tc.token)
		}
	}
}
