// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package form

import (
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/trackerDatabase"
)

type FormProperties struct {
	Layout componentLayout.ComponentLayout `json:"layout"`
}

func (srcProps FormProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*FormProperties, error) {

	destProps := FormProperties{
		Layout: srcProps.Layout.Clone(cloneParams.IDRemapper)}

	return &destProps, nil
}

func newDefaultFormProperties() FormProperties {
	defaultProps := FormProperties{
		Layout: componentLayout.ComponentLayout{}}

	return defaultProps
}
