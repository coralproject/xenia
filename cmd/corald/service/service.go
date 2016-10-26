package service

import (
	"bytes"
	"net/http"

	"github.com/ardanlabs/kit/web"
	"github.com/coralproject/shelf/internal/platform/auth"
)

// SignServiceRequest signs a request with the claims necessary to authenticate
// with downstream services.
func SignServiceRequest(context interface{}, signer auth.Signer, r *http.Request) error {
	claims := map[string]interface{}{}

	return auth.SignRequest(context, signer, claims, r)
}

// Rewrite will add service request headers to the request and add other
// standards.
func Rewrite(c *web.Context) func(*http.Request) {

	f := func(r *http.Request) {

		// Extract the signer from the application context.
		signer, ok := c.Web.Ctx["signer"].(auth.Signer)
		if !ok {
			return
		}

		// Sign the service request with the signer.
		if err := SignServiceRequest(c.SessionID, signer, r); err != nil {
			return
		}

	}

	return f
}

// RewritePath will rewrite the path given a PathRewriter and return the request
// director function.
func RewritePath(c *web.Context, targetPath string) func(*http.Request) {

	f := func(r *http.Request) {

		// Rewrite the request for the services which will add authentication
		// headers and/or other default headers.
		Rewrite(c)(r)

		// Update the target path.
		r.URL.Path = targetPath
	}

	return f
}

func requestService(c *web.Context, verb string, url string, payload []byte) (*http.Response, error) {
	req, err := http.NewRequest(verb, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	// Extract the signer from the application context.
	if signer, ok := c.Web.Ctx["signer"].(auth.Signer); ok {
		// Sign the service request with the signer.
		if err = SignServiceRequest(c.SessionID, signer, req); err != nil {
			return nil, err
		}
	}

	// Get the response.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
