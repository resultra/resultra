package testUtil

import (
	"encoding/json"
	"testing"
)

func EncodeJSONString(t *testing.T, val interface{}) string {
	b, err := json.Marshal(val)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
