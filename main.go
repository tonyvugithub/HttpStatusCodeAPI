package main

import (
	"fmt"

	"github.com/tonyvugithub/statusCodeApi/api"
)

func main() {
	fmt.Println("Status Code API is up and running")
	api.HandleRequests()
}
