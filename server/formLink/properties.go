// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package formLink

import (
	"fmt"
	"github.com/resultra/resultra/server/record"
	"github.com/resultra/resultra/server/trackerDatabase"
)

type FormLinkProperties struct {
	DefaultValues []record.DefaultFieldValue `json:"defaultValues"`
}

func newDefaultNewItemProperties() FormLinkProperties {
	defaultProps := FormLinkProperties{
		DefaultValues: []record.DefaultFieldValue{}}
	return defaultProps
}

func (srcProps FormLinkProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*FormLinkProperties, error) {

	destProps := srcProps

	destDefaultVals, cloneErr := record.CloneDefaultFieldValues(cloneParams.IDRemapper, srcProps.DefaultValues)
	if cloneErr != nil {
		return nil, fmt.Errorf("FormLinkProperties.Clone: %v", cloneErr)
	}

	destProps.DefaultValues = destDefaultVals

	return &destProps, nil

}
