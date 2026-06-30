package test

import (
	"testing"
	"github.com/abhay/JOOJ-Graph/backend/api"
	"github.com/abhay/JOOJ-Graph/backend/model"
)

func TestGraphAddNode(t *testing.T) {
	SJ := model.Graph{}

	node, err := api.CreateUserNode(
		"19annahdksnHKAnskjs01192",	"John",	"Doe",	"Placeholder1",	"JohnDoe@email.com","24/08/2026",	)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	SJ.Nodes = append(SJ.Nodes, node)

	if len(SJ.Nodes) != 1 {
		t.Errorf("Expected 1 node but got %d", len(SJ.Nodes))
	}

	if SJ.Nodes[0].First_name != "John" {
		t.Errorf("Expected John but got %s", SJ.Nodes[0].First_name)
	}
}

func TestGraphAddMultipleNode(t *testing.T) {
	SJ := model.Graph{}

	node1, err := api.CreateUserNode(
		"19annahdksnHKAnskjs01192",	"John",	"Doe",	"Placeholder1",	"JohnDoe@email.com","24/08/2026",	)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	node2, err := api.CreateUserNode(
		"29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA","Jane",	"Smith","Placeholder2",	"JaneSmith@email.com",	"03/12/2026",
	)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	node3, err := api.CreateUserNode(
		"JHnHJy8JBcrYI798NVdj0dn2KOSB9","Joe","Jo",	"Placeholder3",	"JoeJo@email.com",	"02/01/2027",	)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	SJ.Nodes = append(SJ.Nodes, node1)
	SJ.Nodes = append(SJ.Nodes, node2)
	SJ.Nodes = append(SJ.Nodes, node3)

	if len(SJ.Nodes) != 3 {
		t.Errorf("Expected 3 node but got %d", len(SJ.Nodes))
	}
	if SJ.Nodes[0].First_name != "John" {
		t.Errorf("Expected John but got %s", SJ.Nodes[0].First_name)
	}
	if SJ.Nodes[1].First_name != "Jane" {
		t.Errorf("Expected John but got %s", SJ.Nodes[1].First_name)
	}
	if SJ.Nodes[2].First_name != "Joe" {
		t.Errorf("Expected John but got %s", SJ.Nodes[2].First_name)
	}
}

func TestGraphDuplicateNodeDetection(t *testing.T) {
	SJ := model.Graph{}
	SJ.NodeMap = make(map[string]bool)

	node1, err := api.CreateUserNode(
		"19annahdksnHKAnskjs01192",	"John",	"Doe","Placeholder1","JohnDoe@email.com","24/08/2026",	)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	node2, err := api.CreateUserNode(
		"19annahdksnHKAnskjs01192",	"John",	"Doe",	"Placeholder1",	"JohnDoe@email.com","24/08/2026",	)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	node3, err := api.CreateUserNode(
		"JHnHJy8JBcrYI798NVdj0dn2KOSB9","Joe","Jo",	"Placeholder3",	"JoeJo@email.com",	"02/01/2027",	)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	if !SJ.NodeMap[node1.User_id] {
		SJ.Nodes = append(SJ.Nodes, node1)
		SJ.NodeMap[node1.User_id] = true
	}

	if !SJ.NodeMap[node2.User_id] {
		SJ.Nodes = append(SJ.Nodes, node2)
		SJ.NodeMap[node2.User_id] = true
	}

	if !SJ.NodeMap[node3.User_id] {
		SJ.Nodes = append(SJ.Nodes, node3)
		SJ.NodeMap[node3.User_id] = true
	}

	if len(SJ.Nodes) != 2 {
		t.Errorf("Expected 2 node but got %d", len(SJ.Nodes))
	}
	if SJ.Nodes[0].First_name != "John" {
		t.Errorf("Expected John but got %s", SJ.Nodes[0].First_name)
	}
	if SJ.Nodes[1].First_name != "Joe" {
		t.Errorf("Expected John but got %s", SJ.Nodes[1].First_name)
	}
}

