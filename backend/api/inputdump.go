package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Exnode struct {
	First_name string `json:"first name"`
	Last_name string `json:"last name"`
	Email string `json:"Email"`
}

type ExnodeWrapper struct {
	Exnode Exnode `json:"user"`
}


func InputDump(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("data.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var exnode ExnodeWrapper
	err = json.Unmarshal(data, &exnode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("First name: ", exnode.Exnode.First_name)
	fmt.Println("Last name: ", exnode.Exnode.Last_name)
	fmt.Println("Email: ", exnode.Exnode.Email)
}