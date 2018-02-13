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

func ReadUploadFile(w http.ResponseWriter, req *http.Request, fileName string) (*FileUploadInfo, error) {

	const maxMBUpload int64 = 25
	const maxBytesUpload int64 = maxMBUpload * 1024 * 1024

	// MaxBytesReader is the preferred method to detect if the file upload size is too large, since
	// the request headers content length could potentially be manipulated by the client.
	//
	// IMPORTANT: Before calling ReadUploadFile, don't call any functions in the http package which
	// may in turn call http.ParseMultipartForm. The MaxBytesReader method is based around the
	// call to http.ParseMultipartForm to get its error message.
	req.Body = http.MaxBytesReader(w, req.Body, maxBytesUpload)
	formFile, handler, formErr := req.FormFile(fileName)
	if formErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to upload file: invalid api/form input in request: %v req = %+v", formErr, req)
	}
	log.Printf("ReadUploadFile: %v", handler.Filename)

	// Assuming the code above with http.MaxBytesReader does it's job, additional error handling shouldn't
	// need to happen here. However, this secondary check also ensures the file uploads isn't too large.
	if req.ContentLength > maxBytesUpload {
		return nil, fmt.Errorf("Failure uploading file: attachment too large: max bytes = %v, got = %v", maxBytesUpload, req.ContentLength)
	}

	fileContents, readErr := ioutil.ReadAll(formFile)
	if formErr != nil {
		return nil, fmt.Errorf("uploadFile: Unable to read file contents: %v", readErr)
	}
	fileLength := len(fileContents)
	if int64(fileLength) > maxBytesUpload {
		return nil, fmt.Errorf("uploadFile: maximum bytes exceeded: max = %v, got %v", maxBytesUpload, fileLength)
	}
	log.Printf("ReadUploadFile: got file contents: length (bytes) = %v", fileLength)

	uploadInfo := FileUploadInfo{
		FileData:   fileContents,
		FileName:   handler.Filename,
		FileLength: fileLength}

	return &uploadInfo, nil

}
