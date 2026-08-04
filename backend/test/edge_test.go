package test

import (
	// Internal Imports
	"JOOJ-Graph/backend/model"
	"JOOJ-Graph/backend/logic"

	// External Imports
	"testing"
	// "gopkg.in/yaml.v3"
)

var (
	John = model.User_node{
		User_id:             "19annahdksnHKAnskjs01192",
		First_name:          "John",
		Last_name:           "Doe",
		Profile_picture_url: "Placeholder1",
		Email:               "JohnDoe@email.com",
	}
	Jane = model.User_node{
		User_id:             "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA",
		First_name:          "Jane",
		Last_name:           "Smith",
		Profile_picture_url: "Placeholder2",
		Email:               "JaneSmith@email.com",
	}
	Joe = model.User_node{
		User_id:             "JHnHJy8JBcrYI798NVdj0dn2KOSB9",
		First_name:          "Joe",
		Last_name:           "Jo",
		Profile_picture_url: "Placeholder3",
		Email:               "JoeJo@email.com",
	}

	edges = []model.Edge{
		{Source: John, Target: Jane, Edge_tag: "Family", Edge_desc: "Desc AB"},
		{Source: Jane, Target: John, Edge_tag: "Family", Edge_desc: "Desc BA"}}
)

func TestEdgeInvalidFields(t *testing.T) {
	_, exists := logic.CreateEdge(Jane, John, "", "Highschool")
	if exists == nil {
		t.Errorf("Edge Tag should not be empty.")
	}

	_, exists2 := logic.CreateEdge(John, Jane, " ", "Highschool")
	if exists2 == nil {
		t.Errorf("Edge Tag should not be whitespace.")
	}

	_, exists3 := logic.CreateEdge(Joe, John, "Friend", "")
	if exists3 == nil {
		t.Errorf("Edge description should not be empty.")
	}

	_, exists4 := logic.CreateEdge(John, Joe, "Friend", " ")
	if exists4 == nil {
		t.Errorf("Edge description should not be whitespace.")
	}

	_, exists5 := logic.CreateEdge(John, Joe, "Fr1end?", "From Highschool")
	if exists5 == nil {
		t.Errorf("Edge tag should only contain alphabet letters and forward slash.")
	}
}

func TestEdgeSelfEdgeDetection(t *testing.T) {
	_, exists := logic.CreateEdge(John, John, "Myself", "Became concious.")
	if exists == nil {
		t.Errorf("The source and target cannot be the same node.")
	}
}

func TestEdgeMissingSourceTargetDetection(t *testing.T) {
	_, exists := logic.CreateEdge(model.User_node{}, Jane, "Friend", "From Highschool")
	if exists == nil {
		t.Errorf("Expected no edge to be created but got an edge.")
	}
}
// SAME ARE WE STILL USING THIS?

// func TestEdgeYamlUnmarshall(t *testing.T) {
// 	yamlData := []byte(`
// Source:
//   First_name: John
// Target:
//   First_name: Jane
// Edge_tag: relationship tag
// Edge_desc: Desc of relationship
// Created_at_edges: date of start of relationship 
// `)

// 	var edge1 model.Edge
// 	err := yaml.Unmarshal(yamlData, &edge1)

// 	if err != nil {
// 		t.Errorf("Expected no error, got %v", err)
// 	}
// 	if edge1.Target.First_name != "Jane" {
// 		t.Errorf("Expected End Node but got %s", edge1.Target)
// 	}
// }
