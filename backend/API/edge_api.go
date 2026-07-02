package api

// create edge function, enforcing rules from its test file
import (
	"errors"
	"regexp"
	"JOOJ-Graph/backend/model"
)

var (
	valid_tag_field = regexp.MustCompile(`^[a-zA-Z\/]+$`)
	valid_no_whitespace_field = regexp.MustCompile(`^\s+$`)
)

func CreateUserEdge(Source model.User_node, Target model.User_node, Edge_tag string, Edge_desc string, Created_at_edges string) (model.Edge, error) {
	if Source.User_id == "" || Target.User_id == "" || Edge_tag == "" || Edge_desc == "" || Created_at_edges == "" {
		return model.Edge{}, errors.New("No fields can be empty, please try again...")
	}

	if valid_no_whitespace_field.MatchString(Edge_tag) || valid_no_whitespace_field.MatchString(Edge_desc) || valid_no_whitespace_field.MatchString(Created_at_edges) {
		return model.Edge{}, errors.New("Invalid fields, please try again...")
	}

	if Source.User_id == Target.User_id {
		return model.Edge{}, errors.New("Cannot have an edge to yourself, please try again...")
	}

	if !valid_tag_field.MatchString(Edge_tag) {
		return model.Edge{}, errors.New("Edge tag can only contain upper/lower case letters and / only, please try again...")
	}
	
	edge := model.Edge{
		Source: Source,
		Target: Target,
		Edge_tag: Edge_tag,
		Edge_desc: Edge_desc,
		Created_at_edges: Created_at_edges }

	return edge, nil
}