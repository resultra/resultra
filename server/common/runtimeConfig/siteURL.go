package runtimeConfig

import "fmt"

func GetSiteBaseURL() string {
	return fmt.Sprintf("http://localhost:%v/", CurrRuntimeConfig.PortNumber)
}

func GetSiteResourceURL(resourceSuffix string) string {
	return GetSiteBaseURL() + resourceSuffix
}
