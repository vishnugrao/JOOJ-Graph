package API

import (
	"JOOJ-Graph/backend/model"
	"net/http"
)

var inMemoryStore = map[string]model.Graph{}


func StoreResetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method...", http.StatusBadRequest)
	}
	inMemoryStore = map[string]model.Graph{}
	w.WriteHeader(http.StatusOK)
}
