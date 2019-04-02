// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package header

import (
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/trackerDatabase"
)

type HeaderProps struct {
	Geometry   componentLayout.LayoutGeometry `json:"geometry"`
	Title      string                         `json:"title"`
	Size       string                         `json:"size"`
	Underlined bool                           `json:"underlined"`
}

func (srcProps HeaderProps) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*HeaderProps, error) {

	destProps := srcProps
	return &destProps, nil

}

func newDefaultHeaderProps() HeaderProps {
	props := HeaderProps{
		Title:      "Header",
		Size:       "medium",
		Underlined: false}
	return props
}
