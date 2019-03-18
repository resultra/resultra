// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package recordSort

import (
	"resultra/tracker/server/common/recordSortDataModel"
	"resultra/tracker/server/recordValue"
	"sort"
)

// ByRecordValue is the type of a "less" function that defines the ordering
// of its RecordValueResults arguments.
type ByRecordValueLessFunc func(p1, p2 recordValue.RecordValueResults) bool

type RecordValueSorter struct {
	RecordValues   []recordValue.RecordValueResults
	MultiSortFuncs []ByRecordValueLessFunc
}

func (recValSorter *RecordValueSorter) Len() int {
	return len(recValSorter.RecordValues)
}

// Swap is part of sort.Interface.
func (recValSorter *RecordValueSorter) Swap(i, j int) {
	recValSorter.RecordValues[i], recValSorter.RecordValues[j] = recValSorter.RecordValues[j], recValSorter.RecordValues[i]
}

func (recValSorter *RecordValueSorter) Less(i, j int) bool {

	// Use indirection to retrieve the corresponding record values for the
	// given indices.
	p, q := recValSorter.RecordValues[i], recValSorter.RecordValues[j]

	// Iterate through the multi-sort functions in order to compare the
	// 2 record values until one is found to be less than the other using
	// the given sort criteria.
	for _, lessFunc := range recValSorter.MultiSortFuncs {
		switch {
		case lessFunc(p, q):
			// p < q => we have a decision
			return true
		case lessFunc(q, p):
			// q < p => we have a decision
			return false
		}
	}
	// All comparisons to this point haven't yielded a decision; i.e., the 2 elements are equal
	// using all the sort functions. So, return whatever the last function returns.
	if len(recValSorter.MultiSortFuncs) > 0 {
		return recValSorter.MultiSortFuncs[len(recValSorter.MultiSortFuncs)-1](p, q)
	}
	return false
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(lessFuncs ...ByRecordValueLessFunc) *RecordValueSorter {
	return &RecordValueSorter{
		MultiSortFuncs: lessFuncs,
	}
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (recValSorter *RecordValueSorter) Sort(recordValues []recordValue.RecordValueResults) {
	recValSorter.RecordValues = recordValues
	sort.Sort(recValSorter)
}

// sortByNumberField uses a closure to create a less sorting function conforming
// to the ByRecordValueLessFunc type.
func SortByNumberField(fieldID string, direction string) ByRecordValueLessFunc {
	if direction == recordSortDataModel.SortDirectionAsc {
		lessAscFunc := func(p1, p2 recordValue.RecordValueResults) bool {

			// The following assignments will try to find the value with the right
			// type and get those values if they are the correct type. This assignment
			// also does the right thing when no value is found or the value is the incorrect type;
			// in both these cases the value will be returned as not found ()
			p1NumberVal, p1Found := p1.FieldValues[fieldID].(float64)
			p2NumberVal, p2Found := p2.FieldValues[fieldID].(float64)

			if p1Found {
				if p2Found {
					// Both p1 & p2 are defined, so do an actual comparison of their
					// numbers.
					return p1NumberVal < p2NumberVal
				} else {
					// Value for p2 is undefined (blank) => p1 < p2 (put blanks at end)
					return true
				}
			} else { // p1 undefined
				if p2Found {
					// p1 is undefined (blank), but p2 has a value => p2 < p1 (put blanks at end)
					return false
				} else {
					// Both p1 & p2 are undefined, return whatever.
					return false
				}
			}
		}
		return lessAscFunc
	} else { // sortDirectionDesc
		lessDescFunc := func(p1, p2 recordValue.RecordValueResults) bool {

			p1NumberVal, p1Found := p1.FieldValues[fieldID].(float64)
			p2NumberVal, p2Found := p2.FieldValues[fieldID].(float64)

			if p1Found {
				if p2Found {
					// Both p1 & p2 are defined, so do an actual comparison of their
					// numbers.
					return p1NumberVal > p2NumberVal
				} else {
					return false
				}
			} else { // p1 undefined
				if p2Found {
					return true
				} else {
					// Both p1 & p2 are undefined, return whatever.
					return false
				}
			}
		}
		return lessDescFunc
	}
}

func SortByTextField(fieldID string, direction string) ByRecordValueLessFunc {
	if direction == recordSortDataModel.SortDirectionAsc {
		lessAscFunc := func(p1, p2 recordValue.RecordValueResults) bool {

			// The following assignments will try to find the value with the right
			// type and get those values if they are the correct type. This assignment
			// also does the right thing when no value is found or the value is the incorrect type;
			// in both these cases the value will be returned as not found ()
			p1Val, p1Found := p1.FieldValues[fieldID].(string)
			p2Val, p2Found := p2.FieldValues[fieldID].(string)

			if p1Found {
				if p2Found {
					// Both p1 & p2 are defined, so do an actual comparison of their
					// numbers.
					return p1Val < p2Val
				} else {
					// Value for p2 is undefined (blank) => p1 < p2 (blanks at end)
					return true
				}
			} else { // p1 undefined
				if p2Found {
					// p1 is undefined (blank), but p2 has a value => p2 < p1
					return false
				} else {
					// Both p1 & p2 are undefined, return whatever.
					return false
				}
			}
		}
		return lessAscFunc
	} else { // sortDirectionDesc
		lessDescFunc := func(p1, p2 recordValue.RecordValueResults) bool {

			p1Val, p1Found := p1.FieldValues[fieldID].(string)
			p2Val, p2Found := p2.FieldValues[fieldID].(string)

			if p1Found {
				if p2Found {
					// Both p1 & p2 are defined, so do an actual comparison of their
					// numbers.
					return p1Val > p2Val
				} else {
					// Value for p2 is undefined (blank) => p2 < p1
					return false
				}
			} else { // p1 undefined
				if p2Found {
					// p1 is undefined (blank), but p2 has a value => p1 < p2
					return true
				} else {
					// Both p1 & p2 are undefined, return whatever.
					return false
				}
			}
		}
		return lessDescFunc
	}
}

func SortByBoolField(fieldID string, direction string) ByRecordValueLessFunc {
	if direction == recordSortDataModel.SortDirectionAsc {
		lessAscFunc := func(p1, p2 recordValue.RecordValueResults) bool {

			// The following assignments will try to find the value with the right
			// type and get those values if they are the correct type. This assignment
			// also does the right thing when no value is found or the value is the incorrect type;
			// in both these cases the value will be returned as not found ()
			p1Val, p1Found := p1.FieldValues[fieldID].(bool)
			p2Val, p2Found := p2.FieldValues[fieldID].(bool)

			if p1Found {
				if p2Found {
					// Both p1 & p2 are defined, so do an actual comparison
					if (p1Val == false) && (p2Val == true) {
						return true // false < true
					} else {
						return false
					}
				} else {
					// Value for p2 is undefined (blank) => p1 < p2 (put blanks at end)
					return true
				}
			} else { // p1 undefined
				if p2Found {
					// p1 is undefined (blank), but p2 has a value => p1 < p2
					return false
				} else {
					// Both p1 & p2 are undefined, return whatever.
					return false
				}
			}
		}
		return lessAscFunc
	} else { // sortDirectionDesc
		lessDescFunc := func(p1, p2 recordValue.RecordValueResults) bool {

			p1Val, p1Found := p1.FieldValues[fieldID].(bool)
			p2Val, p2Found := p2.FieldValues[fieldID].(bool)

			if p1Found {
				if p2Found {
					// Both p1 & p2 are defined, so do an actual comparison
					if (p1Val == true) && (p2Val == false) {
						return true
					} else {
						return false
					}
				} else {
					return false
				}
			} else { // p1 undefined
				if p2Found {
					return true
				} else {
					return false
				}
			}
		}
		return lessDescFunc
	}
}

func SortByTimeField(fieldID string, direction string) ByRecordValueLessFunc {
	if direction == recordSortDataModel.SortDirectionAsc {
		lessAscFunc := func(p1, p2 recordValue.RecordValueResults) bool {

			// The following assignments will try to find the value with the right
			// type and get those values if they are the correct type. This assignment
			// also does the right thing when no value is found or the value is the incorrect type;
			// in both these cases the value will be returned as not found ()
			p1Val, p1Found := p1.FieldValues.GetTimeFieldValue(fieldID)
			p2Val, p2Found := p2.FieldValues.GetTimeFieldValue(fieldID)

			if p1Found {
				if p2Found {
					// Both p1 & p2 are defined, so do an actual comparison of their
					// numbers.
					return p1Val.Before(p2Val)
				} else {
					// Value for p2 is undefined (blank) => p1 < p2 (put blanks at end)
					return true
				}
			} else { // p1 undefined
				if p2Found {
					// p1 is undefined (blank), but p2 has a value => p2 < p1
					return false
				} else {
					// Both p1 & p2 are undefined, return whatever.
					return false
				}
			}
		}
		return lessAscFunc
	} else { // sortDirectionDesc
		lessDescFunc := func(p1, p2 recordValue.RecordValueResults) bool {

			p1Val, p1Found := p1.FieldValues.GetTimeFieldValue(fieldID)
			p2Val, p2Found := p2.FieldValues.GetTimeFieldValue(fieldID)

			if p1Found {
				if p2Found {
					return p2Val.Before(p1Val)
				} else {
					return false
				}
			} else { // p1 undefined
				if p2Found {
					return true
				} else {
					return false // return whatever
				}
			}
		}
		return lessDescFunc
	}
}
