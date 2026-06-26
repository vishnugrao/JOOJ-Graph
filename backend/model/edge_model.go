package model

import (
)

type Edge struct {
	Source User_node `yaml:"Source"`
	Target User_node `yaml:"Target"`
	Edge_tag string `yaml:"Edge_tag"`
	Edge_desc string `yaml:"Edge_desc"`
	Created_at_edges string `yaml:"Created_at_edges"`
}