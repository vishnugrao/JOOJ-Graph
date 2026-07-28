package test

import(
	// Internal Imports
	"JOOJ-Graph/backend/logic"
	"JOOJ-Graph/backend/model"

	// External Imports
	"testing"
	"regexp"
)

func TestCreateGraph(t *testing.T) {
	Gm := logic.JoojGraphManager{InMemoryStore: make(map[string]model.Graph)}

	Gm.CreateGraph("SJ")
	name := Gm.GetGraph("SJ").Name 
	if name != "SJ" {
		t.Errorf("Expected Graph SJk but got Graph %s", name)
	}
}

func TestDuplicateCreateGraph(t *testing.T) {
	Gm := logic.JoojGraphManager{InMemoryStore: make(map[string]model.Graph)}

	Gm.CreateGraph("SJ")
	Gm.CreateGraph("SJ")
	if len(Gm.InMemoryStore) != 1 {
		t.Errorf("Expected 1 graph but got %d", len(Gm.InMemoryStore))
	}
}

func TestCreateGraphInvalidGraphNames(t *testing.T) {
	Gm := logic.JoojGraphManager{InMemoryStore: make(map[string]model.Graph)}
	NonWhtSpc_Field := regexp.MustCompile(`^\s+$`)

	Gm.CreateGraph("")
	if _, exists := Gm.InMemoryStore[""]; exists{
		t.Errorf("Graph name should not be empty")
	}

	Gm.CreateGraph(" ")
	if _, exists := Gm.InMemoryStore[" "]; exists{
		t.Errorf("Graph name should not be whitespace")
	}

	Gm.CreateGraph("$ME^LB!!!")
	if _, exists := Gm.InMemoryStore["$ME^LB2!!!"]; exists{
		t.Errorf("Graph name should only contain Alphabet letters and Numbers")
	}
}