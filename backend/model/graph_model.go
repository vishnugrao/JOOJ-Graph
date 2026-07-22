package model

type Graph struct{
	Name string
}

type GraphManager interface{
	CreateGraph(Name string)
}

type JoojGraphManager struct{
	InMemoryStore map[string]Graph 
}

func (j JoojGraphManager) CreateGraph(Name string){
	//  Todo
}