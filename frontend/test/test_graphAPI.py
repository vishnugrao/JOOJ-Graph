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