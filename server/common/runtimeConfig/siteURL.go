// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package runtimeConfig

import "fmt"
import "strings"

func LocalHostBaseURL() string {
	return fmt.Sprintf("http://localhost:%v/", CurrRuntimeConfig.ServerConfig.ListenPortNumber)
}

func GetSiteBaseURL() string {
	return *CurrRuntimeConfig.ServerConfig.SiteBaseURL
}

func GetSiteResourceURL(resourceSuffix string) string {

	siteBaseURL := GetSiteBaseURL()

	sep := ""
	if !strings.HasSuffix(siteBaseURL, "/") {
		sep = "/"
	}

	return GetSiteBaseURL() + sep + resourceSuffix
}
