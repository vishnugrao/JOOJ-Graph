package test

import (
// Internal imports
	"JOOJ-Graph/backend/model"

// External imports
	"testing"
	"gopkg.in/yaml.v3"
)

func TestEdgeSingleFields(t *testing.T) {

	John := model.User_node{
		User_id: "19annahdksnHKAnskjs01192",
		First_name: "John",
		Last_name: "Doe",
		Profile_picture_url: "Placeholder1",
		Email: "JohnDoe@email.com",
	}
	Jane := model.User_node{
		User_id: "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA",
		First_name: "Jane",
		Last_name: "Smith",
		Profile_picture_url: "Placeholder2",
		Email: "JaneSmith@email.com",
	}
	edgeAB := model.Edge{
		Source: John,
		Target: Jane,
		Edge_tag: "Family",
		Edge_desc: "Desc AB",
	}

	if edgeAB.Source != John {
		t.Errorf("Expected source: John, but got %s", edgeAB.Source)
	}
	if edgeAB.Target != Jane {
		t.Errorf("Expected target: Jane, but got %s", edgeAB.Target)
	}
	if edgeAB.Edge_tag != "Family" {
		t.Errorf("Expected edge tag: Family, but got %s", edgeAB.Edge_tag)
	}
	if edgeAB.Edge_desc != "Desc AB" {
		t.Errorf("Expected edge description: Desc AB, but got %s", edgeAB.Edge_desc)
	}
	if edgeAB.Source != John || edgeAB.Target != Jane {
		t.Errorf("Expected edge direction from John to Jane, but got %s to %s", edgeAB.Source, edgeAB.Target)
	}
}

func TestEdgeMultiplefields(t *testing.T) {
	John := model.User_node{
		User_id: "19annahdksnHKAnskjs01192",
		First_name: "John",
		Last_name: "Doe",
		Profile_picture_url: "Placeholder1",
		Email: "JohnDoe@email.com",
	}
	Jane := model.User_node{
		User_id: "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA",
		First_name: "Jane",
		Last_name: "Smith",
		Profile_picture_url: "Placeholder2",
		Email: "JaneSmith@email.com",
	}
	Joe := model.User_node{
		User_id: "JHnHJy8JBcrYI798NVdj0dn2KOSB9",
		First_name: "Joe",
		Last_name: "Jo",
		Profile_picture_url: "Placeholder3",
		Email: "JoeJo@email.com",
	}
	edgeAB := model.Edge{Source: John, Target: Jane, Edge_tag: "Family", Edge_desc: "Desc AB"}
	edgeBA := model.Edge{Source: Jane, Target: John, Edge_tag: "Family", Edge_desc: "Desc BA"}
	edgeAC := model.Edge{Source: John, Target: Joe, Edge_tag: "Friend", Edge_desc: "Desc AC"}
	edgeBC := model.Edge{Source: Jane, Target: Joe, Edge_tag: "Friend", Edge_desc: "Desc BC"}

	if edgeAB.Edge_tag != "Family" {
		t.Errorf("Expected edge tag: Family, but got %s", edgeAB.Edge_tag)
	}
	if edgeAB.Edge_desc != "Desc AB" {
		t.Errorf("Expected edge description: Desc AB, but got %s", edgeAB.Edge_desc)
	}
	if edgeBA.Edge_tag != "Family" {
		t.Errorf("Expected edge tag: Family, but got %s", edgeBA.Edge_tag)
	}
	if edgeBA.Edge_desc != "Desc BA" {
		t.Errorf("Expected edge description: Desc BA, but got %s", edgeBA.Edge_desc)
	}
	if edgeAC.Edge_tag != "Friend" {
		t.Errorf("Expected edge tag: Friend, but got %s", edgeAC.Edge_tag)
	}
	if edgeAC.Edge_desc != "Desc AC" {
		t.Errorf("Expected edge description: Desc AC, but got %s", edgeAC.Edge_desc)
	}
	if edgeBC.Edge_tag != "Friend" {
		t.Errorf("Expected edge tag: Friend, but got %s", edgeBC.Edge_tag)
	}
	if edgeBC.Edge_desc != "Desc BC" {
		t.Errorf("Expected edge description: Desc BC, but got %s", edgeBC.Edge_desc)
	}
}

