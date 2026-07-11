package API

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type SimpleNode struct {
	Uid       string `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"Last_name"`
	Email     string `json:"email"`
}

func InputDump(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method...", http.StatusBadRequest)
	} else {
		var node SimpleNode
		if err := json.NewDecoder(r.Body).Decode(&node); err != nil {
			http.Error(w, "Invalid JSON...", http.StatusBadRequest)
			return
		}
		node.Uid = uuid.NewString()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(node)
	}
}
