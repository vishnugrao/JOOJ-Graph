import pytest
import requests
import json
BACKEND_URL = "http://localhost:6767"

def test_ServerReachability(add_edge_response_raw):
    assert add_edge_response_raw.status_code != 404, f"Server not reachable, got {add_edge_response_raw.status_code}"

def test_CorrectPort(add_edge_response_raw):
    try:
        assert add_edge_response_raw.status_code != None
    except requests.exceptions.ConnectionError:
        pytest.fail("Could not connect to the port 6767")


def test_TimeOutHandling(add_edge_response_raw):
    try:
        assert add_edge_response_raw.status_code != None
    except requests.exceptions.Timeout:
        pytest.fail("Server timed out after 5 seconds")


def test_ServerDown():
    try:
        response = requests.get("http://localhost:1403/create/edge", timeout=3)
        pytest.fail("Expected connection error but got a response")
    except requests.exceptions.ConnectionError:
        pass

def test_nonEmptyResponse(add_edge_response_raw):
    assert add_edge_response_raw.text != ""

def test_JsonPayloadPost(add_edge_response_raw):
    assert add_edge_response_raw.status_code == 201, f"Expected 201 but got {add_edge_response_raw.status_code}"


def test_correctHeaders(add_edge_response_raw):
    assert add_edge_response_raw.status_code == 201, f"Expected 201 but got {add_edge_response_raw.status_code}"
    assert "Content-Type" in add_edge_response_raw.request.headers
    assert add_edge_response_raw.request.headers["Content-Type"] == "application/json"


def test_emptyPayload():
    response = requests.post(f"http://localhost:6767/create/edge", data="",headers={"Content-Type":"application/json"} ,timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"


def test_invalidJson():
    response = requests.post(f"http://localhost:6767/create/edge", data="Invalid Json", headers={"Content-Type":"application/json"}, timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"
    

def test_statusCodes(add_edge_response_raw, edge_payload):
    assert add_edge_response_raw.status_code == 201, f"Expected 201 but got {add_edge_response_raw.status_code}"

    response2 = requests.post(f"http://localhost:6767/create/edge/dummy", json=edge_payload, timeout=3)
    assert response2.status_code == 404, f"Expected 404 but got {response2.status_code}"

    response3 = requests.get(f"http://localhost:6767/create/edge")
    assert response3.status_code == 400, f"Expected 400 but got {response3.status_code}"


def test_isString(add_edge_response_raw):
    data = add_edge_response_raw.text
    assert add_edge_response_raw.status_code == 201, f"Expected 201 but got {add_edge_response_raw.status_code}"
    assert isinstance(data, str), f"Expected string but got {type(data)}"

def test_correctPostedFields(add_edge_response, graph_withnodes_payload):
    assert add_edge_response["source"] == graph_withnodes_payload['SJ']['nodes'][0], f"Expected {graph_withnodes_payload['SJ']['nodes'][0]} but got {add_edge_response['source']}"
    assert add_edge_response["target"] == graph_withnodes_payload['SJ']['nodes'][1], f"Expected {graph_withnodes_payload['SJ']['nodes'][1]} but got {add_edge_response['target']}"
    assert add_edge_response["edge_tag"] == "Friend", f"Expected Friend but got {add_edge_response["edge_tag"]}"
    assert add_edge_response["edge_desc"] == "From Highschool", f"Expected From Highschool but got {add_edge_response["edge_desc"]}"

def test_invalidFields(invalid_edge_payloads):
    for i, payload in enumerate(invalid_edge_payloads):
        response = requests.post(f"http://localhost:6767/create/edge", json=payload, timeout=3)
        assert response.status_code == 400, F"Expected 400 from payload {i} but got {response.status_code}"

def test_correctEdgeFieldTypes(add_edge_response):
    assert add_edge_response['source'] != None
    assert add_edge_response['target'] != None
    assert isinstance(add_edge_response['edge_tag'], str)
    assert isinstance(add_edge_response['edge_desc'], str)

def test_selfEdgeCase(graph_with_one_edge_response):
    graph = requests.get(f"{BACKEND_URL}/get/graph?name={graph_with_one_edge_response['name']}",  timeout=5).json()
    edge_payload = {"source" : graph['nodes'][0], "target" : graph['nodes'][0], "edge_tag": "Myself", "edge_desc" : "Birth Certificate"}
    response = requests.post(f"{BACKEND_URL}/create/edge", json=edge_payload, timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"