func TestEdgeDuplicateEdgeDetection(t *testing.T) {
	John := model.User_node{
		User_id: "19annahdksnHKAnskjs01192",
		First_name: "John",
		Last_name: "Doe",
		Profile_picture_url: "Placeholder1",
		Email: "JohnDoe@email.com",
	}
	Jane := model.User_node{
		User_id: "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA",
		First_name: "Jane",
		Last_name: "Smith",
		Profile_picture_url: "Placeholder2",
		Email: "JaneSmith@email.com",
	}
	edges := []model.Edge{
		{Source: John, Target: Jane, Edge_tag: "Family", Edge_desc: "Desc AB"},
		{Source: Jane, Target: John, Edge_tag: "Family", Edge_desc: "Desc BA"}}
	new_edge := model.Edge{Source: John, Target: Jane, Edge_tag: "Family", Edge_desc: "Desc AB"}

	duplicate_found := false
	for _, edge := range edges {
		if edge.Source == new_edge.Source && edge.Target == new_edge.Target {
			duplicate_found = true
			break
		}
	}
	if !duplicate_found {
		t.Errorf("Expected a duplicate edge to be found but it was not found.")
	}
}

func TestEdgeSelfEdgeDetection(t *testing.T) {
	John := model.User_node{
		User_id: "19annahdksnHKAnskjs01192",
		First_name: "John",
		Last_name: "Doe",
		Profile_picture_url: "Placeholder1",
		Email: "JohnDoe@email.com",
	}
	Jane := model.User_node{
		User_id: "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA",
		First_name: "Jane",
		Last_name: "Smith",
		Profile_picture_url: "Placeholder2",
		Email: "JaneSmith@email.com",
	}
	edges := []model.Edge{
		{Source: John, Target: Jane, Edge_tag: "Family", Edge_desc: "Desc AB"},
		{Source: Jane, Target: John, Edge_tag: "Family", Edge_desc: "Desc BA"},
		{Source: Jane, Target: Jane, Edge_tag: "NA", Edge_desc: "NA"}}

	self_edge_found := false
	for _, edge := range edges {
		if edge.Source == edge.Target {
			self_edge_found = true
			break
		}
	}
	if !self_edge_found {
		t.Errorf("Expected a self edge to be found but it was not found.")
	}
}

func TestEdgeMissingSourceTargetDetection(t *testing.T) {
	John := model.User_node{
		User_id: "19annahdksnHKAnskjs01192",
		First_name: "John",
		Last_name: "Doe",
		Profile_picture_url: "Placeholder1",
		Email: "JohnDoe@email.com",
	}
	Jane := model.User_node{
		User_id: "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA",
		First_name: "Jane",
		Last_name: "Smith",
		Profile_picture_url: "Placeholder2",
		Email: "JaneSmith@email.com",
	}
	Joe := model.User_node{
		User_id: "JHnHJy8JBcrYI798NVdj0dn2KOSB9",
		First_name: "Joe",
		Last_name: "Jo",
		Profile_picture_url: "Placeholder3",
		Email: "JoeJo@email.com",
	}
	edges := []model.Edge{
		{Source: John, Target: Jane, Edge_tag: "Family", Edge_desc: "Desc AB"},
		{Source: John, Target: model.User_node{}, Edge_tag: "Family", Edge_desc: "Desc BA"},
		{Source: Jane, Target: Joe, Edge_tag: "Family", Edge_desc: "Desc BC"},
		{Source: John, Target: Joe, Edge_tag: "Friend", Edge_desc: "Desc AC"}}

	missing_SourceOrTarget := false
	for _, edge := range edges {
		if edge.Source.User_id == "" || edge.Target.User_id == "" {
			missing_SourceOrTarget = true
			break
		}
	}
	if !missing_SourceOrTarget {
		t.Errorf("Expected a missing source or target to be found but it was not found.")
	}
}