func TestGraphRemoveNode(t *testing.T) {
	SJ := model.Graph{}
	SJ.NodeMap = make(map[string]bool)
	node, err := api.CreateUserNode(
		"19annahdksnHKAnskjs01192",	"John",	"Doe","Placeholder1","JohnDoe@email.com","24/08/2026",	)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	removeID := node.User_id

	if !SJ.NodeMap[node.User_id] {
		SJ.Nodes = append(SJ.Nodes, node)
		SJ.NodeMap[node.User_id] = true
	}

	if !SJ.NodeMap[removeID] {
		t.Errorf("Node not found")
	} else {
		for i, node := range SJ.Nodes {
			if node.User_id == removeID {
				SJ.Nodes = append(SJ.Nodes[:i], SJ.Nodes[i+1:]...)
				delete(SJ.NodeMap, removeID)
				break
			}
		}
	}

	if len(SJ.Nodes) != 0 {
		t.Errorf("Expected 0 node but got %d", len(SJ.Nodes))
	}
}

func TestGraphAddEdge(t *testing.T) {

	SJ := model.Graph{}
	SJ.EdgeMap = make(map[string]bool)

	John, err := api.CreateUserNode(
		"19annahdksnHKAnskjs01192",	"John",	"Doe",	"Placeholder1",	"JohnDoe@email.com","24/08/2026",)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	Jane, err := api.CreateUserNode(
		"29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA","Jane","Smith",	"Placeholder2",	"JaneSmith@email.com",	"03/12/2026",)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge, err := api.CreateUserEdge(John, Jane, "Family", "Desc AB", "20/08/26")

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge_key := edge.Source.User_id + "-" + edge.Target.User_id

	if !SJ.EdgeMap[edge_key] {
		SJ.Edges = append(SJ.Edges, edge)
		SJ.EdgeMap[edge_key] = true
	}

	if len(SJ.Edges) != 1 {
		t.Errorf("Expected 1 but got %v", len(SJ.Edges))
	}
}

func TestGraphAddEdgeWithInvalidNode(t *testing.T) {

	SJ := model.Graph{}
	SJ.NodeMap = make(map[string]bool)
	SJ.EdgeMap = make(map[string]bool)

	John, err := api.CreateUserNode(
		"19annahdksnHKAnskjs01192",	"John",	"Doe",	"Placeholder1",	"JohnDoe@email.com","24/08/2026",)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	Jane, err := api.CreateUserNode(
		"29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA","Jane","Smith",	"Placeholder2",	"JaneSmith@email.com",	"03/12/2026",)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	Joe, err := api.CreateUserNode(
		"JHnHJy8JBcrYI798NVdj0dn2KOSB9","Joe",	"Jo",	"Placeholder3",	"JoeJo@email.com",	"02/01/2027",	)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge, err := api.CreateUserEdge(John, Jane, "Family", "Desc AB", "20/08/26")

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge2, err := api.CreateUserEdge(Jane, Joe, "Friend", "Desc BC", "21/09/27")

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	if !SJ.NodeMap[John.User_id] {
		SJ.Nodes = append(SJ.Nodes, John)
		SJ.NodeMap[John.User_id] = true
	}
	if !SJ.NodeMap[Jane.User_id] {
		SJ.Nodes = append(SJ.Nodes, Jane)
		SJ.NodeMap[Jane.User_id] = true
	}

	edge_key := edge.Source.User_id + "-" + edge.Target.User_id
	edge2_key := edge2.Source.User_id + "-" + edge2.Target.User_id

	if !SJ.EdgeMap[edge_key] {
		SJ.Edges = append(SJ.Edges, edge)
		SJ.EdgeMap[edge_key] = true
	}

	if SJ.NodeMap[edge2.Source.User_id] && SJ.NodeMap[edge2.Target.User_id] {
		if !SJ.EdgeMap[edge2_key] {
			SJ.Edges = append(SJ.Edges, edge2)
			SJ.EdgeMap[edge2_key] = true
		}
	}

	if len(SJ.Edges) != 1 {
		t.Errorf("Expected 1 but got %v", len(SJ.Edges))
	}
}

