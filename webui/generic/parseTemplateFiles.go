// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package generic

import (
	"html/template"
)

// ParseTemplatesFromFileLists is a helper function to create a combined set of HTML templates
// from a list of list of files. Each package exports a list of template files. A
// file which then depends on other packages can then easily create a set of merged
// templates.
func ParseTemplatesFromFileLists(fileLists [][]string) *template.Template {
	templateFiles := []string{}
	for _, fileList := range fileLists {
		templateFiles = append(templateFiles, fileList...)
	}
	return template.Must(template.ParseFiles(templateFiles...))
}
