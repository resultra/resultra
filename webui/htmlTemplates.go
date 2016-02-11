package webui

import "html/template"

// Parse all the HTML templates at once. Individual templates can then
// be referenced throughout this package using htmlTemplates.ExecuteTemplate(...)
var htmlTemplates = template.Must(template.ParseGlob("template/*.html"))
