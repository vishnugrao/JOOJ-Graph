package model


type User_node struct{
	User_id string `yaml:"User_id" json:"user_id"`
	First_name string `yaml:"First_name" json:"First_name"`
	Last_name string `yaml:"Last_name" json:"Last_name"`
	Profile_picture_url string `yaml:"Profile_picture_url" json:"Profile_picture_url"`
	Email string `yaml:"Email" json:"Email"`
	Created_at string `yaml:"Created_at" json:"Created_at"`
}