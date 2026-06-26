package test

import (
	"testing"
	"gopkg.in/yaml.v3"
	"github.com/abhay/JOOJ-Graph/backend/model"
	"regexp"
)

var (
	valid_name = regexp.MustCompile(`^[a-zA-Z\s\-']+$`)
	valid_field = regexp.MustCompile(`^\s+$`)
	valid_email = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

)

func TestNodeUserFields(t *testing.T) {
	user_node := model.User_node{
		User_id:             "unique_identifier",
		First_name:          "name1",
		Last_name:           "name2",
		Profile_picture_url: "link",
		Email:               "email_address",
		Created_at:          "when_they_joined",
	}

	if user_node.User_id != "unique_identifier" {
		t.Errorf("Expected unique_identifier but got %s", user_node.User_id)
	}
	if user_node.First_name != "name1" {
		t.Errorf("Expected name1 but got %s", user_node.First_name)
	}
	if user_node.Last_name != "name2" {
		t.Errorf("Expected name2 but got %s", user_node.Last_name)
	}
	if user_node.Profile_picture_url != "link" {
		t.Errorf("Expected link but got %s", user_node.Profile_picture_url)
	}
	if user_node.Email != "email_address" {
		t.Errorf("Expected email_address but got %s", user_node.Email)
	}
	if user_node.Created_at != "when_they_joined" {
		t.Errorf("Expected when_they_joined but got %s", user_node.Created_at)
	}

}

func TestNodeUserYamlUnmarshall(t *testing.T) {
	yamlData := []byte(`
User_id: unique_identifier
First_name: name1
Last_name: name2
Profile_picture_url: link
Email: email_address
Created_at: when_they_joined
`)

	var user_node model.User_node
	err := yaml.Unmarshal(yamlData, &user_node)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if user_node.User_id != "unique_identifier" {
		t.Errorf("Expected unique_identifier but got %s", user_node.User_id)
	}
}

func TestNodeUserEmptyFields(t *testing.T) {
	user_node := model.User_node{}

	if user_node.User_id != "" {
		t.Errorf("Expected empty user ID but got %s", user_node.User_id)
	}
}

func TestNodeMultipleUserNodes(t *testing.T) {
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

	if John.First_name != "John" {
		t.Errorf("Expected  but got %s", John.User_id)
	}
	if John.Email != "JohnDoe@email.com" {
		t.Errorf("Expected  but got %s", John.User_id)
	}
	if Jane.First_name != "Jane" {
		t.Errorf("Expected Jane but got %s", Jane.First_name)
	}
	if Jane.Email != "JaneSmith@email.com" {
		t.Errorf("Expected  but got %s", John.User_id)
	}
	if Joe.First_name != "Joe" {
		t.Errorf("Expected  but got %s", John.User_id)
	}
	if Joe.Email != "JoeJo@email.com" {
		t.Errorf("Expected JoeJo@email.com but got %s", Joe.Email)
	}
}

func TestNodeUserNodeIndependency(t *testing.T) {
	John := model.User_node{User_id: "19annahdksnHKAnskjs01192", First_name: "John", Last_name: "Doe",
		Profile_picture_url: "Placeholder1", Email: "JohnDoe@email.com", Created_at: "24/08/2026"}
	Jane := model.User_node{
		User_id: "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA", First_name: "Jane", Last_name: "Smith",
		Profile_picture_url: "Placeholder2", Email: "JaneSmith@email.com", Created_at: "03/12/2026"}

	John.Email = "JohnDoe123@email.com"

	if Jane.Email == "JohnDoe123@email.com" {
		t.Errorf("Expected Jane to be unaffected but got %s", Jane.Email)
	}
	if John.Email != "JohnDoe123@email.com" {
		t.Errorf("Expected JohnDoe123@email.com but got %s", John.Email)
	}
}

func TestNodeSlice(t *testing.T) {
	user_nodes := []model.User_node{
		{User_id: "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA", First_name: "Jane", Last_name: "Smith", Profile_picture_url: "Placeholder2",
			Email: "JaneSmith@email.com", Created_at: "03/12/2026"},
		{User_id: "JHnHJy8JBcrYI798NVdj0dn2KOSB9", First_name: "Joe", Last_name: "Jo", Profile_picture_url: "Placeholder3",
			Email: "JoeJo@email.com", Created_at: "02/01/2027"},
	}

	if len(user_nodes) != 2 {
		t.Errorf("Expected 2 user_nodes but got %v", len(user_nodes))
	}
	if user_nodes[0].First_name != "Jane" {
		t.Errorf("Expected Jane but got %s", user_nodes[0].First_name)
	}
	if user_nodes[1].First_name != "Joe" {
		t.Errorf("Expected Joe but got %s", user_nodes[1].First_name)
	}
}

