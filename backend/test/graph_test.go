package test

import(
	// Internal Imports
	"JOOJ-Graph/backend/logic"
	"JOOJ-Graph/backend/model"

	// External Imports
	"testing"
)

func TestCreateGraph(t *testing.T) {
	Gm := logic.JoojGraphManager{InMemoryStore: make(map[string]model.Graph)}

	Gm.CreateGraph("SJ")

	if Gm.GetGraph("SJ").Name != "SJ" {
		t.Errorf("Wrong graph")
	}
}