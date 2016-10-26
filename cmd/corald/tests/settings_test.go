package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ardanlabs/kit/tests"
	"github.com/coralproject/shelf/internal/sponge/item"
)

// TestSettingsUpsert tests if we can upsert settings.
func TestSetingsUpsert(t *testing.T) {
	tests.ResetLog()
	defer tests.DisplayLog()

	t.Log("Given the need to upsert settings.")
	{
		settings := item.Item{
			Type:    "settings",
			Version: 1,
			Data: map[string]interface{}{
				"moderation_type": "pre",
			},
		}

		payload, err := json.Marshal(settings)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to marshal the settings payload : %v", tests.Failed, err)
		}

		url := "/v1/settings"
		r := httptest.NewRequest("PUT", url, bytes.NewBuffer(payload))
		w := httptest.NewRecorder()

		a.ServeHTTP(w, r)

		t.Logf("\tWhen calling url : %s", url)
		{
			t.Log("\tWhen we use version v1 of the settings endpoint.")

			if w.Code != http.StatusOK {
				t.Fatalf("\t%s\tShould be able to upsert the settings item : %v", tests.Failed, w.Code)
			}
			t.Logf("\t%s\tShould be able to upsert the settings item.", tests.Success)
		}
	}
}

// TestSettingsList tests if we can list current settings.
func TestSettingsList(t *testing.T) {
	tests.ResetLog()
	defer tests.DisplayLog()

	t.Log("Given the need to list settings.")
	{
		url := "/v1/settings"
		r := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()

		a.ServeHTTP(w, r)

		t.Logf("\tWhen calling url : %s", url)
		{
			t.Log("\tWhen we use version v1 of the settings endpoint.")

			if w.Code != http.StatusOK {
				t.Fatalf("\t%s\tShould be able to list the settings : %v", tests.Failed, w.Code)
			}
			t.Logf("\t%s\tShould be able to list the settings.", tests.Success)
		}
	}
}
