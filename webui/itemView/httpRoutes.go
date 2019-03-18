// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