func TestGraphAddMultipleEdge(t *testing.T) {

	SJ := model.Graph{}
	SJ.EdgeMap = make(map[string]bool)

	John, err := api.CreateUserNode(
		"19annahdksnHKAnskjs01192",	"John",	"Doe",	"Placeholder1",	"JohnDoe@email.com","24/08/2026",)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	Jane, err := api.CreateUserNode(
		"29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA","Jane","Smith",	"Placeholder2",	"JaneSmith@email.com",	"03/12/2026",)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	Joe, err := api.CreateUserNode(
		"JHnHJy8JBcrYI798NVdj0dn2KOSB9","Joe",	"Jo",	"Placeholder3",	"JoeJo@email.com",	"02/01/2027",	)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge1, err := api.CreateUserEdge(John, Jane, "Family", "Desc AB", "20/08/26")
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge2, err := api.CreateUserEdge(Jane, John, "Family", "Desc BA", "20/08/26")
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge3, err := api.CreateUserEdge(John, Joe, "Friend", "Desc AC", "11/02/27")
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge4, err := api.CreateUserEdge(Jane, Joe, "Friend", "Desc BC", "21/09/27")

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge1_key := edge1.Source.User_id + "-" + edge1.Target.User_id

	if !SJ.EdgeMap[edge1_key] {
		SJ.Edges = append(SJ.Edges, edge1)
		SJ.EdgeMap[edge1_key] = true
	}

	edge2_key := edge2.Source.User_id + "-" + edge2.Target.User_id

	if !SJ.EdgeMap[edge2_key] {
		SJ.Edges = append(SJ.Edges, edge2)
		SJ.EdgeMap[edge2_key] = true
	}
	edge3_key := edge3.Source.User_id + "-" + edge3.Target.User_id

	if !SJ.EdgeMap[edge3_key] {
		SJ.Edges = append(SJ.Edges, edge3)
		SJ.EdgeMap[edge3_key] = true
	}
	edge4_key := edge4.Source.User_id + "-" + edge4.Target.User_id

	if !SJ.EdgeMap[edge4_key] {
		SJ.Edges = append(SJ.Edges, edge4)
		SJ.EdgeMap[edge4_key] = true
	}

	if len(SJ.Edges) != 4 {
		t.Errorf("Expected 4 but got %v", len(SJ.Edges))
	}

	if SJ.Edges[0].Source.First_name != "John" {
		t.Errorf("Expected John but got %s", SJ.Edges[0].Source.First_name)
	}
	if SJ.Edges[1].Source.First_name != "Jane" {
		t.Errorf("Expected Jane but got %s", SJ.Edges[1].Source.First_name)
	}
	if SJ.Edges[2].Source.First_name != "John" {
		t.Errorf("Expected John but got %s", SJ.Edges[2].Source.First_name)
	}
	if SJ.Edges[3].Source.First_name != "Jane" {
		t.Errorf("Expected Jane but got %s", SJ.Edges[3].Source.First_name)
	}

}

