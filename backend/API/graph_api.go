package API

import (
	// Internal imports
	"JOOJ-Graph/backend/model"

	// External imports
	"encoding/json"
	"errors"
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
		return }

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

func GetGraphHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid Method...", http.StatusBadRequest)
		return }

	name := r.URL.Query().Get("name")
	if name == "" { http.Error(w, "Invalid Graph name, try again...", http.StatusBadRequest)
		return }

	graph, exists := inMemoryStore[name]

	if !exists { http.Error(w, "Graph does not exists, try again...", http.StatusNotFound)
		return }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(graph)
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
		return }

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


func GraphAddEdge(graph *model.Graph, edge model.Edge) (model.Graph, error) {
	if !graph.NodeMap[edge.Source.User_id] || !graph.NodeMap[edge.Target.User_id] {
		return model.Graph{}, errors.New("Source or Target node do not exist...")
	}
	if edge.Source.User_id == edge.Target.User_id {
		return model.Graph{}, errors.New("Source and Target cannot be the same node...")
	}

	graph.Edges = append(graph.Edges, edge)
	edge_key := edge.Source.User_id + "-" + edge.Target.User_id
	graph.EdgeMap[edge_key] = true
	inMemoryStore[graph.Name] = *graph
	return *graph, nil
}

func GraphAddEdgeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method...", http.StatusBadRequest)
		return }

	var req struct { 
		GraphName string `json:"graphname"`
		EdgeObj model.Edge `json:"edge"` }
		
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { http.Error(w, "Invalid JSON...", http.StatusBadRequest)
		return }
	
	graph, exists := inMemoryStore[req.GraphName]
	if !exists {
		http.Error(w, "Graph does not exist...", http.StatusNotFound)
		return }

	e_key := req.EdgeObj.Source.User_id + "-" + req.EdgeObj.Target.User_id

	if graph.EdgeMap[e_key] {
		http.Error(w, "The Edge already exists...", http.StatusBadRequest)
		return }

	updatedGraph, err := GraphAddEdge(&graph, req.EdgeObj)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updatedGraph)
}

// func DeleteGraph(graph *model.Graph) (model.Graph, error){
// 	if  {
// 		return model.Graph{}, errors.New("Graph does not exist")
// 	}
// }


func GraphRemoveNode(graph *model.Graph, node model.User_node) (model.Graph, error) {
	if !graph.NodeMap[node.User_id] {
		return model.Graph{}, errors.New("Node does not exist...")
	}
	for i, n := range graph.Nodes {
		if n.User_id == node.User_id {
			graph.Nodes = append(graph.Nodes[:i], graph.Nodes[i+1:]...)
			delete(graph.NodeMap, node.User_id)
			break
		}
	}

	for i := len(graph.Edges) - 1; i >= 0; i-- {
		if graph.Edges[i].Source.User_id == node.User_id || graph.Edges[i].Target.User_id == node.User_id {
			delete(graph.EdgeMap, graph.Edges[i].Source.User_id+"-"+graph.Edges[i].Target.User_id)
			delete(graph.EdgeMap, graph.Edges[i].Target.User_id+"-"+graph.Edges[i].Source.User_id)
			graph.Edges = append(graph.Edges[:i], graph.Edges[i+1:]...)
		}
	}
	inMemoryStore[graph.Name] = *graph
	return *graph, nil
} 

func GraphRemoveNodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method...", http.StatusBadRequest)
		return }
	var req struct { 
		GraphName string `json:"graphname"`
		NodeObj model.User_node `json:"node"` }
		
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { http.Error(w, "Invalid JSON...", http.StatusBadRequest)
		return }
	
	graph, exists := inMemoryStore[req.GraphName]
	if !exists {
		http.Error(w, "Graph does not exist...", http.StatusNotFound)
		return }
	
	found := false
	for _, n := range graph.Nodes {
	if n.Email == req.NodeObj.Email {
		req.NodeObj = n
		found = true
		break } }

	if !found {
	http.Error(w, "A node with this email does not exist...", http.StatusNotFound)
	return }

	updatedGraph, err := GraphRemoveNode(&graph, req.NodeObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedGraph)
}


func GraphRemoveEdge(graph *model.Graph, edge model.Edge) (model.Graph, error) {
	removeID1 := edge.Source.User_id
	removeID2 := edge.Target.User_id
	removekey := removeID1 + "-" + removeID2

	if !graph.EdgeMap[removekey] {
		return model.Graph{}, errors.New("Edge does not exist")
	}

	for i := len(graph.Edges) - 1; i >= 0; i-- {
		if graph.Edges[i].Source.User_id == removeID1 && graph.Edges[i].Target.User_id == removeID2 {
			delete(graph.EdgeMap, graph.Edges[i].Source.User_id+"-"+graph.Edges[i].Target.User_id)
			graph.Edges = append(graph.Edges[:i], graph.Edges[i+1:]...)
		}
	}
	inMemoryStore[graph.Name] = *graph
	return *graph, nil
}

func GraphRemoveEdgeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method...", http.StatusBadRequest)
		return }
	var req struct { 
		GraphName string `json:"graphname"`
		EdgeObj model.Edge `json:"edge"` }
		
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { http.Error(w, "Invalid JSON...", http.StatusBadRequest)
		return }
	
	graph, exists := inMemoryStore[req.GraphName]
	if !exists {
		http.Error(w, "Graph does not exist...", http.StatusNotFound)
		return }

	if req.EdgeObj.Source.User_id == "" || req.EdgeObj.Target.User_id == "" {
		http.Error(w, "Edge must have a Source and a Target...", http.StatusBadRequest)
		return }
	
	if req.EdgeObj.Source.User_id == req.EdgeObj.Target.User_id{
		http.Error(w, "The source and target node cannot be the same...", http.StatusBadRequest)
		return }

	if !graph.NodeMap[req.EdgeObj.Source.User_id] || !graph.NodeMap[req.EdgeObj.Target.User_id] {
		http.Error(w, "Source or Target node do not exist in this graph", http.StatusNotFound)
		return }

	edge_key := req.EdgeObj.Source.User_id + "-" + req.EdgeObj.Target.User_id
	if !graph.EdgeMap[edge_key]{
		http.Error(w, "Edge does not exist in this graph", http.StatusNotFound)
		return }
	
	updatedGraph, err := GraphRemoveEdge(&graph, req.EdgeObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return }
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedGraph)
}

func GetNodes(graph model.Graph) []model.User_node {
	// todo
	list := graph.Nodes
	return list
}