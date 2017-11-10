package databaseWrapper

type NetworkDatabaseConfig struct {
	DatabaseHostName      string `json:"databaseHostName"`
	DatabaseUser          string `json:"databaseUser"`
	DatabasePassword      string `json:"databasePassword"`
	AttachmentStoragePath string `json:"attachmentStoragePath"`
}
