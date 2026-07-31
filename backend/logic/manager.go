package logic

import (
	// Internal Imports
	"JOOJ-Graph/backend/model"

	// External Imports
	"errors"
	"regexp"
)

type GraphManager interface{
	CreateGraph(Name string)
	GetGraph(Name string) model.Graph
}

type JoojGraphManager struct{
	InMemoryStore map[string]model.Graph 
}

func (j JoojGraphManager) CreateGraph(Name string) error {
	Graph := model.Graph{
		Name : Name,
	}
	NonWhtSpc_Field := regexp.MustCompile(`^\s+$`)


	if Name == "" || NonWhtSpc_Field.MatchString(Name) {
		return errors.New("Graph Name should not be empty, please use another name...")
	} else if _, exists := j.InMemoryStore[Name]; exists {
		return errors.New("Graph with this name already exists, please use another name...")
	} else {
		j.InMemoryStore[Name] = Graph
		return nil
	}
}

func (j JoojGraphManager) GetGraph(Name string) model.Graph{
	return j.InMemoryStore[Name]
}