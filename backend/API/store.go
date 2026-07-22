package API

import (
// Internal imports
	"JOOJ-Graph/backend/model"

// External imports
	"net/http"
)

var inMemoryStore = map[string]model.Graph{}


func StoreResetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method...", http.StatusBadRequest)
		return
	}
	inMemoryStore = map[string]model.Graph{}
	w.WriteHeader(http.StatusOK)
}
