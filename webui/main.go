package webui

import (
	"html/template"
	"net/http"
)

// Parse the templates once
var templates = template.Must(template.ParseFiles("template/main.html"))

type PageInfo struct {
	Title string `json:"title"`
}

func root(w http.ResponseWriter, r *http.Request) {
	//	c := appengine.NewContext(r)

	p := PageInfo{"Home Page"}
	err := templates.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
