// Package endpoint implements users tests for the API layer.
package tests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/coralproject/xenia/pkg/script"
	"github.com/coralproject/xenia/pkg/script/sfix"

	"github.com/ardanlabs/kit/tests"
)

// TestScripts tests the retrieval of scripts.
func TestScripts(t *testing.T) {
	tests.ResetLog()
	defer tests.DisplayLog()

	t.Log("Given the need get a set of scripts.")
	{
		url := "/1.0/script"
		r := tests.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()

		a.ServeHTTP(w, r)

		t.Logf("\tWhen calling url : %s", url)
		{
			t.Log("\tWhen we user version 1.0 of the script endpoint.")
			if w.Code != 200 {
				t.Fatalf("\t%s\tShould be able to retrieve the script list : %v", tests.Failed, w.Code)
			}
			t.Logf("\t%s\tShould be able to retrieve the script list.", tests.Success)

			var scrs []script.Script
			if err := json.Unmarshal(w.Body.Bytes(), &scrs); err != nil {
				t.Fatalf("\t%s\tShould be able to unmarshal the results : %v", tests.Failed, err)
			}
			t.Logf("\t%s\tShould be able to unmarshal the results.", tests.Success)

			var count int
			for _, scr := range scrs {
				if scr.Name[0:5] == "STEST" {
					count++
				}
			}

			if count != 2 {
				t.Fatalf("\t%s\tShould have two scripts : %d", tests.Failed, count)
			}
			t.Logf("\t%s\tShould have two scripts.", tests.Success)
		}
	}
}

// TestScriptByName tests the retrieval of a specific script.
func TestScriptByName(t *testing.T) {
	tests.ResetLog()
	defer tests.DisplayLog()

	t.Log("Given the need to get a specific script.")
	{
		url := "/1.0/script/STEST_basic_script_pre"
		r := tests.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()

		a.ServeHTTP(w, r)

		t.Logf("\tWhen calling url : %s", url)
		{
			if w.Code != 200 {
				t.Fatalf("\t%s\tShould be able to retrieve the script : %v", tests.Failed, w.Code)
			}
			t.Logf("\t%s\tShould be able to retrieve the script.", tests.Success)

			var scr script.Script
			if err := json.Unmarshal(w.Body.Bytes(), &scr); err != nil {
				t.Fatalf("\t%s\tShould be able to unmarshal the results : %v", tests.Failed, err)
			}
			t.Logf("\t%s\tShould be able to unmarshal the results.", tests.Success)

			if scr.Name != "STEST_basic_script_pre" {
				t.Fatalf("\t%s\tShould have the correct script : %s", tests.Failed, scr.Name)
			}
			t.Logf("\t%s\tShould have the correct script.", tests.Success)
		}
	}
}

// TestScriptUpsert tests the insert and update of a script.
func TestScriptUpsert(t *testing.T) {
	tests.ResetLog()
	defer tests.DisplayLog()

	t.Log("Given the need to insert and then update a script.")
	{
		//----------------------------------------------------------------------
		// Get the fixture.

		scr, err := sfix.Get("upsert.json")
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve the fixture : %v", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve the fixture.", tests.Success)

		scrStrData, err := json.Marshal(&scr)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to marshal the fixture : %v", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to marshal the fixture.", tests.Success)

		//----------------------------------------------------------------------
		// Insert the Script.

		url := "/1.0/script"
		r := tests.NewRequest("PUT", url, bytes.NewBuffer(scrStrData))
		w := httptest.NewRecorder()

		a.ServeHTTP(w, r)

		t.Logf("\tWhen calling url to insert : %s", url)
		{
			if w.Code != 204 {
				t.Fatalf("\t%s\tShould be able to insert the script : %v", tests.Failed, w.Code)
			}
			t.Logf("\t%s\tShould be able to insert the script.", tests.Success)
		}

		//----------------------------------------------------------------------
		// Retrieve the Script.

		url = "/1.0/script/STEST_upsert"
		r = tests.NewRequest("GET", url, nil)
		w = httptest.NewRecorder()

		a.ServeHTTP(w, r)

		t.Logf("\tWhen calling url to get : %s", url)
		{
			if w.Code != 200 {
				t.Fatalf("\t%s\tShould be able to retrieve the script : %v", tests.Failed, w.Code)
			}
			t.Logf("\t%s\tShould be able to retrieve the script.", tests.Success)

			recv := w.Body.String()
			resp := `{"name":"STEST_upsert","commands":[{"command.one":1},{"command":2},{"command":3}]}`

			if resp != recv {
				t.Log(resp)
				t.Log(w.Body.String())
				t.Fatalf("\t%s\tShould get the expected result.", tests.Failed)
			}
			t.Logf("\t%s\tShould get the expected result.", tests.Success)
		}

		//----------------------------------------------------------------------
		// Update the Script.

		scr.Commands = append(scr.Commands, map[string]interface{}{"command": 4})

		scrStrData, err = json.Marshal(&scr)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to marshal the changed fixture : %v", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to marshal the changed fixture.", tests.Success)

		url = "/1.0/script"
		r = tests.NewRequest("PUT", url, bytes.NewBuffer(scrStrData))
		w = httptest.NewRecorder()

		a.ServeHTTP(w, r)

		t.Logf("\tWhen calling url to update : %s", url)
		{
			if w.Code != 204 {
				t.Fatalf("\t%s\tShould be able to update the script : %v", tests.Failed, w.Code)
			}
			t.Logf("\t%s\tShould be able to update the script.", tests.Success)
		}

		//----------------------------------------------------------------------
		// Retrieve the Script.

		url = "/1.0/script/STEST_upsert"
		r = tests.NewRequest("GET", url, nil)
		w = httptest.NewRecorder()

		a.ServeHTTP(w, r)

		t.Logf("\tWhen calling url to get : %s", url)
		{
			if w.Code != 200 {
				t.Fatalf("\t%s\tShould be able to retrieve the script : %v", tests.Failed, w.Code)
			}
			t.Logf("\t%s\tShould be able to retrieve the script.", tests.Success)

			recv := w.Body.String()
			resp := `{"name":"STEST_upsert","commands":[{"command.one":1},{"command":2},{"command":3},{"command":4}]}`

			if resp != recv {
				t.Log(resp)
				t.Log(w.Body.String())
				t.Fatalf("\t%s\tShould get the expected result.", tests.Failed)
			}
			t.Logf("\t%s\tShould get the expected result.", tests.Success)
		}
	}
}

