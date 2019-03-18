package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"resultra/tracker/server"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/databaseController"
	"resultra/tracker/server/generic/timestamp"
	"resultra/tracker/server/generic/uniqueID"
	"resultra/tracker/server/trackerDatabase"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func init() {

	// The following dummy functions are called to legitimize the includes
	// of the server and webui packages. In other words, these includes
	// are needed so the packages are compiled into the Google App Engine
	// executable.
	server.DummyFunctionForImportFromGoogleAppEngineProjectFolder()
}

const templateAdminUserName string = "templateAdmin"

func initDestDBDatabase(destDBFileName string) (*sql.DB, error) {

	dbFileAlreadyExists := databaseWrapper.TrackerDatabaseFileExists(destDBFileName)

	dbHandle, err := sql.Open("sqlite3", destDBFileName)
	if err != nil {
		return nil, fmt.Errorf("can't establish connection to database %v: %v", destDBFileName, err)
	}

	// Configure the maximum number of open connections.
	dbHandle.SetMaxOpenConns(25)

	// Open doesn't directly open the database connection. To verify the connection, the Ping() function
	// is needed.
	if err := dbHandle.Ping(); err != nil {
		return nil, fmt.Errorf("can't establish connection to database: %v", err)
	}

	log.Printf("Database connection established: %v", destDBFileName)

	if !dbFileAlreadyExists {
		log.Printf("New database found, initializing: %v", destDBFileName)
		if initErr := databaseWrapper.InitNewTrackerDatabaseToDest(dbHandle); initErr != nil {
			return nil, fmt.Errorf("failure initializing tracker database: %v", initErr)
		} else {
			log.Printf("New database initialization complete: %v", destDBFileName)
		}

		templateDBUserParams := userAuth.RegisterSingleUserParams{
			FirstName: templateAdminUserName,
			LastName:  templateAdminUserName,
			UserName:  templateAdminUserName}

		authResp := userAuth.SaveNewSingleUser(dbHandle, templateDBUserParams)
		if !authResp.Success {
			return nil, fmt.Errorf("Can't create template user info in destination database %v: %v", destDBFileName, authResp.Msg)
		}

	} else {
		log.Printf("Existing tracker database found.")
	}

	log.Printf("Done initializing local sqlite connection: %v \n", destDBFileName)

	return dbHandle, nil

}

func openSourceDB(sourceDBFileName string) (*sql.DB, error) {
	dbFileAlreadyExists := databaseWrapper.TrackerDatabaseFileExists(sourceDBFileName)

	if !dbFileAlreadyExists {
		return nil, fmt.Errorf("Can't open source database %v: file doesn't exist", sourceDBFileName)
	}

	dbHandle, err := sql.Open("sqlite3", sourceDBFileName)
	if err != nil {
		return nil, fmt.Errorf("can't establish connection to source database %v: %v", sourceDBFileName, err)
	}

	// Configure the maximum number of open connections.
	dbHandle.SetMaxOpenConns(25)

	return dbHandle, nil

}

func main() {

	sourceDBFile := flag.String("sourcedb", "", "Source tracker database file for templates")
	sourcdDBID := flag.String("source-tracker-db-id", "", "Unique tracker database ID in sourcedb")
	destDBFile := flag.String("destdb", "", "Destination tracker database file for templates")
	testIDs := flag.Bool("test-ids", false, "Generate template with test versions of unique ID and time stamp generators.")
	flag.Parse()

	if (sourceDBFile != nil) && (len(*sourceDBFile) > 0) {
		log.Printf("Opening %v as source tracker database", *sourceDBFile)
	} else {
		log.Printf("ERROR: Missing --sourcedb option for source template database")
		os.Exit(255)
	}
	sourceDBFileName := *sourceDBFile

	if (destDBFile != nil) && (len(*destDBFile) > 0) {
		log.Printf("Opening %v as destination tracker database", *destDBFile)
	} else {
		log.Printf("ERROR: Missing --destdb option for destination template database")
		os.Exit(255)
	}
	destDBFileName := *destDBFile

	if (sourcdDBID != nil) && (len(*sourcdDBID) > 0) {
		log.Printf("Using %v as source tracker database ID", *sourcdDBID)
	} else {
		log.Printf("ERROR: Missing --source-tracker-db-id option for source tracker database ID")
		os.Exit(255)
	}
	sourceDBID := *sourcdDBID

	sourceDBHandle, initDestErr := openSourceDB(sourceDBFileName)
	if initDestErr != nil {
		log.Printf("ERROR: Can't open source database %v: %v", sourceDBFileName, initDestErr)
		os.Exit(255)
	}

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(sourceDBHandle, sourceDBID)
	if dbInfoErr != nil {
		log.Printf("ERROR: Unable to retrieve database info for tracker database ID = %v: %v", sourceDBID, dbInfoErr)
		os.Exit(255)
	}

	destDBHandle, openSrcErr := initDestDBDatabase(destDBFileName)
	if openSrcErr != nil {
		log.Printf("ERROR: Can't initialize destination database %v: %v", destDBFileName, openSrcErr)
		os.Exit(255)
	}

	if *testIDs {
		currTestUniqueID := 0
		testIDFunc := func() string {
			currTestUniqueID++
			return fmt.Sprintf("UNIQUEID%04d", currTestUniqueID)
		}
		uniqueID.OverrideProductionUniqueIDFuncWithTestFunc(testIDFunc)

		currTimestamp := time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)
		testTimestampFunc := func() time.Time {
			currTimestamp = currTimestamp.Add(1 * time.Second)
			return currTimestamp
		}
		timestamp.OverrideProductionTimestampGeneratorWithTestFunc(testTimestampFunc)
	}

	cloneParams := trackerDatabase.CloneDatabaseParams{
		SourceDatabaseID: sourceDBID,
		NewName:          dbInfo.DatabaseName,
		IsTemplate:       true,
		CreatedByUserID:  templateAdminUserName,
		SrcDBHandle:      sourceDBHandle,
		DestDBHandle:     destDBHandle,
		IDRemapper:       uniqueID.UniqueIDRemapper{}}
	_, cloneErr := databaseController.CloneIntoNewTrackerDatabase(&cloneParams)
	if cloneErr != nil {
		log.Printf("ERROR: Error cloning tracker database: %v", cloneErr)
		os.Exit(255)
	}

	os.Exit(0)
}
