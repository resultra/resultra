// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package header

import (
	"fmt"
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/form/components/common"
	"github.com/resultra/resultra/server/trackerDatabase"
)

const headerSizeMedium string = "medium"

type HeaderProperties struct {
	Label      string                         `json:"label"`
	Geometry   componentLayout.LayoutGeometry `json:"geometry"`
	HeaderSize string                         `json:"headerSize"`
	Underlined bool                           `json:"underlined"`
	common.ComponentVisibilityProperties
}

func (srcProps HeaderProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*HeaderProperties, error) {

	destProps := srcProps

	destVisibilityConditions, err := srcProps.VisibilityConditions.Clone(cloneParams)
	if err != nil {
		return nil, fmt.Errorf("CaptionProperties.Clone: %v")
	}
	destProps.VisibilityConditions = *destVisibilityConditions

	return &destProps, nil
}

func newDefaultHeaderProperties() HeaderProperties {
	props := HeaderProperties{
		ComponentVisibilityProperties: common.NewDefaultComponentVisibilityProperties(),
		HeaderSize:                    headerSizeMedium,
		Underlined:                    false}
	return props
}
