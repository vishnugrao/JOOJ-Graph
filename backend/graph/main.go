package main

import (
	// External Imports
	"fmt"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	fmt.Println("Server running on port 6767...")
	if err := http.ListenAndServe(":6767", router); err != nil {
		panic(err)
	}

}
