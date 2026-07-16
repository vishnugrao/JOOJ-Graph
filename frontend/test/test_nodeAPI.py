import pytest
import requests

BACKEND_URL = "http://localhost:6767"

def test_ServerReachability(add_node_response_raw):
    assert add_node_response_raw.status_code != 404, f"Server not reachable, got {add_node_response_raw.status_code}"

def test_CorrectPort(add_node_response_raw):
    try:
        assert add_node_response_raw.status_code != None
    except requests.exceptions.ConnectionError:
        pytest.fail("Could not connect to the port 6767")


def test_TimeOutHandling(add_node_response_raw):
    try:
        assert add_node_response_raw.status_code != None
    except requests.exceptions.Timeout:
        pytest.fail("Server timed out after 5 seconds")

def test_nonEmptyResponse(add_node_response_raw):
    assert add_node_response_raw.text != ""

def test_ServerDown():
    try:
        response = requests.get("http://localhost:1403/create/node", timeout=3)
        pytest.fail("Expected connection error but got a response")
    except requests.exceptions.ConnectionError:
        pass


def test_JsonPayloadPost(add_node_response_raw):
    assert add_node_response_raw.status_code == 201, f"Expected 201 but got {add_node_response_raw.status_code}"


def test_correctHeaders(add_node_response_raw):
    assert add_node_response_raw.status_code == 201, f"Expected 201 but got {add_node_response_raw.status_code}"
    assert "Content-Type" in add_node_response_raw.request.headers
    assert add_node_response_raw.request.headers["Content-Type"] == "application/json"


def test_emptyPayload():
    response = requests.post(f"{BACKEND_URL}/create/node", data="",headers={"Content-Type":"application/json"} ,timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"


def test_invalidJson():
    response = requests.post(f"{BACKEND_URL}/create/node", data="Invalid Json", headers={"Content-Type":"application/json"}, timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"
    

def test_statusCodes(add_node_response_raw, node_payload):
    assert add_node_response_raw.status_code == 201, f"Expected 201 but got {add_node_response_raw.status_code}"

    response2 = requests.post(f"{BACKEND_URL}/create/node/dummy", json=node_payload, timeout=3)
    assert response2.status_code == 404, f"Expected 404 but got {response2.status_code}"

    response3 = requests.get(f"{BACKEND_URL}/create/node")
    assert response3.status_code == 400, f"Expected 400 but got {response3.status_code}"


def test_isString(add_node_response_raw):
    data = add_node_response_raw.text
    assert add_node_response_raw.status_code == 201, f"Expected 201 but got {add_node_response_raw.status_code}"
    assert isinstance(data, str), f"Expected string but got {type(data)}"

def test_nonEmptyUserID(add_node_response):
    assert add_node_response["user_id"] != "", f"Expected a generated user ID but got {add_node_response["user_id"]}"

def test_nonEmptyCreatedAtDate(add_node_response):
    assert add_node_response["Created_at"] != "", f"Expected a Created At date but got {add_node_response["Created_at"]}"

def test_userIDisStr(add_node_response):
    assert isinstance(add_node_response["user_id"], str), f"Expected generated user ID to be a string but got {type(add_node_response["user_id"])}"

def test_correctPostedFields(add_node_response):
    assert add_node_response["First_name"] == "John", f"Expected John but got {add_node_response["First_name"]}"
    assert add_node_response["Last_name"] == "Smith", f"Expected Smith but got {add_node_response["Last_name"]}"
    assert add_node_response["Email"] == "Jsmith@gmail.com", f"Expected Jsmith@gmail.com but got {add_node_response["Email"]}"

def test_invalidFields(invalid_node_payloads):
    for i, payload in enumerate(invalid_node_payloads):
        response = requests.post(f"{BACKEND_URL}/create/node", json=payload, timeout=3)
        assert response.status_code == 400, F"Expected 400 from payload {i} but got {response.status_code}"

def test_addNodeToGraph(graph_with_one_node_response):
    assert len(graph_with_one_node_response['nodes']) == 1, f"Expected 1 but got {len(graph_with_one_node_response['nodes'])}"

def test_addTwoNodesToGraph(add_graph_response):
    node1 = {"first_name": "John", "last_name": "Smith", "email": "Jsmith@gmail.com"}
    node2 = {"first_name": "Jay", "last_name": "Beck", "email": "Jbeck@gmail.com"}

    nodeout1 = requests.post(f"{BACKEND_URL}/create/node", json=node1, timeout=5).json()
    nodeout2 = requests.post(f"{BACKEND_URL}/create/node", json=node2, timeout=5).json()

    payload1 = {"graphname" : 'SJ' , "node" : nodeout1}
    payload2 = {"graphname" : 'SJ' , "node" : nodeout2}

    requests.post(f"{BACKEND_URL}/graph/add/node", json=payload1,  timeout=5)
    requests.post(f"{BACKEND_URL}/graph/add/node", json=payload2,  timeout=5)

    graph = requests.get(f"{BACKEND_URL}/get/graph?name=SJ", json=payload1,  timeout=5).json()
    assert len(graph['nodes']) == 2, f"Expected 2 but got {len(graph['nodes'])}"

def test_addNodeCorrectNodeMapUpdate(graph_with_one_node_response):
    nodeID = graph_with_one_node_response['nodes'][0]['user_id']
    assert nodeID in graph_with_one_node_response['node_map']

def test_addNodeToNoGraph(add_node_response):
    payload = {"graphname" : 'DNE' , "node" : add_node_response}
    response = requests.post(f"{BACKEND_URL}/graph/add/node", json=payload,  timeout=5)
    assert response.status_code == 404, f"Expected 404 but got {response.status_code}"

def test_addNodeToEmptyGraph(add_node_response):
    payload = {"graphname" : '' , "node" : add_node_response}
    response = requests.post(f"{BACKEND_URL}/graph/add/node", json=payload,  timeout=5)
    assert response.status_code == 404, f"Expected 404 but got {response.status_code}"

def test_duplicateAddNode(graph_with_one_node_response, add_node_response):
    payload = {'graphname': 'SJ', 'node': add_node_response}
    response = requests.post(f"{BACKEND_URL}/graph/add/node", json=payload,  timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"

def test_updatedGraphReturned(graph_with_one_node_response):
    assert graph_with_one_node_response['name'] == 'SJ'
    assert isinstance(graph_with_one_node_response['nodes'], list)
    assert graph_with_one_node_response['nodes'][0]['Email'] == 'Jsmith@gmail.com'