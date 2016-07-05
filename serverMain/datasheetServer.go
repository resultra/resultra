package main

import (
	"log"
	"net/http"
	"resultra/datasheet/server"
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

	// Serve static CSS, Javascript and image files from a common "static" directory.
	http.Handle(staticSiteResourcesPrefix, http.StripPrefix(staticSiteResourcesPrefix,
		http.FileServer(http.Dir("./static"))))

	log.Println("Server started: listening on port 8080")

	http.ListenAndServe(":8080", nil)

}
