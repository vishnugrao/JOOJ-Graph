package API

import (
// Internal imports
	"JOOJ-Graph/backend/model"

// External imports
	"errors"
	"regexp"
	"net/http"
	"encoding/json"
)

var (
	valid_tag_field = regexp.MustCompile(`^[a-zA-Z\/]+$`)
)

func CreateUserEdge(Source model.User_node, Target model.User_node, Edge_tag string, Edge_desc string) (model.Edge, error) {
	if Source.User_id == "" || Target.User_id == "" || Edge_tag == "" || Edge_desc == "" {
		return model.Edge{}, errors.New("No fields can be empty, please try again...")}

	if valid_field.MatchString(Edge_tag) || valid_field.MatchString(Edge_desc) {
		return model.Edge{}, errors.New("Invalid fields, please try again...")}

	if Source.User_id == Target.User_id {
		return model.Edge{}, errors.New("Cannot have an edge to yourself, please try again...")}

	if !valid_tag_field.MatchString(Edge_tag) {
		return model.Edge{}, errors.New("Edge tag can only contain upper/lower case letters and / only, please try again...")}
	
	edge := model.Edge{
		Source: Source,
		Target: Target,
		Edge_tag: Edge_tag,
		Edge_desc: Edge_desc }

	return edge, nil
}

func CreateEdgeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method...", http.StatusBadRequest)
		return }

	var edge model.Edge
	if err := json.NewDecoder(r.Body).Decode(&edge); err != nil {
		http.Error(w, "Invalid JSON...", http.StatusBadRequest)
		return
	}
	edge, err := CreateUserEdge(edge.Source, edge.Target, edge.Edge_tag, edge.Edge_desc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(edge)
}