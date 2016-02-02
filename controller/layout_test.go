package controller

import (
	"appengine/aetest"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestNewLayoutAPI(t *testing.T) {

	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	jsonParams := new(bytes.Buffer)
	json.NewEncoder(jsonParams).Encode(map[string]string{"name": "My Layout"})
	t.Logf("JSON param data: %v", jsonParams.String())

	apiReq, err := inst.NewRequest("POST", "/api/newLayout", jsonParams)
	if err != nil {
		t.Fatalf("Failed to create apiReq: %v", err)
	}
	record := httptest.NewRecorder()

	t.Logf("apiReq Body: %v", apiReq.Body)

	newLayout(record, apiReq)

	t.Logf("api response: %v", record.Code)

	if record.Code != 200 {
		t.Errorf("Unexpected result from newLayout API call: %v", record.Code)
	}

}
