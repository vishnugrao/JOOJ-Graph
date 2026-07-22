package main

import (
// Internal imports
	api "JOOJ-Graph/backend/API"

// External imports
	"fmt"
	"net/http"
)

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/reset", api.StoreResetHandler)	

	router.HandleFunc("/create/graph", api.CreateGraphHandler)
	router.HandleFunc("/create/node", api.CreateNodeHandler)
	router.HandleFunc("/create/edge", api.CreateEdgeHandler)
	
	router.HandleFunc("/graph/add/node", api.GraphAddNodeHandler)
	router.HandleFunc("/graph/add/edge", api.GraphAddEdgeHandler)

	router.HandleFunc("/graph/delete/node", api.GraphRemoveNodeHandler)
	router.HandleFunc("/graph/delete/edge", api.GraphRemoveEdgeHandler)

	router.HandleFunc("/get/allgraphs", api.GetAllGraphHandler)

	router.HandleFunc("/get/graph", api.GetGraphHandler)

	router.HandleFunc("/graph/get/node", api.GetNodeHandler)

	// router.HandleFunc("/read/graph", api.ReadGraphHandler) Visualization, Opens browswer to show graph
	// router.HandleFunc("/read/node", api.ReadNodeHandler) A page with information on the node, GET request
	// router.HandleFunc("/read/edge", api.ReadEdgeHandler) A page with information on the node, GET request

	// router.HandleFunc("/update/graph", api.UpdateGraphHandler)
	// router.HandleFunc("/update/node", api.UpdateNodeHandler)
	// router.HandleFunc("/update/edge", api.UpdateEdgeHandler)

	// router.HandleFunc("/delete/graph", api.DeleteGraphHandler)
	// router.HandleFunc("/delete/node", api.DeleteNodeHandler)
	// router.HandleFunc("/delete/edge", api.DeleteEdgeHandler)

	fmt.Println("Server running on port 6767...")
	if err := http.ListenAndServe(":6767", router); err != nil {
		panic(err)
	}

}
