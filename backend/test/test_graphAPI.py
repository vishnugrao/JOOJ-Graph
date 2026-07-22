import pytest
import requests
import json
BACKEND_URL = "http://localhost:6767"

def test_ServerReachability(add_graph_response_raw):
    assert add_graph_response_raw.status_code != 404, f"Server not reachable, got {add_graph_response_raw.status_code}"

def test_CorrectPort(add_graph_response_raw):
    try:
        assert add_graph_response_raw.status_code != None
    except requests.exceptions.ConnectionError:
        pytest.fail("Could not connect to the port 6767")


def test_TimeOutHandling(add_graph_response_raw):
    try:
        assert add_graph_response_raw.status_code != None
    except requests.exceptions.Timeout:
        pytest.fail("Server timed out after 5 seconds")

def test_nonEmptyResponse(add_graph_response_raw):
    assert add_graph_response_raw.text != ""

def test_ServerDown():
    try:
        response = requests.get("http://localhost:1403/create/graph", timeout=3)
        pytest.fail("Expected connection error but got a response")
    except requests.exceptions.ConnectionError:
        pass


def test_JsonPayloadPost(add_graph_response_raw):
    assert add_graph_response_raw.status_code == 201, f"Expected 201 but got {add_graph_response_raw.status_code}"


def test_correctHeaders(add_graph_response_raw):
    assert add_graph_response_raw.status_code == 201, f"Expected 201 but got {add_graph_response_raw.status_code}"
    assert "Content-Type" in add_graph_response_raw.request.headers
    assert add_graph_response_raw.request.headers["Content-Type"] == "application/json"


