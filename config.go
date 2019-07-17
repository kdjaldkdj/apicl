package apicl

import "os"

// ConfigAPI detects if APIARY_API_KEY has been set.

func CheckConfigEnv() bool {
	apikey := os.Getenv("APIARY_API_KEY")
	if apikey == "" {
		return false
	}
	return true
}
