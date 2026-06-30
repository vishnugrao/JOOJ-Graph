package model

import (
)

type User_node struct{
	User_id string `yaml:"User_id"`
	First_name string `yaml:"First_name"`
	Last_name string `yaml:"Last_name"`
	Profile_picture_url string `yaml:"Profile_picture_url"`
	Email string `yaml:"Email"`
	Created_at string `yaml:"Created_at"`
}