def test_emptyPayload():
    response = requests.post(f"http://localhost:6767/create/graph", data="",headers={"Content-Type":"application/json"} ,timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"


def test_invalidJson():
    response = requests.post(f"http://localhost:6767/create/graph", data="Invalid Json", headers={"Content-Type":"application/json"}, timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"
    

def test_statusCodes(add_graph_response_raw, graph_payload):
    assert add_graph_response_raw.status_code == 201, f"Expected 201 but got {add_graph_response_raw.status_code}"

    response2 = requests.post(f"http://localhost:6767/create/graph/dummy", json=graph_payload, timeout=3)
    assert response2.status_code == 404, f"Expected 404 but got {response2.status_code}"

    response3 = requests.get(f"http://localhost:6767/create/graph")
    assert response3.status_code == 400, f"Expected 400 but got {response3.status_code}"


def test_isString(add_graph_response_raw):
    data = add_graph_response_raw.text
    assert add_graph_response_raw.status_code == 201, f"Expected 201 but got {add_graph_response_raw.status_code}"
    assert isinstance(data, str), f"Expected string but got {type(data)}"

def test_graphname(add_graph_response):
    assert add_graph_response["name"] == "SJ", f"Expected SJ but got {add_graph_response['name']}"

def test_graph_empty_lists_and_maps(add_graph_response):
    assert len(add_graph_response['nodes']) == 0, f"Expected nodes to be empty but got {len(add_graph_response['nodes'])} items"
    assert len(add_graph_response['edges']) == 0, f"Expected edges to be empty but got {len(add_graph_response['edges'])} items"
    assert len(add_graph_response['node_map']) == 0, f"Expected node_map to be empty but got {len(add_graph_response['node_map'])} items"
    assert len(add_graph_response['edge_map']) == 0, f"Expected edge_map to be empty but got {len(add_graph_response['edge_map'])} items"

def test_duplicate_graph_names(graph_payload):
    requests.post(f"http://localhost:6767/create/graph", json=graph_payload, timeout=3)
    response = requests.post(f"http://localhost:6767/create/graph", json=graph_payload, timeout=3)
    assert response.status_code == 409, f"Expected 409 but got {response.status_code}"
    

def test_nonEmptyGraphName():
    payload = {"name": ""}
    response = requests.post(f"http://localhost:6767/create/graph", json=payload, timeout=3)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"

def test_graph_name_str(add_graph_response):
    assert isinstance(add_graph_response["name"], str)

def test_graph_expected_fields(add_graph_response):
    assert add_graph_response["name"] != "", f"Expected name SJ but got {add_graph_response["name"]}"
    assert add_graph_response["nodes"] == [], f"Expected nodes to be a list but got {add_graph_response["nodes"]}"
    assert add_graph_response["edges"] == [], f"Expected edges to be a list but got {add_graph_response["edges"]}"
    assert add_graph_response["node_map"] == {}, f"Expected node_map to be a map but got {add_graph_response["node_map"]}"
    assert add_graph_response["edge_map"] == {}, f"Expected edge_map to be a map but got {add_graph_response["edge_map"]}"

def test_correctGraphRetrievalMethod(get_graph_response_raw):
    assert get_graph_response_raw.status_code == 200, f"Expected 200 but got {get_graph_response_raw.status_code}"

def test_invalidGraphRetrieval():
    response = requests.get(f"{BACKEND_URL}/get/graph?name=Nothing", timeout=3)
    assert response.status_code == 404, f"Expected 404 but got {response.status_code}"

def test_returnsCorrectGraphFields(get_graph_response):
    data = get_graph_response
    assert data['name'] == 'SJ'
    assert data['nodes'] == []
    assert data['edges'] == []
    assert data['node_map'] == {}
    assert data['edge_map'] == {}

def test_emptyNameParamGraph():
    response = requests.get(f"{BACKEND_URL}/get/graph", timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"

def test_correctlyRetrivesAllGraphs(get_all_graphs_response):
    assert len(get_all_graphs_response) == 2
    assert 'ADL' in get_all_graphs_response
    assert 'SJ' in get_all_graphs_response

def test_getGraphCorrectHeaders(get_graph_response_raw):
    response = get_graph_response_raw
    assert response.headers['Content-Type'].startswith('application/json')

def test_correctlyGetAllGraphs(get_all_graphs_response_raw):
    assert get_all_graphs_response_raw.status_code == 200, f"Expected 200 but got {get_all_graphs_response_raw}"
    data = get_all_graphs_response_raw.json()
    assert 'SJ' in data

def test_getAllGraphsDict(get_all_graphs_response_raw):
    data = get_all_graphs_response_raw.json()
    assert isinstance(data, dict)

def test_getAllGraphsEmpty():
    response = requests.get(f"{BACKEND_URL}/get/allgraphs", timeout=5)
    data = response.json()
    assert isinstance(data, dict)
    assert len(data) == 0

def test_correctlyAddedEdgeInGraph(graph_with_one_edge_response):
    assert len(graph_with_one_edge_response['edges']) == 1, f"Expected 1 but got {len(graph_with_one_edge_response['edges'])}"

def test_correctlyAddedEdgeMapInGraph(graph_with_one_edge_response):
    edge = graph_with_one_edge_response['edges'][0]
    edge_key = edge['source']['user_id'] + "-" + edge['target']['user_id']
    assert edge_key in graph_with_one_edge_response['edge_map']
    assert graph_with_one_edge_response['edge_map'][edge_key] == True

def test_correctGraphEdgeFields(graph_with_one_edge_response):
    edge = graph_with_one_edge_response['edges'][0]
    assert edge['edge_tag'] == "Friend"
    assert edge['edge_desc'] == "From Highschool"

def test_addEdgeInvalidGraph(add_edge_response):
    payload = {"graphname" : "Invalid", "edge" : add_edge_response}
    response = requests.post(f"{BACKEND_URL}/graph/add/edge", json=payload, timeout=3)
    assert response.status_code == 404, f"Expected 404 but got {response.status_code}"

def test_addEdgeEmptyGraph(add_edge_response):
    payload = {"graphname" : "", "edge" : add_edge_response}
    response = requests.post(f"{BACKEND_URL}/graph/add/edge", json=payload, timeout=3)
    assert response.status_code == 404, f"Expected 404 but got {response.status_code}"

def test_invalidEdgeJsonPayload():
    response = requests.post(f"{BACKEND_URL}/graph/add/edge", data="Invalid Json", headers={"Content-Type" : "application/json"}, timeout=3)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"

def test_duplicateGraphEdgeCase(graph_with_one_edge_response):
    graph = requests.get(f"{BACKEND_URL}/get/graph?name={graph_with_one_edge_response['name']}",  timeout=5).json()

    edge_payload = {"source" : graph['nodes'][0], "target" : graph['nodes'][1], "edge_tag": "Friend", "edge_desc" : "From Highschool"}
    edge = requests.post(f"{BACKEND_URL}/create/edge", json=edge_payload, timeout=5).json()

    payload = {"graphname" : graph['name'] , "edge" : edge}
    response = requests.post(f"{BACKEND_URL}/graph/add/edge", json=payload,  timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"

def test_selfGraphEdgeCase(graph_with_one_edge_response):
    graph = requests.get(f"{BACKEND_URL}/get/graph?name={graph_with_one_edge_response['name']}",  timeout=5).json()

    edge_payload = {"source" : graph['nodes'][0], "target" : graph['nodes'][0], "edge_tag": "Myself", "edge_desc" : "Birth Certificate"}

    payload = {"graphname" : graph['name'] , "edge" : edge_payload}
    response = requests.post(f"{BACKEND_URL}/graph/add/edge", json=payload,  timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"

def test_invalidEdgeSource(graph_with_one_edge_response):
    graph = graph_with_one_edge_response
    fakeNode = {"user_id": "X", "First_name" : "X" , "Last_name": "X", "Email" :"X", "Profile_picture_url": ""}

    edge = {"source" : fakeNode, "target" : graph_with_one_edge_response['nodes'][1], "edge_tag": "Friend", "edge_desc" : "From Highschool"}
    payload = {"graphname" : graph['name'] , "edge" : edge}
    response = requests.post(f"{BACKEND_URL}/graph/add/edge", json=payload,  timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"

def test_invalidEdgeTarget(graph_with_one_edge_response):
    graph = graph_with_one_edge_response
    fakeNode = {"user_id": "X", "First_name" : "X" , "Last_name": "X", "Email" :"X", "Profile_picture_url": ""}

    edge = {"source" : graph_with_one_edge_response['nodes'][0], "target" : fakeNode, "edge_tag": "Friend", "edge_desc" : "From Highschool"}
    payload = {"graphname" : graph['name'] , "edge" : edge}
    response = requests.post(f"{BACKEND_URL}/graph/add/edge", json=payload,  timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"

def test_correctEdgeHeaderResponse(graph_with_one_edge_response_raw):
    assert graph_with_one_edge_response_raw.status_code == 200, f"Expected 200 but got {graph_with_one_edge_response_raw}"
    assert graph_with_one_edge_response_raw.headers['Content-Type'].startswith("application/json")

def test_correctlyRemovesGraphNode(graph_with_one_node_response):
    payload = {"graphname": graph_with_one_node_response["name"], "node": graph_with_one_node_response["nodes"][0]}
    response = requests.post(f"{BACKEND_URL}/graph/delete/node", json=payload,  timeout=5)
    assert response.status_code == 200, f"Expected 200 but got {response.status_code}"
    graph = response.json()
    assert len(graph['node_map']) == 0, f"Expected 0 keys in the node map but got {len(graph['node_map'])}"
    assert len(graph['nodes']) == 0, f"Expected 0 nodes in the nodes list but got {len(graph['nodes'])}"
    assert response.headers["Content-Type"].startswith("application/json")

def test_removeNodeFromInvalidGraph(graph_with_one_node_response):
    node = graph_with_one_node_response['nodes'][0]
    payload = {"graphname": "Invalid", "node": node}
    response = requests.post(f"{BACKEND_URL}/graph/delete/node", json=payload,  timeout=5)    
    assert response.status_code == 404, f"Expected 404 but got {response.status_code}"

def test_removeFakeNodefromGraph(graph_with_one_node_response):
    fakeNode = {"user_id": "X", "First_name" : "X" , "Last_name": "X", "Email" :"X", "Profile_picture_url": ""}
    payload = {"graphname": graph_with_one_node_response['name'], "node": fakeNode}
    response = requests.post(f"{BACKEND_URL}/graph/delete/node", json=payload,  timeout=5)    
    assert response.status_code == 404, f"Expected 404 but got {response.status_code}"

def test_removeNodeFromEmptyGraphName(graph_with_one_node_response):
    node = graph_with_one_node_response['nodes'][0]
    payload = {"graphname": "", "node": node}
    response = requests.post(f"{BACKEND_URL}/graph/delete/node", json=payload,  timeout=5)    
    assert response.status_code == 404, f"Expected 404 but got {response.status_code}"

def test_removeNodeInvalidJson():
    response = requests.post(f"http://localhost:6767/graph/delete/node", data="Invalid Json", headers={"Content-Type":"application/json"}, timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"

def test_correctlyRemovedNodeDeletesEdges(graph_with_one_edge_response):
    node = graph_with_one_edge_response["nodes"][0]
    payload = {"graphname": graph_with_one_edge_response["name"], "node": node}
    response = requests.post(f"{BACKEND_URL}/graph/delete/node", json=payload,  timeout=5)
    assert response.status_code == 200, f"Expected 200 but got {response.status_code}"
    graph = response.json()
    assert len(graph['edges']) == 0, f"Expected 0 but got {len(graph['edges'])}"
    assert len(graph['edge_map']) == 0, f"Expected 0 but got {len(graph['edge_map'])}"
    assert len(graph['nodes']) == 1, f"Expected 1 but got {len(graph['nodes'])}"
    assert graph['nodes'][0]['Email'] != node['Email']

def test_removesSameNodeTwice(graph_with_one_node_response):
    graph = graph_with_one_node_response
    node = graph['nodes'][0]
    payload = {"graphname": graph["name"], "node": node}
    successful_response = requests.post(f"{BACKEND_URL}/graph/delete/node", json=payload,  timeout=5)
    assert successful_response.status_code == 200, f"Expected 200 but got {successful_response.status_code}"

    fail_response = requests.post(f"{BACKEND_URL}/graph/delete/node", json=payload,  timeout=5)
    assert fail_response.status_code == 404, f"Expected 404 but got {fail_response.status_code}"

def test_correctlyRemovesGraphEdge(graph_with_one_edge_response):
    payload = {"graphname": graph_with_one_edge_response["name"], "edge": graph_with_one_edge_response["edges"][0]}
    response = requests.post(f"{BACKEND_URL}/graph/delete/edge", json=payload,  timeout=5)
    assert response.status_code == 200, f"Expected 200 but got {response.status_code}"
    graph = response.json()
    assert len(graph['edge_map']) == 0, f"Expected 0 keys in the edge map but got {len(graph['edge_map'])}"
    assert len(graph['edges']) == 0, f"Expected 0 edges in the edge list but got {len(graph['edges'])}"
    assert len(graph['node_map']) == 2, f"Expected 2 keys in the node map but got {len(graph['node_map'])}"
    assert len(graph['nodes']) == 2, f"Expected 2 nodes in the nodes list but got {len(graph['nodes'])}"
    assert response.headers["Content-Type"].startswith("application/json")

def test_removeEdgeFromInvalidGraph(graph_with_one_edge_response):
    edge = graph_with_one_edge_response['nodes'][0]
    payload = {"graphname": "Invalid", "edge": edge}
    response = requests.post(f"{BACKEND_URL}/graph/delete/edge", json=payload,  timeout=5)    
    assert response.status_code == 404, f"Expected 404 but got {response.status_code}"

def test_removeFakeEdgefromGraph(graph_with_one_edge_response):
    fakeEdge = {'source': {'user_id': 'X', 'First_name': 'X', 'Last_name': 'X', 'Profile_picture_url': '', 'Email': 'X'}, 'target': {'user_id': 'Y', 'First_name': 'Y', 'Last_name': 'Y', 'Profile_picture_url': '', 'Email': 'Y'}, 'edge_tag': 'X', 'edge_desc': 'X'}
    payload = {"graphname": graph_with_one_edge_response['name'], "edge": fakeEdge}
    response = requests.post(f"{BACKEND_URL}/graph/delete/edge", json=payload,  timeout=5)    
    assert response.status_code == 404, f"Expected 404 but got {response.status_code}"

def test_removeEdgeFromEmptyGraphName(graph_with_one_edge_response):
    edge = graph_with_one_edge_response['edges'][0]
    payload = {"graphname": "", "edge": edge}
    response = requests.post(f"{BACKEND_URL}/graph/delete/edge", json=payload,  timeout=5)    
    assert response.status_code == 404, f"Expected 404 but got {response.status_code}"

def test_removeEdgeInvalidJson():
    response = requests.post(f"http://localhost:6767/graph/delete/edge", data="Invalid Json", headers={"Content-Type":"application/json"}, timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"

def test_removesSameEdgeTwice(graph_with_one_edge_response):
    graph = graph_with_one_edge_response
    edge = graph['edges'][0]
    payload = {"graphname": graph["name"], "edge": edge}
    successful_response = requests.post(f"{BACKEND_URL}/graph/delete/edge", json=payload,  timeout=5)
    assert successful_response.status_code == 200, f"Expected 200 but got {successful_response.status_code}"

    fail_response = requests.post(f"{BACKEND_URL}/graph/delete/edge", json=payload,  timeout=5)
    assert fail_response.status_code == 404, f"Expected 404 but got {fail_response.status_code}"

def test_removeSelfEdge(graph_with_one_edge_response):
    node = graph_with_one_edge_response['nodes'][0]
    fakeEdge = {'source': node, 'target': node, 'edge_tag': 'X', 'edge_desc': 'X'}
    payload = {"graphname": graph_with_one_edge_response["name"], "edge": fakeEdge}
    response = requests.post(f"{BACKEND_URL}/graph/delete/edge", json=payload,  timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"

def test_removeEdgeFakeSourceNode(graph_with_one_edge_response):
    fakeNode = {"user_id": "X", "First_name" : "X" , "Last_name": "X", "Email" :"X", "Profile_picture_url": ""}
    edge = {'source': fakeNode, 'target': graph_with_one_edge_response['nodes'][0], 'edge_tag': 'Family', 'edge_desc': 'From Highschool'}
    payload = {"graphname": graph_with_one_edge_response["name"], "edge": edge}
    response = requests.post(f"{BACKEND_URL}/graph/delete/edge", json=payload,  timeout=5)
    assert response.status_code == 404, f"Expected 404 but got {response.status_code}"

def test_removeEdgeFakeTargetNode(graph_with_one_edge_response):
    fakeNode = {"user_id": "X", "First_name" : "X" , "Last_name": "X", "Email" :"X", "Profile_picture_url": ""}
    edge = {'source': graph_with_one_edge_response['nodes'][0], 'target': fakeNode, 'edge_tag': 'Family', 'edge_desc': 'From Highschool'}
    payload = {"graphname": graph_with_one_edge_response["name"], "edge": edge}
    response = requests.post(f"{BACKEND_URL}/graph/delete/edge", json=payload,  timeout=5)
    assert response.status_code == 404, f"Expected 404 but got {response.status_code}"