// TestScriptDelete tests the insert and deletion of a script.
func TestScriptDelete(t *testing.T) {
	tests.ResetLog()
	defer tests.DisplayLog()

	t.Log("Given the need to insert and then delete a script.")
	{
		//----------------------------------------------------------------------
		// Get the fixture.

		scr, err := sfix.Get("upsert.json")
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve the fixture : %v", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve the fixture.", tests.Success)

		scrStrData, err := json.Marshal(&scr)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to marshal the fixture : %v", tests.Failed, err)
		}
		t.Logf("\t%s\tShould be able to marshal the fixture.", tests.Success)

		//----------------------------------------------------------------------
		// Insert the Script.

		url := "/1.0/script"
		r := tests.NewRequest("PUT", url, bytes.NewBuffer(scrStrData))
		w := httptest.NewRecorder()

		a.ServeHTTP(w, r)

		t.Logf("\tWhen calling url to insert : %s", url)
		{
			if w.Code != 204 {
				t.Fatalf("\t%s\tShould be able to insert the script : %v", tests.Failed, w.Code)
			}
			t.Logf("\t%s\tShould be able to insert the script.", tests.Success)
		}

		//----------------------------------------------------------------------
		// Retrieve the Script.

		url = "/1.0/script/STEST_upsert"
		r = tests.NewRequest("GET", url, nil)
		w = httptest.NewRecorder()

		a.ServeHTTP(w, r)

		t.Logf("\tWhen calling url to get : %s", url)
		{
			if w.Code != 200 {
				t.Fatalf("\t%s\tShould be able to retrieve the script : %v", tests.Failed, w.Code)
			}
			t.Logf("\t%s\tShould be able to retrieve the script.", tests.Success)

			recv := w.Body.String()
			resp := `{"name":"STEST_upsert","commands":[{"command.one":1},{"command":2},{"command":3}]}`

			if resp != recv {
				t.Log(resp)
				t.Log(w.Body.String())
				t.Fatalf("\t%s\tShould get the expected result.", tests.Failed)
			}
			t.Logf("\t%s\tShould get the expected result.", tests.Success)
		}

		//----------------------------------------------------------------------
		// Delete the Script.

		url = "/1.0/script/STEST_upsert"
		r = tests.NewRequest("DELETE", url, nil)
		w = httptest.NewRecorder()

		a.ServeHTTP(w, r)

		t.Logf("\tWhen calling url to delete : %s", url)
		{
			if w.Code != 204 {
				t.Fatalf("\t%s\tShould be able to delete the script : %v", tests.Failed, w.Code)
			}
			t.Logf("\t%s\tShould be able to delete the script.", tests.Success)
		}

		//----------------------------------------------------------------------
		// Retrieve the Script.

		url = "/1.0/script/STEST_upsert"
		r = tests.NewRequest("GET", url, nil)
		w = httptest.NewRecorder()

		a.ServeHTTP(w, r)

		t.Logf("\tWhen calling url to get : %s", url)
		{
			if w.Code != 404 {
				t.Fatalf("\t%s\tShould Not be able to retrieve the script : %v", tests.Failed, w.Code)
			}
			t.Logf("\t%s\tShould Not be able to retrieve the script.", tests.Success)
		}
	}
}
