// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package api

import (
	"log"
	"net/http"
)

func WriteErrorResponse(w http.ResponseWriter, err error) {
	// TBD - Also log the error somewhere
	log.Printf("ERROR: Couldn't process server request: %v", err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
