package logic

import (
	// Internal Imports
	"JOOJ-Graph/backend/model"
	"fmt"
)

type GraphManager interface{
	CreateGraph(Name string)
	GetGraph(Name string) model.Graph
}

type JoojGraphManager struct{
	InMemoryStore map[string]model.Graph 
}

func (j JoojGraphManager) CreateGraph(Name string){
	Graph := model.Graph{
		Name : Name,
	}
	// if Name == "" || Name == "" // whitespace check{fmt.Println("Invalid Graph name, please use another name...")}
	if _, exists := j.InMemoryStore[Name]; exists { // change this to else if after top line is written...
		fmt.Println("Graph with this name already exists, please use another name...")
	} else {
		j.InMemoryStore[Name] = Graph
	}
}

func (j JoojGraphManager) GetGraph(Name string) model.Graph{
	return j.InMemoryStore[Name]
}