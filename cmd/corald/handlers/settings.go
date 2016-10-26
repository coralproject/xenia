package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ardanlabs/kit/web"
	"github.com/coralproject/shelf/cmd/corald/service"
	"github.com/coralproject/shelf/internal/sponge/item"
	"github.com/coralproject/shelf/internal/xenia/query"
)

// settingsHandle maintains the set of handlers for the settings api.
type settingsHandle struct{}

// Settings fronts the access to the settings functionality.
var Settings settingsHandle

// allQuery is a xenia query to return all view documents.
const allQuery = `{
    "name": "all",
    "params": [],
    "queries": [
      {
        "name": "all",
        "type": "pipeline",
        "collection": "view",
        "commands": [
            {"$skip": 0}
        ],
        "indexes": [],
        "return": true
      }
    ],
    "enabled": true,
    "explain": false
}`

const (

	// settingsView is the view name that is used to return settings items.
	settingsView = "settings"

	// settingsKey is the item key used to return settings items.
	settingsKey = "settings"
)

//==============================================================================

// Upsert a settings item.
// 200 Success, 404 Not Found, 500 Internal.
func (settingsHandle) Upsert(c *web.Context) error {

	// Decode the item.
	var itm item.Item
	if err := json.NewDecoder(c.Request.Body).Decode(&itm); err != nil {
		return err
	}

	// Add the indication that this item is a settings item.
	itm.Data["has_type"] = "settings"

	// Upsert the item.
	if err := service.UpsertItem(c, &itm); err != nil {
		return err
	}

	c.Respond(nil, http.StatusOK)
	return nil
}

// List returns the current list of settings items.
// 200 Success, 404 Not Found, 500 Internal.
func (settingsHandle) List(c *web.Context) error {

	// Prepare the query.
	payload := []byte(allQuery)

	// Query the view.
	resp, err := service.ExecQueryOnView(c, settingsView, settingsKey, payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response Body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Unmarshal the results.
	var results query.Result
	if err := json.Unmarshal(body, &results); err != nil {
		return err
	}

	c.Respond(results, http.StatusOK)
	return nil
}
