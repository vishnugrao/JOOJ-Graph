package api

import (
	"errors"
	"regexp"
	"github.com/abhay/JOOJ-Graph/backend/model"
)

var (
	valid_email = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	valid_name = regexp.MustCompile(`^[a-zA-Z\s\-']+$`)
	valid_field = regexp.MustCompile(`^\s+$`)
)

func CreateUserNode(User_id string, First_name string, Last_name string, Profile_picture_url string, Email string, Created_at string) (model.User_node, error) {
	if User_id == "" || First_name == "" || Last_name == "" || Email == "" || Created_at == "" {
		return model.User_node{}, errors.New("User ID, First name, Last name, Email or Created At date cannot be empty, please try again...")
	}
	
	if !valid_email.MatchString(Email)  {
		return model.User_node{}, errors.New("Invalid email, please try again...")
	}

	if !valid_name.MatchString(First_name) || !valid_name.MatchString(Last_name)  {
		return model.User_node{}, errors.New("Invalid First or Last name, please try again...")
	}

	if valid_field.MatchString(First_name) || valid_field.MatchString(Last_name) || valid_field.MatchString(User_id) || valid_field.MatchString(Profile_picture_url) || valid_field.MatchString(Email) || valid_field.MatchString(Created_at) {
		return model.User_node{}, errors.New("Invalid fields, please try again...")
	}
	user := model.User_node{
		User_id: User_id,
		First_name: First_name,
		Last_name: Last_name,
		Profile_picture_url: Profile_picture_url,
		Email: Email,
		Created_at: Created_at }

	return user, nil
}