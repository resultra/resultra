package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type JSONParams map[string]string

func DecodeJSONRequest(r *http.Request, decodedVal interface{}) error {

	if err := json.NewDecoder(r.Body).Decode(decodedVal); err != nil {
		return fmt.Errorf("DecodeJSONRequest:Error decoding server JSON request: decode error = %v", err)
	} else {
		log.Printf("INFO: API: Decoded JSON: %+v", decodedVal)
		return nil
	}
}

func WriteJSONResponse(w http.ResponseWriter, responseVals interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encodeErr := json.NewEncoder(w).Encode(responseVals)
	if encodeErr != nil {
		WriteErrorResponse(w, encodeErr)
	}
}

type FileUploadInfo struct {
	FileData   []byte
	FileName   string
	FileLength int
}

func ReadUploadFile(req *http.Request, fileName string) (*FileUploadInfo, error) {

	formFile, handler, formErr := req.FormFile(fileName)
	if formErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to upload file: invalid api/form input in request: %v req = %+v", formErr, req)
	}
	log.Printf("ReadUploadFile: %v", handler.Filename)

	fileContents, readErr := ioutil.ReadAll(formFile)
	if formErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to read file contents: %v", readErr)
	}
	fileLength := len(fileContents)
	log.Printf("ReadUploadFile: got file contents: length (bytes) = %v", fileLength)

	uploadInfo := FileUploadInfo{
		FileData:   fileContents,
		FileName:   handler.Filename,
		FileLength: fileLength}

	return &uploadInfo, nil

}
