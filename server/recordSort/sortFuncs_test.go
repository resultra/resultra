// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package recordSort

import (
	"github.com/resultra/resultra/server/common/recordSortDataModel"
	"github.com/resultra/resultra/server/common/testUtil"
	"github.com/resultra/resultra/server/record"
	"github.com/resultra/resultra/server/recordValue"
	"testing"
	"time"
)

func TestNumberValueSort(t *testing.T) {

	fieldID := "fieldA"

	rec1Values := record.RecFieldValues{}
	rec1Values[fieldID] = 42.5
	recVal1 := recordValue.RecordValueResults{
		RecordID:    "Record_01",
		FieldValues: rec1Values}

	rec2Values := record.RecFieldValues{}
	rec2Values[fieldID] = 100.0
	recVal2 := recordValue.RecordValueResults{
		RecordID:    "Record_02",
		FieldValues: rec2Values}

	// recValues initialized with records out of order
	recValues := []recordValue.RecordValueResults{recVal2, recVal1}

	fieldAsc := SortByNumberField(fieldID, recordSortDataModel.SortDirectionAsc)
	OrderedBy(fieldAsc).Sort(recValues)
	if recValues[0].RecordID != "Record_01" {
		t.Errorf("TestRecordValueSort: expecting Record_01 first")
	}
	t.Logf("TestRecordValueSort: sort asc results %+v", testUtil.EncodeJSONString(t, recValues))

	fieldDesc := SortByNumberField(fieldID, recordSortDataModel.SortDirectionDesc)
	OrderedBy(fieldDesc).Sort(recValues)
	if recValues[0].RecordID != "Record_02" {
		t.Errorf("TestRecordValueSort: expecting Record_02 first")
	}
	t.Logf("TestRecordValueSort: sort desc results %+v", testUtil.EncodeJSONString(t, recValues))

}

func TestBlankNumberValueSort(t *testing.T) {

	fieldID := "fieldA"

	rec1Values := record.RecFieldValues{}
	// field value not set.
	recVal1 := recordValue.RecordValueResults{
		RecordID:    "Record_01",
		FieldValues: rec1Values}

	rec2Values := record.RecFieldValues{}
	rec2Values[fieldID] = 100.0
	recVal2 := recordValue.RecordValueResults{
		RecordID:    "Record_02",
		FieldValues: rec2Values}

	// recValues initialized with records out of order
	recValues := []recordValue.RecordValueResults{recVal1, recVal2}

	fieldAsc := SortByNumberField(fieldID, recordSortDataModel.SortDirectionAsc)
	OrderedBy(fieldAsc).Sort(recValues)
	if recValues[0].RecordID != "Record_02" {
		t.Errorf("TestRecordValueSort: expecting Record_01 first (blanks last)")
	}
	t.Logf("TestRecordValueSort: sort asc results %+v", testUtil.EncodeJSONString(t, recValues))

	fieldDesc := SortByNumberField(fieldID, recordSortDataModel.SortDirectionDesc)
	OrderedBy(fieldDesc).Sort(recValues)
	if recValues[0].RecordID != "Record_01" {
		t.Errorf("TestRecordValueSort: expecting Record_02 first (blanks first)")
	}
	t.Logf("TestRecordValueSort: sort desc results %+v", testUtil.EncodeJSONString(t, recValues))

}

func TestTextValueSort(t *testing.T) {

	fieldID := "fieldA"

	rec1Values := record.RecFieldValues{}
	rec1Values[fieldID] = "ABC"
	recVal1 := recordValue.RecordValueResults{
		RecordID:    "Record_01",
		FieldValues: rec1Values}

	rec2Values := record.RecFieldValues{}
	rec2Values[fieldID] = "DEF"
	recVal2 := recordValue.RecordValueResults{
		RecordID:    "Record_02",
		FieldValues: rec2Values}

	// recValues initialized with records out of order
	recValues := []recordValue.RecordValueResults{recVal2, recVal1}

	fieldAsc := SortByTextField(fieldID, recordSortDataModel.SortDirectionAsc)
	OrderedBy(fieldAsc).Sort(recValues)
	if recValues[0].RecordID != "Record_01" {
		t.Errorf("TestTextValueSort: expecting Record_01 first")
	}
	t.Logf("TestTextValueSort: sort asc results %+v", testUtil.EncodeJSONString(t, recValues))

	fieldDesc := SortByTextField(fieldID, recordSortDataModel.SortDirectionDesc)
	OrderedBy(fieldDesc).Sort(recValues)
	if recValues[0].RecordID != "Record_02" {
		t.Errorf("TestRecordValueSort: expecting Record_02 first")
	}
	t.Logf("TestTextValueSort: sort desc results %+v", testUtil.EncodeJSONString(t, recValues))

}

