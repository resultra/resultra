package runtimeConfig

import (
	"fmt"
	"golang.org/x/oauth2/jwt"
	"resultra/datasheet/server/generic/cloudStorageWrapper"
)

const CloudStorageBucketName string = "resultra-db-dev"
const cloudStorageJSONAuthInfoFile string = "/Users/sroehling/Development/Datasheet-Dev-60167588e163.json"

var CloudStorageAuthConfig *jwt.Config

func init() {

	cloudStorageAuthConfig, configErr := cloudStorageWrapper.ReadServiceAuthConfig(cloudStorageJSONAuthInfoFile)
	if configErr != nil {
		panic(fmt.Sprintf("Server startup configuration: Unable to read cloud storage credentials: %v", configErr))
	}
	CloudStorageAuthConfig = cloudStorageAuthConfig

}
