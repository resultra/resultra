package homePage

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/", home)
}

// Parse the templates once
var homePageTemplates = template.Must(template.ParseFiles("static/homePage/homePage.html"))

type PageInfo struct {
	Title string `json:"title"`
}

func home(w http.ResponseWriter, r *http.Request) {
	//	c := appengine.NewContext(r)

	p := PageInfo{"Home Page"}
	err := homePageTemplates.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
