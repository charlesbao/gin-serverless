package gateway

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/tencentyun/scf-go-lib/cloudevents/scf"
)

// NewRequest returns a new http.Request from the given Lambda event.
func NewRequest(ctx context.Context, e scf.APIGatewayProxyRequest) (*http.Request, error) {
	// path
	u, err := url.Parse(e.Path)
	if err != nil {
		return nil, errors.Wrap(err, "parsing path")
	}

	// querystring
	q := u.Query()
	for k, v := range e.QueryString {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()

	// base64 encoded body
	body := e.Body
	if e.IsBase64Encoded {
		b, err := base64.StdEncoding.DecodeString(body)
		if err != nil {
			return nil, errors.Wrap(err, "decoding base64 body")
		}
		body = string(b)
	}

	// new request
	req, err := http.NewRequest(e.HTTPMethod, u.String(), strings.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}

	// remote addr
	req.RequestURI = e.Path
	req.RemoteAddr = e.RequestContext.SourceIP
	req.Header.Set("X-Forwarded-For", req.RemoteAddr)
	req.Header.Set("X-Real-Ip", req.RemoteAddr)

	// header fields
	for k, v := range e.Headers {
		req.Header.Set(k, v)
	}

	// content-length
	if req.Header.Get("content-length") == "" && body != "" {
		req.Header.Set("content-length", strconv.Itoa(len(body)))
	}

	// custom context values
	req = req.WithContext(newContext(ctx, e))

	// host
	req.URL.Host = req.Header.Get("host")
	req.Host = req.URL.Host

	return req, nil
}
