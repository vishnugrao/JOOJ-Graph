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

# def test_duplicate_graph_names():
#     May need to make some sort of map on the main TUI to keep track of duplicate graph names

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

def test_debug_response(add_graph_response):
    print(add_graph_response)
