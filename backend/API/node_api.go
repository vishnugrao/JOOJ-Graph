package API

import (
// Internal imports
	"JOOJ-Graph/backend/model"

//  External imports	
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"github.com/google/uuid"
)

var (
	valid_email = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	valid_name  = regexp.MustCompile(`^[a-zA-Z\s\-']+$`)
	valid_field = regexp.MustCompile(`^\s+$`)
)

func CreateUserNode(First_name string, Last_name string, Profile_picture_url string, Email string) (model.User_node, error) {
	if First_name == "" || Last_name == "" || Email == "" {
		return model.User_node{}, errors.New("User ID, First name, Last name, Email or Created At date cannot be empty, please try again...") }

	if !valid_email.MatchString(Email) {
		return model.User_node{}, errors.New("Invalid email, please try again...") }

	if !valid_name.MatchString(First_name) || !valid_name.MatchString(Last_name) {
		return model.User_node{}, errors.New("Invalid First or Last name, please try again...") }

	if valid_field.MatchString(First_name) || valid_field.MatchString(Last_name) || valid_field.MatchString(Email) {
		return model.User_node{}, errors.New("Invalid fields, please try again...") }

	user := model.User_node{
		User_id: uuid.NewString(),
		First_name: First_name,
		Last_name: Last_name,
		Profile_picture_url: Profile_picture_url,
		Email: Email }

	return user, nil
}

func CreateNodeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method...", http.StatusBadRequest)
		return
	} else {
		var node model.User_node
		
		if err := json.NewDecoder(r.Body).Decode(&node); err != nil {
			http.Error(w, "Invalid JSON...", http.StatusBadRequest)
			return }

		node, err := CreateUserNode(node.First_name, node.Last_name, node.Profile_picture_url, node.Email)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return }

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(node)
	}
}

func GetNodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid Method...", http.StatusBadRequest)
		return }

	GraphName := r.URL.Query().Get("graph_name")
	email :=r.URL.Query().Get("email")

	if GraphName == "" || email == "" {
		http.Error(w, "A graph name and email is required", http.StatusBadRequest)
		return }
	
	graph, exists := inMemoryStore[GraphName]
	if !exists {
		http.Error(w, "Graph does not exist...", http.StatusNotFound)
		return }

	for _, node := range graph.Nodes {
		if node.Email == email {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(node)
			return } }

	http.Error(w, "Node not found...", http.StatusNotFound)
}
