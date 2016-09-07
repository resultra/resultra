package global

import (
	"fmt"
)

func GetFileURL(cloudFileName string) string {
	// TODO - Replace localhost part with dynamically configured host name.
	fileURL := "http://localhost:8080/api/global/getFile/" + cloudFileName

	return fileURL
}

type GetGlobalValUrlParams struct {
	GlobalID      string `json:"globalID"`
	CloudFileName string `json:"cloudFileName"`
}

type GlobalValURLResponse struct {
	Url string `json:"url"`
}

func getGlobalValUrl(params GetGlobalValUrlParams) (*GlobalValURLResponse, error) {

	// TODO check the global is valid.

	if len(params.CloudFileName) == 0 {
		return nil, fmt.Errorf(
			"getGlobalValUrl: Unabled to get global value url with params = %+v", params)
	}

	fileURL := GetFileURL(params.CloudFileName)

	return &GlobalValURLResponse{Url: fileURL}, nil

}
