package test

import (
	"JOOJ-Graph/backend/model"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestEdgeSingleFields(t *testing.T) {

	John := model.User_node{
		User_id:             "19annahdksnHKAnskjs01192",
		First_name:          "John",
		Last_name:           "Doe",
		Profile_picture_url: "Placeholder1",
		Email:               "JohnDoe@email.com",
		Created_at:          "24/08/2026",
	}
	Jane := model.User_node{
		User_id:             "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA",
		First_name:          "Jane",
		Last_name:           "Smith",
		Profile_picture_url: "Placeholder2",
		Email:               "JaneSmith@email.com",
		Created_at:          "03/12/2026",
	}
	edgeAB := model.Edge{
		Source:           John,
		Target:           Jane,
		Edge_tag:         "Family",
		Edge_desc:        "Desc AB",
		Created_at_edges: "20/08/26",
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
	if edgeAB.Created_at_edges != "20/08/26" {
		t.Errorf("Expected creation date: 20/08/26, but got %s", edgeAB.Created_at_edges)
	}
	if edgeAB.Source != John || edgeAB.Target != Jane {
		t.Errorf("Expected edge direction from John to Jane, but got %s to %s", edgeAB.Source, edgeAB.Target)
	}
}

func TestEdgeMultiplefields(t *testing.T) {
	John := model.User_node{
		User_id:             "19annahdksnHKAnskjs01192",
		First_name:          "John",
		Last_name:           "Doe",
		Profile_picture_url: "Placeholder1",
		Email:               "JohnDoe@email.com",
		Created_at:          "24/08/2026",
	}
	Jane := model.User_node{
		User_id:             "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA",
		First_name:          "Jane",
		Last_name:           "Smith",
		Profile_picture_url: "Placeholder2",
		Email:               "JaneSmith@email.com",
		Created_at:          "03/12/2026",
	}
	Joe := model.User_node{
		User_id:             "JHnHJy8JBcrYI798NVdj0dn2KOSB9",
		First_name:          "Joe",
		Last_name:           "Jo",
		Profile_picture_url: "Placeholder3",
		Email:               "JoeJo@email.com",
		Created_at:          "02/01/2027",
	}
	edgeAB := model.Edge{Source: John, Target: Jane, Edge_tag: "Family", Edge_desc: "Desc AB", Created_at_edges: "20/08/26"}
	edgeBA := model.Edge{Source: Jane, Target: John, Edge_tag: "Family", Edge_desc: "Desc BA", Created_at_edges: "20/08/26"}
	edgeAC := model.Edge{Source: John, Target: Joe, Edge_tag: "Friend", Edge_desc: "Desc AC", Created_at_edges: "11/02/27"}
	edgeBC := model.Edge{Source: Jane, Target: Joe, Edge_tag: "Friend", Edge_desc: "Desc BC", Created_at_edges: "21/09/27"}

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
		User_id:             "19annahdksnHKAnskjs01192",
		First_name:          "John",
		Last_name:           "Doe",
		Profile_picture_url: "Placeholder1",
		Email:               "JohnDoe@email.com",
		Created_at:          "24/08/2026",
	}
	Jane := model.User_node{
		User_id:             "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA",
		First_name:          "Jane",
		Last_name:           "Smith",
		Profile_picture_url: "Placeholder2",
		Email:               "JaneSmith@email.com",
		Created_at:          "03/12/2026",
	}
	edges := []model.Edge{
		{Source: John, Target: Jane, Edge_tag: "Family", Edge_desc: "Desc AB", Created_at_edges: "20/08/26"},
		{Source: Jane, Target: John, Edge_tag: "Family", Edge_desc: "Desc BA", Created_at_edges: "20/08/26"}}

	new_edge := model.Edge{Source: John, Target: Jane, Edge_tag: "Family", Edge_desc: "Desc AB", Created_at_edges: "20/08/26"}

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
		User_id:             "19annahdksnHKAnskjs01192",
		First_name:          "John",
		Last_name:           "Doe",
		Profile_picture_url: "Placeholder1",
		Email:               "JohnDoe@email.com",
		Created_at:          "24/08/2026",
	}
	Jane := model.User_node{
		User_id:             "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA",
		First_name:          "Jane",
		Last_name:           "Smith",
		Profile_picture_url: "Placeholder2",
		Email:               "JaneSmith@email.com",
		Created_at:          "03/12/2026",
	}
	edges := []model.Edge{
		{Source: John, Target: Jane, Edge_tag: "Family", Edge_desc: "Desc AB", Created_at_edges: "20/08/26"},
		{Source: Jane, Target: John, Edge_tag: "Family", Edge_desc: "Desc BA", Created_at_edges: "20/08/26"},
		{Source: Jane, Target: Jane, Edge_tag: "NA", Edge_desc: "NA", Created_at_edges: "NA"}}

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
		User_id:             "19annahdksnHKAnskjs01192",
		First_name:          "John",
		Last_name:           "Doe",
		Profile_picture_url: "Placeholder1",
		Email:               "JohnDoe@email.com",
		Created_at:          "24/08/2026",
	}
	Jane := model.User_node{
		User_id:             "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA",
		First_name:          "Jane",
		Last_name:           "Smith",
		Profile_picture_url: "Placeholder2",
		Email:               "JaneSmith@email.com",
		Created_at:          "03/12/2026",
	}
	Joe := model.User_node{
		User_id:             "JHnHJy8JBcrYI798NVdj0dn2KOSB9",
		First_name:          "Joe",
		Last_name:           "Jo",
		Profile_picture_url: "Placeholder3",
		Email:               "JoeJo@email.com",
		Created_at:          "02/01/2027",
	}
	edges := []model.Edge{
		{Source: John, Target: Jane, Edge_tag: "Family", Edge_desc: "Desc AB", Created_at_edges: "20/08/26"},
		{Source: John, Target: model.User_node{}, Edge_tag: "Family", Edge_desc: "Desc BA", Created_at_edges: "20/08/26"},
		{Source: Jane, Target: Joe, Edge_tag: "Family", Edge_desc: "Desc BC", Created_at_edges: "02/12/26"},
		{Source: John, Target: Joe, Edge_tag: "Friend", Edge_desc: "Desc AC", Created_at_edges: "11/02/27"}}

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
		User_id:             "19annahdksnHKAnskjs01192",
		First_name:          "John",
		Last_name:           "Doe",
		Profile_picture_url: "Placeholder1",
		Email:               "JohnDoe@email.com",
		Created_at:          "24/08/2026",
	}
	Jane := model.User_node{
		User_id:             "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA",
		First_name:          "Jane",
		Last_name:           "Smith",
		Profile_picture_url: "Placeholder2",
		Email:               "JaneSmith@email.com",
		Created_at:          "03/12/2026",
	}
	Joe := model.User_node{
		User_id:             "JHnHJy8JBcrYI798NVdj0dn2KOSB9",
		First_name:          "Joe",
		Last_name:           "Jo",
		Profile_picture_url: "Placeholder3",
		Email:               "JoeJo@email.com",
		Created_at:          "02/01/2027",
	}
	edgeAB := model.Edge{Source: John, Target: Jane, Edge_tag: "Family", Edge_desc: "Desc AB", Created_at_edges: "20/08/26"}
	edgeAC := model.Edge{Source: John, Target: Joe, Edge_tag: "Friend", Edge_desc: "Desc AC", Created_at_edges: "11/02/27"}

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
		User_id:             "19annahdksnHKAnskjs01192",
		First_name:          "John",
		Last_name:           "Doe",
		Profile_picture_url: "Placeholder1",
		Email:               "JohnDoe@email.com",
		Created_at:          "24/08/2026",
	}
	Jane := model.User_node{
		User_id:             "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA",
		First_name:          "Jane",
		Last_name:           "Smith",
		Profile_picture_url: "Placeholder2",
		Email:               "JaneSmith@email.com",
		Created_at:          "03/12/2026",
	}
	Joe := model.User_node{
		User_id:             "JHnHJy8JBcrYI798NVdj0dn2KOSB9",
		First_name:          "Joe",
		Last_name:           "Jo",
		Profile_picture_url: "Placeholder3",
		Email:               "JoeJo@email.com",
		Created_at:          "02/01/2027",
	}
	Edges := []model.Edge{
		{Source: John, Target: Jane, Edge_tag: "Family", Edge_desc: "Desc AB", Created_at_edges: "20/08/26"},
		{Source: Jane, Target: John, Edge_tag: "Family", Edge_desc: "Desc BA", Created_at_edges: "20/08/26"},
		{Source: John, Target: Joe, Edge_tag: "Friend", Edge_desc: "Desc AC", Created_at_edges: "11/02/27"}}

	if len(Edges) != 3 {
		t.Errorf("Expected 3 edges but got %v", len(Edges))
	}
	if Edges[0].Edge_tag != "Family" {
		t.Errorf("Expected Family but got %s", Edges[0].Edge_tag)
	}
	if Edges[1].Created_at_edges != "20/08/26" {
		t.Errorf("Expected 20/08/26 but got %s", Edges[1].Created_at_edges)
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
Created_at_edges: date of start of relationship 
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
