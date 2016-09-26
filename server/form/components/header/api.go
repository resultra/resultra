package header

import (
	"github.com/gorilla/mux"
	"net/http"
	"resultra/datasheet/server/generic/api"
)

func init() {
	headerRouter := mux.NewRouter()

	headerRouter.HandleFunc("/api/frm/header/new", newHeader)
	//	headerRouter.HandleFunc("/api/frm/header/resize", resizeHeader)

	http.Handle("/api/frm/header/", headerRouter)
}

func newHeader(w http.ResponseWriter, r *http.Request) {

	params := NewHeaderParams{}
	if err := api.DecodeJSONRequest(r, &params); err != nil {
		api.WriteErrorResponse(w, err)
		return
	}

	if headerRef, err := saveNewHeader(params); err != nil {
		api.WriteErrorResponse(w, err)
	} else {
		api.WriteJSONResponse(w, *headerRef)
	}

}
