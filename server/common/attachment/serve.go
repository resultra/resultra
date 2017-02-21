package attachment

func GetAttachmentURL(cloudFileName string) string {
	// TODO - Replace localhost part with dynamically configured host name.
	fileURL := "http://localhost:8080/api/record/getFile/" + cloudFileName

	return fileURL
}
