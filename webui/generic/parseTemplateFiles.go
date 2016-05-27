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
