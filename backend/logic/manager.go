package logic

import (
	// Internal Imports
	"JOOJ-Graph/backend/model"

	// External Imports
	"errors"
	"regexp"
	"github.com/google/uuid"
)

var (
	ValidName_Field  = regexp.MustCompile(`^[a-zA-Z\s\-']+$`)
	NonWhtSpc_Field = regexp.MustCompile(`^\s+$`)
	ValidEmail_Field = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
)

type GraphManager interface {
	CreateGraph(Name string)
	GetGraph(Name string) model.Graph
}

type NodeManager interface {
	CreateNode(First_name string, Last_name string, Profile_picture_url string, Email string) model.User_node
}

type JoojGraphManager struct {
	InMemoryStore map[string]model.Graph 
}

func (j JoojGraphManager) CreateGraph(Name string) error {
	Graph := model.Graph{
		Name : Name,
	}
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

func CreateNode(First_name string, Last_name string, Profile_picture_url string, Email string) (model.User_node, error) {
	Node := model.User_node{
		User_id: uuid.NewString(),
		First_name : First_name,
		Last_name: Last_name,
		Profile_picture_url: Profile_picture_url,
		Email: Email,
	}

	if First_name == "" || Last_name == "" || Email == "" {
		return model.User_node{}, errors.New("Node first name, last name and email should not be empty...")
	} else if NonWhtSpc_Field.MatchString(First_name) || NonWhtSpc_Field.MatchString(Last_name) || NonWhtSpc_Field.MatchString(Profile_picture_url) || NonWhtSpc_Field.MatchString(Email) {
		return model.User_node{}, errors.New("All node fields should not contain whitespace...")
	} else if !ValidName_Field.MatchString(First_name) || !ValidName_Field.MatchString(Last_name) {
		return model.User_node{}, errors.New("Node first name, last name fields are invalid, please try again...")
	} else if !ValidEmail_Field.MatchString(Email) {
		return model.User_node{}, errors.New("Node email is invalid, please try again...")
	} else {
		return Node, nil
	}
}