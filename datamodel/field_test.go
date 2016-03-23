package datamodel

import (
	"appengine"
	"testing"
)

// Helper functions for testing fields

func newTestNumField(appEngContext appengine.Context, t *testing.T, refName string) string {

	testNumField := NewFieldParams{Name: refName, Type: FieldTypeNumber, RefName: refName}
	testNumFieldID, numFieldErr := NewField(appEngContext, testNumField)
	if numFieldErr != nil {
		t.Fatal(numFieldErr)
	}

	return testNumFieldID
}

func newTestTextField(appEngContext appengine.Context, t *testing.T, refName string) string {

	testField := NewFieldParams{Name: "Test Text Field", Type: FieldTypeText, RefName: "TestRef1"}
	testFieldID, err := NewField(appEngContext, testField)
	if err != nil {
		t.Fatal(err)
	}

	return testFieldID
}
