package gateway

import (
	"context"

	"github.com/tencentyun/scf-go-lib/cloudevents/scf"
)

// key is the type used for any items added to the request context.
type key int

// requestContextKey is the key for the api gateway proxy `RequestContext`.
const requestContextKey key = iota

// newContext returns a new Context with specific api gateway proxy values.
func newContext(ctx context.Context, e scf.APIGatewayProxyRequest) context.Context {
	return context.WithValue(ctx, requestContextKey, e.RequestContext)
}

// RequestContext returns the APIGatewayProxyRequestContext value stored in ctx.
func RequestContext(ctx context.Context) (scf.APIGatewayProxyRequestContext, bool) {
	c, ok := ctx.Value(requestContextKey).(scf.APIGatewayProxyRequestContext)
	return c, ok
}
