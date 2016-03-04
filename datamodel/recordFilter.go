package datamodel

import ()

const recordFilterEntityKind string = "RecordFilter"

type RecordFilter struct {
	rules []RecordFilterRule
}
