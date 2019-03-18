// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package global

import (
	"database/sql"
	"fmt"
)

func GetFileURL(trackerDBHandle *sql.DB, cloudFileName string) string {
	// TODO - Replace localhost part with dynamically configured host name.
	fileURL := "http://localhost:8080/api/global/getFile/" + cloudFileName

	return fileURL
}

type GetGlobalValUrlParams struct {
	GlobalID      string `json:"globalID"`
	CloudFileName string `json:"cloudFileName"`
}

type GlobalValURLResponse struct {
	Url string `json:"url"`
}

func getGlobalValUrl(trackerDBHandle *sql.DB, params GetGlobalValUrlParams) (*GlobalValURLResponse, error) {

	// TODO check the global is valid.

	if len(params.CloudFileName) == 0 {
		return nil, fmt.Errorf(
			"getGlobalValUrl: Unabled to get global value url with params = %+v", params)
	}

	fileURL := GetFileURL(trackerDBHandle, params.CloudFileName)

	return &GlobalValURLResponse{Url: fileURL}, nil

}
