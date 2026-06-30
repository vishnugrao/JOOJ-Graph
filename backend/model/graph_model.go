package model

import (

)

type Graph struct {
	Nodes []User_node
	Edges []Edge
	NodeMap map[string]bool
	EdgeMap map[string]bool
}