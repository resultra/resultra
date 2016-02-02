package webui

import (
	"html/template"
	"net/http"
)

var editLayoutTemplates = template.Must(template.ParseFiles("template/editLayout.html"))

func editLayout(w http.ResponseWriter, r *http.Request) {
	//	c := appengine.NewContext(r)

	p := PageInfo{"Edit Layout"}
	err := editLayoutTemplates.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
