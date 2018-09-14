package oauth2

import "golang.org/x/net/context"

// Keys to use with context.Context
type oauth2Key struct{}
type oauth2JWTKey struct{}

// newContext creates a new context with `info` attached.
func newContext(ctx context.Context, info *TokenInfo) context.Context {
	return context.WithValue(ctx, oauth2Key{}, info)
}

// FromContext returns the TokenInfo from ctx if it exists.
func FromContext(ctx context.Context) (info *TokenInfo, ok bool) {
	info, ok = ctx.Value(oauth2Key{}).(*TokenInfo)
	return
}

// newJWTContext creates a new context with the raw JWT populated.
func newJWTContext(ctx context.Context, jwt string) context.Context {
	return context.WithValue(ctx, oauth2JWTKey{}, jwt)
}

// JWTFromContext returns the raw JWT from ctx if it exists.
func JWTFromContext(ctx context.Context) (jwt string, ok bool) {
	jwt, ok = ctx.Value(oauth2JWTKey{}).(string)
	return
}