func TestEdgeIndependency(t *testing.T) {
	John := model.User_node{
		User_id: "19annahdksnHKAnskjs01192",
		First_name: "John",
		Last_name: "Doe",
		Profile_picture_url: "Placeholder1",
		Email: "JohnDoe@email.com",
	}
	Jane := model.User_node{
		User_id: "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA",
		First_name: "Jane",
		Last_name: "Smith",
		Profile_picture_url: "Placeholder2",
		Email: "JaneSmith@email.com",
	}
	Joe := model.User_node{
		User_id: "JHnHJy8JBcrYI798NVdj0dn2KOSB9",
		First_name: "Joe",
		Last_name: "Jo",
		Profile_picture_url: "Placeholder3",
		Email: "JoeJo@email.com",
	}
	edgeAB := model.Edge{Source: John, Target: Jane, Edge_tag: "Family", Edge_desc: "Desc AB"}
	edgeAC := model.Edge{Source: John, Target: Joe, Edge_tag: "Friend", Edge_desc: "Desc AC"}

	edgeAC.Edge_tag = "Friend/Highschool"

	if edgeAB.Edge_tag == "Friend/Highschool" {
		t.Errorf("Expected edge from A -> B to be unaffected but got %s", edgeAB.Edge_tag)
	}
	if edgeAC.Edge_tag != "Friend/Highschool" {
		t.Errorf("Expected Friend/Highschool but got %s", edgeAC.Edge_tag)
	}
}

func TestEdgeSlice(t *testing.T) {
	John := model.User_node{
		User_id: "19annahdksnHKAnskjs01192",
		First_name: "John",
		Last_name: "Doe",
		Profile_picture_url: "Placeholder1",
		Email: "JohnDoe@email.com",
	}
	Jane := model.User_node{
		User_id: "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA",
		First_name: "Jane",
		Last_name: "Smith",
		Profile_picture_url: "Placeholder2",
		Email: "JaneSmith@email.com",
	}
	Joe := model.User_node{
		User_id: "JHnHJy8JBcrYI798NVdj0dn2KOSB9",
		First_name: "Joe",
		Last_name: "Jo",
		Profile_picture_url: "Placeholder3",
		Email: "JoeJo@email.com",
	}
	Edges := []model.Edge{
		{Source: John, Target: Jane, Edge_tag: "Family", Edge_desc: "Desc AB"},
		{Source: Jane, Target: John, Edge_tag: "Family", Edge_desc: "Desc BA"},
		{Source: John, Target: Joe, Edge_tag: "Friend", Edge_desc: "Desc AC"}}

	if len(Edges) != 3 {
		t.Errorf("Expected 3 edges but got %v", len(Edges))
	}
	if Edges[0].Edge_tag != "Family" {
		t.Errorf("Expected Family but got %s", Edges[0].Edge_tag)
	}
}

func TestEdgeEmptyFields(t *testing.T) {
	edgeAB := model.Edge{}

	if edgeAB.Edge_desc != "" {
		t.Errorf("Expected empty edge description but got %s", edgeAB.Edge_desc)
	}
}

func TestEdgeYamlUnmarshall(t *testing.T) {
	yamlData := []byte(`
Source:
  First_name: John
Target:
  First_name: Jane
Edge_tag: relationship tag
Edge_desc: Desc of relationship
`)

	var edge1 model.Edge
	err := yaml.Unmarshal(yamlData, &edge1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if edge1.Target.First_name != "Jane" {
		t.Errorf("Expected End Node but got %s", edge1.Target)
	}
}
