package helpers

import "os"

//Port to get the port. In development, use port 8080"
func Port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
