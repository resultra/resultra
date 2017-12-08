package webui

import (
	"resultra/datasheet/server/common/userAuth"
)

// Dummy variables to force inclusion of the packages (and not trigger an error from the Golang compiler).
// This is needed since these packages are essentially plug-ins which register their own HTTP handlers upon startup.
var dummyUserAuthVar = userAuth.DummyStructForPkgImport{}

func DummyFunctionForImportFromGoogleAppEngineProjectFolder() {
	// This dummy function is needed so standaline packages inside
	// the server will be compiled into the google app engine executable.
	// The stand-alone packages won't be compiled in unless they are included somewhere.

}
