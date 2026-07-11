package main

import (
	api "JOOJ-Graph/backend/API"
	"fmt"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/create/graph", api.CreateGraphHandler)
	router.HandleFunc("/create/node", api.CreateNodeHandler)

	fmt.Println("Server running on port 6767...")
	if err := http.ListenAndServe(":6767", router); err != nil {
		panic(err)
	}

}
