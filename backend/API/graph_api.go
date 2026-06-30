package api

import (
	"errors"
	"github.com/abhay/JOOJ-Graph/backend/api"
	"github.com/abhay/JOOJ-Graph/backend/model"
)

func CreateGraph() model.Graph{
	graph := model.Graph{}
	graph.NodeMap = make(map[string]bool)
	graph.EdgeMap = make(map[string]bool)
	return graph
}

func GraphAddNode (graph *model.Graph, User_id string, First_name string, Last_name string, Profile_picture_url string, Email string, Created_at string) (model.Graph, error){
	if !graph.NodeMap[User_id] {
	node, err := api.CreateUserNode(User_id, First_name, Last_name, Profile_picture_url, Email, Created_at)
	if err != nil {	return model.Graph{}, err}
	graph.Nodes = append(graph.Nodes, node)
	graph.NodeMap[User_id] = true
	return *graph, nil
	} else {
		return model.Graph{}, errors.New("Node already exists...")
	}
}

func GraphRemoveNode (graph *model.Graph, removeid string) (model.Graph, error) {
	if graph.NodeMap[removeid] {
		for i, node := range graph.Nodes {
			if node.User_id == removeid {
				graph.Nodes = append(graph.Nodes[:i], graph.Nodes[i+1:]...)
				delete(graph.NodeMap, removeid)
				break
			}
		}

		for i:= len(graph.Edges) -1; i>=0; i-- {
			if graph.Edges[i].Source.User_id == removeid || graph.Edges[i].Target.User_id == removeid {
				delete(graph.EdgeMap, graph.Edges[i].Source.User_id + "-" + graph.Edges[i].Target.User_id)
				delete(graph.EdgeMap, graph.Edges[i].Target.User_id + "-" + graph.Edges[i].Source.User_id)
				graph.Edges = append(graph.Edges[:i], graph.Edges[i+1:]...)
			}
		}
		return *graph, nil
	} else {
		return model.Graph{}, errors.New("Node does not exist...")
	}

}

func GraphAddEdge(graph *model.Graph, Source model.User_node, Target model.User_node, Edge_tag string, Edge_desc string, Created_at_edges string) (model.Graph, error) {
	if graph.NodeMap[Source.User_id] && graph.NodeMap[Target.User_id]  {
		edge_key := Source.User_id + "-" + Target.User_id
		if !graph.EdgeMap[edge_key] {
			edge, err := api.CreateUserEdge(Source, Target, Edge_tag, Edge_desc, Created_at_edges)
			if err != nil {	return model.Graph{}, err}
			graph.Edges = append(graph.Edges, edge)
			graph.EdgeMap[edge_key] = true
			return *graph, nil
		} else {
			return model.Graph{}, errors.New("Edge already exists")
		}

	} else {
		return model.Graph{}, errors.New("Source or Target node do not exist...")
	}

}


func GraphRemoveEdge(graph *model.Graph, removeSource model.User_node, removeTarget model.User_node) (model.Graph, error) {
	removeID1 := removeSource.User_id
	removeID2 := removeTarget.User_id

	removekey := removeID1 + "-" + removeID2

	if graph.EdgeMap[removekey] {
		for i:= len(graph.Edges) -1; i>=0; i-- {
			if graph.Edges[i].Source.User_id == removeID1 && graph.Edges[i].Target.User_id == removeID2 {
				delete(graph.EdgeMap, graph.Edges[i].Source.User_id + "-" + graph.Edges[i].Target.User_id)
				graph.Edges = append(graph.Edges[:i], graph.Edges[i+1:]...)
			}
		}
		return *graph, nil
	} else {
		return model.Graph{}, errors.New("Edge does not exist")
	}
}