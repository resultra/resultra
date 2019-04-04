// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package main

import (
	"flag"
	"fmt"
	"github.com/resultra/resultra/server"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/common/runtimeConfig"
	"github.com/resultra/resultra/webui"
	"log"
	"net/http"
	"os"

	"github.com/pkg/profile"
)

const staticSiteResourcesPrefix string = `/static/`
const defaultLocalPortNumber int = 43409

func init() {

	// The following dummy functions are called to legitimize the includes
	// of the server and webui packages. In other words, these includes
	// are needed so the packages are compiled into the Google App Engine
	// executable.
	webui.DummyFunctionForImportFromGoogleAppEngineProjectFolder()
	server.DummyFunctionForImportFromGoogleAppEngineProjectFolder()
}

func main() {

	trackerBasePath := flag.String("tracker-path", "", "Tracker database and attachment base path")
	templatesBasePath := flag.String("templates-path", "", "Factory templates base path")
	enableProfiling := flag.Bool("profile", false, "Enable pprof performance profiling")
	flag.Parse()

	if (trackerBasePath == nil) || (len(*trackerBasePath) == 0) {
		log.Printf("ERROR: Tracker database base path is required")
		os.Exit(255)
	}
	trackerDBConfig := databaseWrapper.LocalSQLiteTrackerDatabaseConnectionConfig{DatabaseBasePath: *trackerBasePath}

	attachmentBasePath := (*trackerBasePath) + `/attachments`
	attachmentConfig := databaseWrapper.LocalAttachmentStorageConfig{AttachmentBasePath: attachmentBasePath}

	if (templatesBasePath == nil) || (len(*templatesBasePath) == 0) {
		log.Printf("ERROR: Template base path is required")
		os.Exit(255)
	}
	localTemplateDBConfig := databaseWrapper.LocalSQLiteTrackerDatabaseConnectionConfig{DatabaseBasePath: *templatesBasePath}
	factoryTemplateConfig := runtimeConfig.FactoryTemplateDatabaseConfig{LocalDatabaseConfig: &localTemplateDBConfig}

	config := runtimeConfig.NewDefaultRuntimeConfig()
	config.TrackerDatabaseConfig.LocalDatabaseConfig = &trackerDBConfig
	config.TrackerDatabaseConfig.LocalAttachmentConfig = &attachmentConfig
	config.FactoryTemplateDatabaseConfig = &factoryTemplateConfig
	isSingleWorkspace := true
	config.IsSingleUserWorkspace = &isSingleWorkspace

	config.ServerConfig.ListenPortNumber = defaultLocalPortNumber
	siteURL := runtimeConfig.LocalHostBaseURL()
	config.ServerConfig.SiteBaseURL = &siteURL
	// Hard-code the cookie keys. Since the service is only used locally and the user is
	// not authencicated, this should sufice.
	config.ServerConfig.CookieAuthKey = "nRrHLlHcHH0u7fUxyzHje9m7uJ5SnJzP"
	config.ServerConfig.CookieEncryptionKey = "CAp1KsJncuMzARfookqSFLqsBi5ag2bE"

	if err := runtimeConfig.InitRuntimeConfig(config); err != nil {
		log.Printf("Error setting configuration: %v\n", err)
		os.Exit(255)
	}

	if *enableProfiling {
		log.Println("Profiling enabled (pprof)")
		defer profile.Start().Stop()
	}

	// Serve static CSS, Javascript and image files from a common "static" directory.
	http.Handle(staticSiteResourcesPrefix, http.StripPrefix(staticSiteResourcesPrefix,
		http.FileServer(http.Dir("./static"))))

	log.Printf("Server started: listening on port: %v", runtimeConfig.CurrRuntimeConfig.ServerConfig.ListenPortNumber)
	// Only listen on the localhost/loopback port. T
	portNumString := fmt.Sprintf("127.0.0.1:%v", runtimeConfig.CurrRuntimeConfig.ServerConfig.ListenPortNumber)
	http.ListenAndServe(portNumString, nil)

}