func TestGraphDuplicateEdgeDetection(t *testing.T) {
	SJ := model.Graph{}
	SJ.EdgeMap = make(map[string]bool)

	John, err := api.CreateUserNode(
		"19annahdksnHKAnskjs01192",	"John",	"Doe",	"Placeholder1",	"JohnDoe@email.com","24/08/2026",	)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	Joe, err := api.CreateUserNode(
		"JHnHJy8JBcrYI798NVdj0dn2KOSB9","Joe",	"Jo","Placeholder3","JoeJo@email.com",	"02/01/2027",)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge1, err := api.CreateUserEdge(John, Joe, "Friend", "Desc AB", "20/08/26")
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge2, err := api.CreateUserEdge(Joe, John, "Family", "Desc BA", "20/08/26")
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge3, err := api.CreateUserEdge(John, Joe, "Friend", "Desc AB", "20/08/26")
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge1_key := edge1.Source.User_id + "-" + edge1.Target.User_id
	edge2_key := edge2.Source.User_id + "-" + edge2.Target.User_id
	edge3_key := edge3.Source.User_id + "-" + edge3.Target.User_id

	if !SJ.EdgeMap[edge1_key] {
		SJ.Edges = append(SJ.Edges, edge1)
		SJ.EdgeMap[edge1_key] = true
	}
	if !SJ.EdgeMap[edge2_key] {
		SJ.Edges = append(SJ.Edges, edge2)
		SJ.EdgeMap[edge2_key] = true
	}
	if !SJ.EdgeMap[edge3_key] {
		SJ.Edges = append(SJ.Edges, edge3)
		SJ.EdgeMap[edge3_key] = true
	}

	if len(SJ.Edges) != 2 {
		t.Errorf("Expected 2 edges but got %d", len(SJ.Edges))
	}
	if SJ.Edges[0].Source.First_name != "John" {
		t.Errorf("Expected John but got %s", SJ.Edges[0].Source.First_name)
	}
	if SJ.Edges[1].Source.First_name != "Joe" {
		t.Errorf("Expected Joe but got %s", SJ.Edges[1].Source.First_name)
	}
}
func TestGraphRemoveEdge(t *testing.T) {
	SJ := model.Graph{}
	SJ.NodeMap = make(map[string]bool)
	SJ.EdgeMap = make(map[string]bool)

	John, err := api.CreateUserNode(
		"19annahdksnHKAnskjs01192",	"John",	"Doe","Placeholder1","JohnDoe@email.com","24/08/2026",	)
	
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	Jane, err := api.CreateUserNode(
		"29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA","Jane","Smith",	"Placeholder2",	"JaneSmith@email.com",	"03/12/2026",)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	Joe, err := api.CreateUserNode(
		"JHnHJy8JBcrYI798NVdj0dn2KOSB9","Joe",	"Jo",	"Placeholder3",	"JoeJo@email.com",	"02/01/2027",	)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	if !SJ.NodeMap[John.User_id] {
		SJ.Nodes = append(SJ.Nodes, John)
		SJ.NodeMap[John.User_id] = true
	}
	if !SJ.NodeMap[Jane.User_id] {
		SJ.Nodes = append(SJ.Nodes, Jane)
		SJ.NodeMap[Jane.User_id] = true
	}
	if !SJ.NodeMap[Joe.User_id] {
		SJ.Nodes = append(SJ.Nodes, Joe)
		SJ.NodeMap[Joe.User_id] = true
	}

	edge1, err := api.CreateUserEdge(John, Jane, "Friend", "Desc AB", "20/08/26")
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge2, err := api.CreateUserEdge(Jane, John, "Family", "Desc BA", "20/08/26")
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge3, err := api.CreateUserEdge(Jane, Joe, "Friend", "Desc AC", "11/09/26")
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge1_key := edge1.Source.User_id + "-" + edge1.Target.User_id
	edge2_key := edge2.Source.User_id + "-" + edge2.Target.User_id
	edge3_key := edge3.Source.User_id + "-" + edge3.Target.User_id

	if !SJ.EdgeMap[edge1_key] {
		SJ.Edges = append(SJ.Edges, edge1)
		SJ.EdgeMap[edge1_key] = true
	}
	if !SJ.EdgeMap[edge2_key] {
		SJ.Edges = append(SJ.Edges, edge2)
		SJ.EdgeMap[edge2_key] = true
	}
	if !SJ.EdgeMap[edge3_key] {
		SJ.Edges = append(SJ.Edges, edge3)
		SJ.EdgeMap[edge3_key] = true
	}

	removeID1 := Jane.User_id
	removeID2 := Joe.User_id

	removekey := removeID1 + "-" + removeID2

	if !SJ.EdgeMap[removekey] {
		t.Errorf("Node not found")
	} else {
		for i:= len(SJ.Edges) -1; i>=0; i-- {
			if SJ.Edges[i].Source.User_id == removeID1 && SJ.Edges[i].Target.User_id == removeID2 {
				delete(SJ.EdgeMap, SJ.Edges[i].Source.User_id + "-" + SJ.Edges[i].Target.User_id)
				SJ.Edges = append(SJ.Edges[:i], SJ.Edges[i+1:]...)
			}
		}
	}

	if len(SJ.Edges) != 2 {
		t.Errorf("Expected 2 edges but got %d", len(SJ.Edges))
	}
	if len(SJ.Nodes) != 3 {
		t.Errorf("Expected 3 but got %d", len(SJ.Nodes))
	}
	if SJ.Edges[0].Source.First_name != "John"  {
		t.Errorf("Expected John but got %s", SJ.Edges[0].Source.First_name)
	}
	if SJ.Edges[1].Source.First_name != "Jane"  {
		t.Errorf("Expected Jane but got %s", SJ.Edges[1].Source.First_name)
	}
}
func TestGraphRemoveNodeAndEdge(t *testing.T) {
	SJ := model.Graph{}
	SJ.NodeMap = make(map[string]bool)
	SJ.EdgeMap = make(map[string]bool)

	John, err := api.CreateUserNode(
		"19annahdksnHKAnskjs01192",	"John",	"Doe","Placeholder1","JohnDoe@email.com","24/08/2026",	)
	
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	Jane, err := api.CreateUserNode(
		"29fhhfnskjJKaspMWaBsJHSKKNSGG02nJassmee93n3NA","Jane","Smith",	"Placeholder2",	"JaneSmith@email.com",	"03/12/2026",)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	Joe, err := api.CreateUserNode(
		"JHnHJy8JBcrYI798NVdj0dn2KOSB9","Joe",	"Jo",	"Placeholder3",	"JoeJo@email.com",	"02/01/2027",	)

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	if !SJ.NodeMap[John.User_id] {
		SJ.Nodes = append(SJ.Nodes, John)
		SJ.NodeMap[John.User_id] = true
	}
	if !SJ.NodeMap[Jane.User_id] {
		SJ.Nodes = append(SJ.Nodes, Jane)
		SJ.NodeMap[Jane.User_id] = true
	}
		if !SJ.NodeMap[Joe.User_id] {
		SJ.Nodes = append(SJ.Nodes, Joe)
		SJ.NodeMap[Joe.User_id] = true
	}

	edge1, err := api.CreateUserEdge(John, Jane, "Friend", "Desc AB", "20/08/26")
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge2, err := api.CreateUserEdge(Jane, John, "Family", "Desc BA", "20/08/26")
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge3, err := api.CreateUserEdge(Jane, Joe, "Friend", "Desc AC", "11/09/26")
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	edge1_key := edge1.Source.User_id + "-" + edge1.Target.User_id
	edge2_key := edge2.Source.User_id + "-" + edge2.Target.User_id
	edge3_key := edge3.Source.User_id + "-" + edge3.Target.User_id

	if !SJ.EdgeMap[edge1_key] {
		SJ.Edges = append(SJ.Edges, edge1)
		SJ.EdgeMap[edge1_key] = true
	}
	if !SJ.EdgeMap[edge2_key] {
		SJ.Edges = append(SJ.Edges, edge2)
		SJ.EdgeMap[edge2_key] = true
	}
	if !SJ.EdgeMap[edge3_key] {
		SJ.Edges = append(SJ.Edges, edge3)
		SJ.EdgeMap[edge3_key] = true
	}

	removeID := Joe.User_id

	if !SJ.NodeMap[removeID] {
		t.Errorf("Node not found")
	} else {
		for i, node := range SJ.Nodes {
			if node.User_id == removeID {
				SJ.Nodes = append(SJ.Nodes[:i], SJ.Nodes[i+1:]...)
				delete(SJ.NodeMap, removeID)
				break
			}
		}

		for i:= len(SJ.Edges) -1; i>=0; i-- {
			if SJ.Edges[i].Source.User_id == removeID || SJ.Edges[i].Target.User_id == removeID {
				delete(SJ.EdgeMap, SJ.Edges[i].Source.User_id + "-" + SJ.Edges[i].Target.User_id)
				delete(SJ.EdgeMap, SJ.Edges[i].Target.User_id + "-" + SJ.Edges[i].Source.User_id)
				SJ.Edges = append(SJ.Edges[:i], SJ.Edges[i+1:]...)
			}
		}
	}

	if len(SJ.Nodes) != 2 {
		t.Errorf("Expected 2 node but got %d", len(SJ.Nodes))
	}
	if len(SJ.Edges) != 2 {
		t.Errorf("Expected 2 edges but got %d", len(SJ.Edges))
	}
	if SJ.Edges[0].Source.First_name != "John"  {
		t.Errorf("Expected John but got %s", SJ.Edges[0].Source.First_name)
	}
	if SJ.Edges[1].Source.First_name != "Jane"  {
		t.Errorf("Expected Jane but got %s", SJ.Edges[1].Source.First_name)
	}
}



