package gateway

import (
	"context"
	"net/http"

	"github.com/tencentyun/scf-go-lib/cloudevents/scf"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

func Serve(h http.Handler) error {
	if h == nil {
		h = http.DefaultServeMux
	}

	cloudfunction.Start(func(ctx context.Context, e scf.APIGatewayProxyRequest) (scf.APIGatewayProxyResponse, error) {
		r, err := NewRequest(ctx, e)
		if err != nil {
			return scf.APIGatewayProxyResponse{}, err
		}

		w := NewResponse()
		h.ServeHTTP(w, r)
		return w.End(), nil
	})

	return nil
}
