package oauth2

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"golang.org/x/net/context"
)

type testHandler struct {
	receivedToken string
	response      interface{}
	status        int
}

func (t *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.receivedToken = r.FormValue("token")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(t.status)
	json.NewEncoder(w).Encode(t.response)
}

func TestVerifier(t *testing.T) {
	handler := &testHandler{}
	ts := httptest.NewServer(handler)
	defer ts.Close()

	tt := []struct {
		token     string
		status    int
		tokenInfo *TokenInfo
		err       error
	}{
		{"token", 200, &TokenInfo{Active: false}, ErrTokenValidationFailed},
		{"token", 501, &TokenInfo{Active: true}, ErrTokenValidationFailed},
		{"token", 200, &TokenInfo{Active: true, Scope: "openid"}, nil},
	}
	verifier := NewVerifier(ts.URL)
	ctx := context.Background()

	for _, tc := range tt {
		handler.response = tc.tokenInfo
		handler.status = tc.status
		ti, err := verifier.VerifyToken(ctx, tc.token)
		if handler.receivedToken != tc.token {
			t.Errorf("Handler called with token %s; expected %s", handler.receivedToken, tc.token)
		}
		if err != tc.err {
			t.Errorf("VerifyToken(ctx, %s) = _, %v; expected %v", tc.token, err, tc.err)
		}
		if err == nil {
			if !reflect.DeepEqual(ti, tc.tokenInfo) {
				t.Errorf("VerifyToken(ctx, %s) = %v; expected %v", tc.token, ti, tc.tokenInfo)
			}
		}
	}
}
