// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package valueList

import (
	"github.com/resultra/resultra/server/trackerDatabase"
)

type ValueListValue struct {
	NumValue  *float64 `json:"numValue,omitempty"`
	TextValue *string  `json:"textValue,omitempty"`
}

type ValueListProperties struct {
	Values []ValueListValue `json:"values"`
}

func newDefaultValueListProperties() ValueListProperties {
	defaultProps := ValueListProperties{
		Values: []ValueListValue{}}
	return defaultProps
}

func (srcProps ValueListProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*ValueListProperties, error) {

	destProps := srcProps

	return &destProps, nil
}
