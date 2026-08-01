package test

import (
	// Internal Imports
	"JOOJ-Graph/backend/model"
	"JOOJ-Graph/backend/logic"

	// External Imports
	"regexp"
	"testing"
	// "gopkg.in/yaml.v3"
)

var (
	valid_name  = regexp.MustCompile(`^[a-zA-Z\s\-']+$`)
	valid_field = regexp.MustCompile(`^\s+$`)
	valid_email = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	user_nodes = []model.User_node{
		{User_id: "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA", First_name: "Jane", Last_name: "Smith", Profile_picture_url: "Placeholder2",
			Email: "JaneSmith@email.com"},
		{User_id: "JHnHJy8JBcrYI798NVdj0dn2KOSB9", First_name: "Joe", Last_name: "Jo", Profile_picture_url: "Placeholder3",
			Email: "JoeJo@email.com"},
	}
)


func TestNodeUserEmptyFields(t *testing.T) {
	_ , exists := logic.CreateNode("", "", "", "")
	if exists == nil {
		t.Errorf("Expected node fields to not be empty.")
	}
}

func TestNodeNoDuplicateUserNodes(t *testing.T) {
	duplcte_user_node, _ := logic.CreateNode("jane", "smith", "placeholder2", "JaneSmith@email.com")
	duplicate_found := false
	for _, user_node := range user_nodes {
		if user_node.Email == duplcte_user_node.Email {
			duplicate_found = true
			break
		}
	}
	if !duplicate_found {
		t.Errorf("Expected a duplicate user node to be found but it was not found.")
	}
}

func TestNodeInvalidEmail(t *testing.T) {
	_ , exists := logic.CreateNode("Joe", "Joe", "url", "JoeJoemail.com");
	if exists == nil{
		t.Errorf("Expected node email to be correctly formatted.")
		}
}

func TestNodeInvalidName(t *testing.T) {
	_ , exists := logic.CreateNode("J0e", "Joe", "url", "JoeJoemail.com")
	if exists == nil {
		t.Errorf("Expected node name to be correctly formatted.")
	}
}

func TestNodeWhiteSpaceOnly(t *testing.T) {
	_, exists := logic.CreateNode(" ", "Joe", "url", "JoeJoemail.com")
		if exists == nil {
			t.Errorf("Expected the node fields to not contain whitespace.")
		}
}

// func TestNodeUserNodeIndependency(t *testing.T) {
// 	John := model.User_node{User_id: "19annahdksnHKAnskjs01192", First_name: "John", Last_name: "Doe",
// 		Profile_picture_url: "Placeholder1", Email: "JohnDoe@email.com", Created_at: "24/08/2026"}
// 	Jane := model.User_node{
// 		User_id: "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA", First_name: "Jane", Last_name: "Smith",
// 		Profile_picture_url: "Placeholder2", Email: "JaneSmith@email.com", Created_at: "03/12/2026"}

// 	John.Email = "JohnDoe123@email.com"

// 	if Jane.Email == "JohnDoe123@email.com" {
// 		t.Errorf("Expected Jane to be unaffected but got %s", Jane.Email)
// 	}
// 	if John.Email != "JohnDoe123@email.com" {
// 		t.Errorf("Expected JohnDoe123@email.com but got %s", John.Email)
// 	}
// }

// ------ This function tests if we the YAML can deserialize a Node struct into a readabel struct, ARE WE STILL USING YAML OR JSON?

// func TestNodeUserYamlUnmarshall(t *testing.T) {
// 	yamlData := []byte(`
// User_id: unique_identifier
// First_name: name1
// Last_name: name2
// Profile_picture_url: link
// Email: email_address
// Created_at: when_they_joined
// `)

// 	var user_node model.User_node
// 	err := yaml.Unmarshal(yamlData, &user_node)

// 	if err != nil {
// 		t.Errorf("Expected no error, got %v", err)
// 	}
// 	if user_node.User_id != "unique_identifier" {
// 		t.Errorf("Expected unique_identifier but got %s", user_node.User_id)
// 	}
// }