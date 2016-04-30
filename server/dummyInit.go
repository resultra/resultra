package server

import (
	"resultra/datasheet/server/database"
	"resultra/datasheet/server/table"
)

var dummyUnuzedDBParams = database.NewDatabaseParams{}
var dummyUnuzedTableParams = table.NewTableParams{}

func DummyFunctionForImportFromGoogleAppEngineProjectFolder() {
	// This dummy function is needed so standaline packages inside
	// the server will be compiled into the google app engine executable.
	// The stand-alone packages won't be compiled in unless they are included somewhere.
}
