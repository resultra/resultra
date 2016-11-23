package databaseController

import (
	"fmt"
	"resultra/datasheet/server/database"
	"resultra/datasheet/server/table"
)

type SaveTemplateParams struct {
	SourceDatabaseID string `json:"sourceDatabaseID"`
	NewTemplateName  string `json:"newTemplateName"`
}

func saveDatabaseToTemplate(params SaveTemplateParams) (*database.Database, error) {

	remappedIDs := map[string]string{}

	templateDB, err := database.CloneDatabase(remappedIDs, params.NewTemplateName, params.SourceDatabaseID)
	if err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	if err := table.CloneTables(remappedIDs, params.SourceDatabaseID); err != nil {
		return nil, fmt.Errorf("copyDatabaseToTemplate: %v", err)
	}

	return templateDB, nil

}
