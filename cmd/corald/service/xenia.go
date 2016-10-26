package service

import (
	"fmt"
	"net/http"

	"github.com/ardanlabs/kit/web"
)

// ExecQueryOnView executes a custom xenia query on a view and returns the results.
func ExecQueryOnView(c *web.Context, viewName string, itemKey string, query []byte) (*http.Response, error) {

	// Get the xenia URL.
	xeniadURL, ok := c.Web.Ctx["xeniadURL"].(string)
	if !ok {
		return nil, fmt.Errorf("No valid xeniad URL defined")
	}

	// Define the URL for the request.
	url := xeniadURL + "/v1/exec/view/" + viewName + "/" + itemKey

	// Get the query results.
	return requestService(c, http.MethodPost, url, query)
}
