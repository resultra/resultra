package recordFilter

import (
	"github.com/gorilla/mux"
	"net/http"
)

func init() {

	filterRouter := mux.NewRouter()

	http.Handle("/api/filter/", filterRouter)
}