func TestBoolValueSort(t *testing.T) {

	fieldID := "fieldA"

	rec1Values := record.RecFieldValues{}
	rec1Values[fieldID] = false
	recVal1 := recordValue.RecordValueResults{
		RecordID:    "Record_01",
		FieldValues: rec1Values}

	rec2Values := record.RecFieldValues{}
	rec2Values[fieldID] = true
	recVal2 := recordValue.RecordValueResults{
		RecordID:    "Record_02",
		FieldValues: rec2Values}

	// recValues initialized with records out of order
	recValues := []recordValue.RecordValueResults{recVal2, recVal1}

	fieldAsc := SortByBoolField(fieldID, recordSortDataModel.SortDirectionAsc)
	OrderedBy(fieldAsc).Sort(recValues)
	if recValues[0].RecordID != "Record_01" {
		t.Errorf("TestTextValueSort: expecting Record_01 first")
	}
	t.Logf("TestTextValueSort: sort asc results %+v", testUtil.EncodeJSONString(t, recValues))

	fieldDesc := SortByBoolField(fieldID, recordSortDataModel.SortDirectionDesc)
	OrderedBy(fieldDesc).Sort(recValues)
	if recValues[0].RecordID != "Record_02" {
		t.Errorf("TestRecordValueSort: expecting Record_02 first")
	}
	t.Logf("TestTextValueSort: sort desc results %+v", testUtil.EncodeJSONString(t, recValues))

}

func TestTimeValueSort(t *testing.T) {

	fieldID := "fieldA"

	rec1Values := record.RecFieldValues{}
	timeFormat := "2006-Jan-02"
	time1, _ := time.Parse(timeFormat, "2015-Dec-01")
	rec1Values[fieldID] = time1.Format(time.RFC3339)
	recVal1 := recordValue.RecordValueResults{
		RecordID:    "Record_01",
		FieldValues: rec1Values}

	rec2Values := record.RecFieldValues{}
	time2, _ := time.Parse(timeFormat, "2016-Jan-31")
	rec2Values[fieldID] = time2.Format(time.RFC3339)
	recVal2 := recordValue.RecordValueResults{
		RecordID:    "Record_02",
		FieldValues: rec2Values}

	// recValues initialized with records out of order
	recValues := []recordValue.RecordValueResults{recVal2, recVal1}

	fieldAsc := SortByTimeField(fieldID, recordSortDataModel.SortDirectionAsc)
	OrderedBy(fieldAsc).Sort(recValues)
	if recValues[0].RecordID != "Record_01" {
		t.Errorf("TestTextValueSort: expecting Record_01 first")
	}
	t.Logf("TestTextValueSort: sort asc results %+v", testUtil.EncodeJSONString(t, recValues))

	fieldDesc := SortByTimeField(fieldID, recordSortDataModel.SortDirectionDesc)
	OrderedBy(fieldDesc).Sort(recValues)
	if recValues[0].RecordID != "Record_02" {
		t.Errorf("TestRecordValueSort: expecting Record_02 first")
	}
	t.Logf("TestTextValueSort: sort desc results %+v", testUtil.EncodeJSONString(t, recValues))

}

const numberFieldID string = "numField"
const textFieldID string = "textField"

func createMultiRecordVal(recID string, textVal string, numVal float64) recordValue.RecordValueResults {
	recValues := record.RecFieldValues{}
	recValues[numberFieldID] = numVal
	recValues[textFieldID] = textVal
	recValResults := recordValue.RecordValueResults{
		RecordID:    recID,
		FieldValues: recValues}
	return recValResults
}

func TestMultiFieldValueSort(t *testing.T) {

	rec1Vals := createMultiRecordVal("rec1", "A", 1)
	rec2Vals := createMultiRecordVal("rec2", "A", 2)
	rec3Vals := createMultiRecordVal("rec3", "B", 4)
	rec4Vals := createMultiRecordVal("rec4", "D", 3)

	// recValues initialized with records out of order
	recValues := []recordValue.RecordValueResults{rec4Vals, rec1Vals, rec3Vals, rec2Vals}

	numberAsc := SortByNumberField(numberFieldID, recordSortDataModel.SortDirectionAsc)
	textAsc := SortByTextField(textFieldID, recordSortDataModel.SortDirectionAsc)
	OrderedBy(textAsc, numberAsc).Sort(recValues)
	if recValues[0].RecordID != "rec1" {
		t.Errorf("TestRecordValueSort: expecting rec1 first")
	}
	if recValues[1].RecordID != "rec2" {
		t.Errorf("TestRecordValueSort: expecting rec2 2nd")
	}
	if recValues[2].RecordID != "rec3" {
		t.Errorf("TestRecordValueSort: expecting rec3 3rd")
	}
	if recValues[3].RecordID != "rec4" {
		t.Errorf("TestRecordValueSort: expecting rec4 4th")
	}

	OrderedBy(numberAsc, textAsc).Sort(recValues)
	if recValues[0].RecordID != "rec1" {
		t.Errorf("TestRecordValueSort: expecting rec1 first")
	}
	if recValues[1].RecordID != "rec2" {
		t.Errorf("TestRecordValueSort: expecting rec2 2nd")
	}
	if recValues[2].RecordID != "rec4" {
		t.Errorf("TestRecordValueSort: expecting rec4 3rd")
	}
	if recValues[3].RecordID != "rec3" {
		t.Errorf("TestRecordValueSort: expecting rec3 4th")
	}

	t.Logf("TestRecordValueSort: sort asc results %+v", testUtil.EncodeJSONString(t, recValues))

}
