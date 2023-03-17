package a

import (
	"os"
)

func main() {
	password := "password123"       // want "hardcoded credential found: password123"
	password = os.Getenv("DB_PASS") // ok
	_ = password

	cred := "credential123"     // want "hardcoded credential found: credential123"
	cred = os.Getenv("DB_CRED") // ok
	_ = cred

	apiKey := "apikey123"         // want "hardcoded credential found: apikey123"
	apiKey = os.Getenv("API_KEY") // ok
	_ = apiKey
}
