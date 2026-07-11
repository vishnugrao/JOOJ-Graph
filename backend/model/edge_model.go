package model


type Edge struct {
	Source User_node `yaml:"Source" json:"source"`
	Target User_node `yaml:"Target" json:"target"`
	Edge_tag string `yaml:"Edge_tag" json:"edge_tag"` 
	Edge_desc string `yaml:"Edge_desc" json:"edge_desc"`
	Created_at_edges string `yaml:"Created_at_edges" json:"created_at_edges"`
}