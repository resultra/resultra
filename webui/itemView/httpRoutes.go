package itemView

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {

	mainRouter.HandleFunc("/itemView/newItemContentLayout", newItemContentLayout)
	mainRouter.HandleFunc("/itemView/newItemOffPageContent", newItemOffPageContent)

	mainRouter.HandleFunc("/itemView/existingItemContentLayout", existingItemContentLayout)
	mainRouter.HandleFunc("/itemView/existingItemOffPageContent", existingItemOffPageContent)
}