func TestNodeNoDuplicateUserNodes(t *testing.T) {
	user_nodes := []model.User_node{
		{User_id: "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA", First_name: "Jane", Last_name: "Smith", Profile_picture_url: "Placeholder2",
			Email: "JaneSmith@email.com", Created_at: "03/12/2026"},
		{User_id: "JHnHJy8JBcrYI798NVdj0dn2KOSB9", First_name: "Joe", Last_name: "Jo", Profile_picture_url: "Placeholder3",
			Email: "JoeJo@email.com", Created_at: "02/01/2027"},
	}

	new_user_node := model.User_node{User_id: "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA", First_name: "Jane", Last_name: "Smith", Profile_picture_url: "Placeholder2",
		Email: "JaneSmith@email.com", Created_at: "03/12/2026"}

	duplicate_found := false
	for _, user_node := range user_nodes {
		if user_node.First_name == new_user_node.First_name && user_node.User_id == new_user_node.User_id {
			duplicate_found = true
			break
		}
	}
	if !duplicate_found {
		t.Errorf("Expected a duplicate user node to be found but it was not found.")
	}
}

func TestNodeInvalidEmail(t *testing.T) {
	user_nodes := []model.User_node{
		{User_id: "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA", First_name: "Jane", Last_name: "Smith", Profile_picture_url: "Placeholder2",
			Email: "JaneSmith@email.com", Created_at: "03/12/2026"},
		{User_id: "JHnHJy8JBcrYI798NVdj0dn2KOSB9", First_name: "Joe", Last_name: "Jo", Profile_picture_url: "Placeholder3",
			Email: "JoeJoemail.com", Created_at: "02/01/2027"},
	}

	invalid_email := false
	for _, user_node := range user_nodes {
		if !valid_email.MatchString(user_node.Email)  {
			invalid_email = true
			break
		}
	}
	if !invalid_email {
		t.Errorf("Expected an invalid email in the email field to be found but it was not found.")
	}
}

func TestNodeInvalidName(t *testing.T) {
	user_nodes := []model.User_node{
		{User_id: "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA", First_name: "Jane", Last_name: "Smith", Profile_picture_url: "Placeholder2",
			Email: "JaneSmith@email.com", Created_at: "03/12/2026"},
		{User_id: "JHnHJy8JBcrYI798NVdj0dn2KOSB9", First_name: "J0e", Last_name: "Jo", Profile_picture_url: "Placeholder3",
			Email: "JoeJo@email.com", Created_at: "02/01/2027"},
	}

	invalid_name := false
	for _, user_node := range user_nodes {
		if !valid_name.MatchString(user_node.First_name) || !valid_name.MatchString(user_node.Last_name)  {
			invalid_name = true
			break
		}
	}
	if !invalid_name {
		t.Errorf("Expected an invalid name in the First name field to be found but it was not found.")
	}
}

func TestNodeWhiteSpaceOnly(t *testing.T) {
	user_nodes := []model.User_node{
		{User_id: "29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA", First_name: "Jane", Last_name: "Smith", Profile_picture_url: "Placeholder2",
			Email: "JaneSmith@email.com", Created_at: "03/12/2026"},
		{User_id: "JHnHJy8JBcrYI798NVdj0dn2KOSB9", First_name: " ", Last_name: "Jo", Profile_picture_url: "Placeholder3",
			Email: "JoeJo@email.com", Created_at: "02/01/2027"},
	}

	invalid_field := false
	for _, user_node := range user_nodes {
		if valid_field.MatchString(user_node.First_name) || valid_field.MatchString(user_node.Last_name) || valid_field.MatchString(user_node.User_id) || valid_field.MatchString(user_node.Profile_picture_url) || valid_field.MatchString(user_node.Email) || valid_field.MatchString(user_node.Created_at) {
			invalid_field = true
			break
		}
	}
	if !invalid_field {
		t.Errorf("Expected an invalid field containing only whitespace to be found but it was not found.")
	}
}



