package common

import "os"

// GetApplicationDomain returns the current application domain name
func GetApplicationDomain() string {

	var domain string

	val, ok := os.LookupEnv("GO_SERVER_APP_DOMAIN")
	if !ok {
		domain = "localhost"
	} else {
		domain = val
	}

	return domain
}

// GetApplicationEndPoint returns the current application domain name
func GetApplicationEndPoint() string {

	var domain string

	val, ok := os.LookupEnv("GO_SERVER_APP_ENDPOINT")
	if !ok {
		domain = GoServerAppEndPoint
	} else {
		domain = val
	}

	return domain
}
