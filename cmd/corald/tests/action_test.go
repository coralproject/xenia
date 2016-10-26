package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ardanlabs/kit/tests"
)

// TestActionsPOST sample test for the POST call to add actions to items.
func TestActionsPOST(t *testing.T) {
	tests.ResetLog()
	defer tests.DisplayLog()

	t.Log("Given the need to add an action to an item.")
	{
		action := "flagged_by"
		userID := "ITEST_80aa936a-f618-4234-a7be-df59a14cf8de"
		itemID := "ITEST_d16790f8-13e9-4cb4-b9ef-d82835589660"
		url := fmt.Sprintf("/v1/action/%s/user/%s/on/item/%s", action, userID, itemID)
		r := httptest.NewRequest("POST", url, nil)
		w := httptest.NewRecorder()

		a.ServeHTTP(w, r)

		t.Logf("\tWhen calling url : %s", url)
		{
			t.Log("\tWhen we use version v1 of the actions endpoint.")

			if w.Code != http.StatusOK {
				t.Fatalf("\t%s\tShould be able to add the action : %v", tests.Failed, w.Code)
			}
			t.Logf("\t%s\tShould be able to add the action.", tests.Success)
		}
	}

	t.Log("Given the need to catch an error when adding an action to an item.")
	{
		action := "flagged_by"
		userID := "ITEST_80aa936a-f618-4234-a7be-df59a14cf8de"
		itemID := "wrongitem"
		url := fmt.Sprintf("/v1/action/%s/user/%s/on/item/%s", action, userID, itemID)
		r := httptest.NewRequest("POST", url, nil)
		w := httptest.NewRecorder()

		a.ServeHTTP(w, r)

		t.Logf("\tWhen calling url : %s", url)
		{
			t.Log("\tWhen we use version v1 of the actions endpoint.")

			if w.Code != http.StatusInternalServerError {
				t.Fatalf("\t%s\tShould fail on finding the target : %v", tests.Failed, w.Code)
			}
			t.Logf("\t%s\tShould fail on finding the target.", tests.Success)
		}
	}

}

// TestActionsDELETE sample test for the DELETE call to remove actions from items.
func TestActionsDELETE(t *testing.T) {
	tests.ResetLog()
	defer tests.DisplayLog()

	t.Log("Given the need to remove actions from items.")
	{
		action := "flagged_by"
		userID := "ITEST_a63af637-58af-472b-98c7-f5c00743bac6"
		itemID := "ITEST_d1dfa366-d2f7-4a4a-a64f-af89d4c97d82"

		url := fmt.Sprintf("/v1/action/%s/user/%s/on/item/%s", action, userID, itemID)
		r := httptest.NewRequest("DELETE", url, nil)
		w := httptest.NewRecorder()

		a.ServeHTTP(w, r)

		t.Logf("\tWhen calling url : %s", url)
		{
			t.Log("\tWhen we use version v1 of the actions endpoint.")

			if w.Code != http.StatusOK {
				t.Fatalf("\t%s\tShould be able to remove the action : %v", tests.Failed, w.Code)
			}
			t.Logf("\t%s\tShould be able to remove the action.", tests.Success)
		}
	}
}
