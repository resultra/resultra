package main

import (
	"flag"
	"fmt"
	"github.com/pkg/profile"
	"log"
	"net/http"
	"os"
	"resultra/datasheet/server"
	"resultra/datasheet/server/common/attachment"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/common/runtimeConfig"
	"resultra/datasheet/webui"
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
		if err := runtimeConfig.InitConfig(*configFile); err != nil {
			log.Printf("Error setting configuration: %v\n", err)
			os.Exit(255)
		}
	}
	if err := attachment.InitAttachmentBasePath(); err != nil {
		log.Printf("Error initializing attachment directory: %v\n", err)
		os.Exit(255)
	}

	if err := databaseWrapper.InitDatabaseConnection(); err != nil {
		log.Printf("Error initializing database connection: %v\n", err)
		os.Exit(255)

	}

	runtimeConfig.PrintCurrentConfig()

	if *enableProfiling {
		log.Println("Profiling enabled (pprof)")
		defer profile.Start().Stop()
	}

	// Serve static CSS, Javascript and image files from a common "static" directory.
	http.Handle(staticSiteResourcesPrefix, http.StripPrefix(staticSiteResourcesPrefix,
		http.FileServer(http.Dir("./static"))))

	log.Printf("Server started: listening on port: %v", runtimeConfig.CurrRuntimeConfig.PortNumber)
	portNumString := fmt.Sprintf(":%v", runtimeConfig.CurrRuntimeConfig.PortNumber)
	http.ListenAndServe(portNumString, nil)

}
