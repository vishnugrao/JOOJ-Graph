package main

import (
	"fmt"
	"net/http"

	"github.com/abhay/JOOJ-Graph/backend/api"
)

func main() {
	
	router := http.NewServeMux()

	router.HandleFunc("/inputdump", api.InputDump)


	fmt.Println("Server running on port 6767...")
	if err:= http.ListenAndServe(":6767", router); err != nil {
		panic(err)
	}

}