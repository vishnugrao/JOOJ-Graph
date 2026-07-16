package API

import (
	"JOOJ-Graph/backend/model"
	"encoding/json"
	"errors"
	// "fmt"
	"net/http"
)

func CreateGraph(Name string) (model.Graph, error) {
	if Name == "" || valid_field.MatchString(Name) { return model.Graph{}, errors.New("Invalid Name, please try again...") }
	graph := model.Graph{}
	graph.Name = Name
	graph.Nodes = []model.User_node{}
	graph.Edges = []model.Edge{}
	graph.NodeMap = make(map[string]bool)
	graph.EdgeMap = make(map[string]bool)
	return graph, nil
}

func CreateGraphHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method...", http.StatusBadRequest)
	} else {
		var namereq struct { Name string `json:"name"` }
		if err := json.NewDecoder(r.Body).Decode(&namereq); err != nil { http.Error(w, "Invalid JSON...", http.StatusBadRequest)
			return }
		if _, exists := inMemoryStore[namereq.Name]; exists { http.Error(w, "Graph with that name already exists", http.StatusConflict)
			return }

		graph, err := CreateGraph(namereq.Name)
		if err != nil { http.Error(w, err.Error(), http.StatusBadRequest)
			return }

		inMemoryStore[graph.Name] = graph

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(graph)
	}
}

func GetGraphHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid Method...", http.StatusBadRequest)
	} else {
		name := r.URL.Query().Get("name")
		if name == "" { http.Error(w, "Invalid Graph name, try again...", http.StatusBadRequest)
			return }

		graph, exists := inMemoryStore[name]

		if !exists { http.Error(w, "Graph does not exists, try again...", http.StatusNotFound)
			return }

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(graph)
	}
}

func GetAllGraphHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid Method...", http.StatusBadRequest)
		return } 
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inMemoryStore)
}

func GraphAddNode(graph *model.Graph, node model.User_node) (model.Graph, error) {
	if !graph.NodeMap[node.User_id] {
		graph.Nodes = append(graph.Nodes, node)
		graph.NodeMap[node.User_id] = true
		inMemoryStore[graph.Name] = *graph
		return *graph, nil
	} else {
		return model.Graph{}, errors.New("Node already exists...")
	}
}

func GraphAddNodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method...", http.StatusBadRequest)
	} else {
		var req struct { 
			GraphName string `json:"graphname"`
			NodeObj model.User_node `json:"node"` }
			
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil { http.Error(w, "Invalid JSON...", http.StatusBadRequest)
			return }
		
		graph, exists := inMemoryStore[req.GraphName]
		if !exists {
			http.Error(w, "Graph does not exist...", http.StatusNotFound)
			return }

		for _, n := range graph.Nodes {
		if n.Email == req.NodeObj.Email {
			http.Error(w, "A node with this email already exists...", http.StatusBadRequest)
			return } }

		updatedGraph, err := GraphAddNode(&graph, req.NodeObj)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return }

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(updatedGraph)
	}
}


// func GraphRemoveNode(graph *model.Graph, removeid string) (model.Graph, error) {
// 	if graph.NodeMap[removeid] {
// 		for i, node := range graph.Nodes {
// 			if node.User_id == removeid {
// 				graph.Nodes = append(graph.Nodes[:i], graph.Nodes[i+1:]...)
// 				delete(graph.NodeMap, removeid)
// 				break
// 			}
// 		}

// 		for i := len(graph.Edges) - 1; i >= 0; i-- {
// 			if graph.Edges[i].Source.User_id == removeid || graph.Edges[i].Target.User_id == removeid {
// 				delete(graph.EdgeMap, graph.Edges[i].Source.User_id+"-"+graph.Edges[i].Target.User_id)
// 				delete(graph.EdgeMap, graph.Edges[i].Target.User_id+"-"+graph.Edges[i].Source.User_id)
// 				graph.Edges = append(graph.Edges[:i], graph.Edges[i+1:]...)
// 			}
// 		}
// 		return *graph, nil
// 	} else {
// 		return model.Graph{}, errors.New("Node does not exist...")
// 	}

// }

// func GraphAddEdge(graph *model.Graph, Source model.User_node, Target model.User_node, Edge_tag string, Edge_desc string, Created_at_edges string) (model.Graph, error) {
// 	if graph.NodeMap[Source.User_id] && graph.NodeMap[Target.User_id] {
// 		edge_key := Source.User_id + "-" + Target.User_id
// 		if !graph.EdgeMap[edge_key] {
// 			edge, err := CreateUserEdge(Source, Target, Edge_tag, Edge_desc, Created_at_edges)
// 			if err != nil {
// 				return model.Graph{}, err
// 			}
// 			graph.Edges = append(graph.Edges, edge)
// 			graph.EdgeMap[edge_key] = true
// 			return *graph, nil
// 		} else {
// 			return model.Graph{}, errors.New("Edge already exists")
// 		}

// 	} else {
// 		return model.Graph{}, errors.New("Source or Target node do not exist...")
// 	}

// }

// func GraphRemoveEdge(graph *model.Graph, removeSource model.User_node, removeTarget model.User_node) (model.Graph, error) {
// 	removeID1 := removeSource.User_id
// 	removeID2 := removeTarget.User_id

// 	removekey := removeID1 + "-" + removeID2

// 	if graph.EdgeMap[removekey] {
// 		for i := len(graph.Edges) - 1; i >= 0; i-- {
// 			if graph.Edges[i].Source.User_id == removeID1 && graph.Edges[i].Target.User_id == removeID2 {
// 				delete(graph.EdgeMap, graph.Edges[i].Source.User_id+"-"+graph.Edges[i].Target.User_id)
// 				graph.Edges = append(graph.Edges[:i], graph.Edges[i+1:]...)
// 			}
// 		}
// 		return *graph, nil
// 	} else {
// 		return model.Graph{}, errors.New("Edge does not exist")
// 	}
// }