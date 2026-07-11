package model

type Graph struct {
	Name string `json:"name"`
	Nodes []User_node `json:"nodes"`
	Edges []Edge `json:"edges"`
	NodeMap map[string]bool `json:"node_map"`
	EdgeMap map[string]bool `json:"edge_map"`
}