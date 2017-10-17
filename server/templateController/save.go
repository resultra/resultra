package templateController

type SaveTemplateParams struct {
	srcDatabaseID string `json:"srcDatabaseID"`
}

func saveTemplate(currUserID string, params SaveTemplateParams) error {
	return nil
}
