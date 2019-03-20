// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package main

import (
	"flag"
	"fmt"
	"github.com/pkg/profile"
	"log"
	"net/http"
	"os"
	"resultra/tracker/server"
	"resultra/tracker/server/common/runtimeConfig"
	"resultra/tracker/webui"
)

const staticSiteResourcesPrefix string = `/static/`

func init() {

	// The following dummy functions are called to legitimize the includes
	// of the server and webui packages. In other words, these includes
	// are needed so the packages are compiled into the Google App Engine
	// executable.
	webui.DummyFunctionForImportFromGoogleAppEngineProjectFolder()
	server.DummyFunctionForImportFromGoogleAppEngineProjectFolder()
}

func main() {

	configFile := flag.String("config", "", "Configuration file")
	enableProfiling := flag.Bool("profile", false, "Enable pprof performance profiling")
	flag.Parse()

	if (configFile != nil) && (len(*configFile) > 0) {
		if err := runtimeConfig.InitConfigFromConfigFile(*configFile); err != nil {
			log.Printf("Error setting configuration: %v\n", err)
			os.Exit(255)
		}
	}

	if *enableProfiling {
		log.Println("Profiling enabled (pprof)")
		defer profile.Start().Stop()
	}

	// Serve static CSS, Javascript and image files from a common "static" directory.
	http.Handle(staticSiteResourcesPrefix, http.StripPrefix(staticSiteResourcesPrefix,
		http.FileServer(http.Dir("./static"))))

	log.Printf("Server started: listening on port: %v", runtimeConfig.CurrRuntimeConfig.ServerConfig.ListenPortNumber)
	portNumString := fmt.Sprintf(":%v", runtimeConfig.CurrRuntimeConfig.ServerConfig.ListenPortNumber)

	listenErr := http.ListenAndServe(portNumString, nil)

	log.Fatal(fmt.Errorf("Error starting up server on given port: %v", listenErr))

}
