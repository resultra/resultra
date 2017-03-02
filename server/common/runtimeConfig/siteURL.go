package runtimeConfig

func GetSiteBaseURL() string {
	return "http://localhost:8080/"
}

func GetSiteResourceURL(resourceSuffix string) string {
	return GetSiteBaseURL() + resourceSuffix
}
