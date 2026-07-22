package logic

import(
	// Internal Imports
	"JOOJ-Graph/backend/model"
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
	j.InMemoryStore[Name] = Graph
}

func (j JoojGraphManager) GetGraph(Name string) model.Graph{
	return j.InMemoryStore[Name]
}