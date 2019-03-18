// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package formPage

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {

	mainRouter.HandleFunc("/submitForm/{sharedLinkID}", submitFormPage)

	mainRouter.Path("/viewItem/{formID}/{recordID}").HandlerFunc(viewFormPage)

	mainRouter.Path("/viewItem/{formID}/{recordID}").Queries("col", "{col}").HandlerFunc(viewFormPage)
	mainRouter.Path("/viewItem/{formID}/{recordID}").Queries("frm", "{frm}").HandlerFunc(viewFormPage)

}
