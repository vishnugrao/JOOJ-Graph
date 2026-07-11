import pytest
import requests
import json

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
    response = requests.post(f"http://localhost:6767/create/node", data="",headers={"Content-Type":"application/json"} ,timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"


def test_invalidJson():
    response = requests.post(f"http://localhost:6767/create/node", data="Invalid Json", headers={"Content-Type":"application/json"}, timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"
    

def test_statusCodes(add_node_response_raw, node_payload):
    assert add_node_response_raw.status_code == 201, f"Expected 201 but got {add_node_response_raw.status_code}"

    response2 = requests.post(f"http://localhost:6767/create/node/dummy", json=node_payload, timeout=3)
    assert response2.status_code == 404, f"Expected 404 but got {response2.status_code}"

    response3 = requests.get(f"http://localhost:6767/create/node")
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
    response = requests.post(f"http://localhost:6767/create/node", json=invalid_node_payloads[0], timeout=3)
    assert response.status_code == 400, F"Expected 400 but got {response.status_code}"

    response2 = requests.post(f"http://localhost:6767/create/node", json=invalid_node_payloads[1], timeout=3)
    assert response2.status_code == 400, F"Expected 400 but got {response2.status_code}"

    response3 = requests.post(f"http://localhost:6767/create/node", json=invalid_node_payloads[2], timeout=3)
    assert response3.status_code == 400, F"Expected 400 but got {response3.status_code}"

    response4 = requests.post(f"http://localhost:6767/create/node", json=invalid_node_payloads[3], timeout=3)
    assert response4.status_code == 400, F"Expected 400 but got {response4.status_code}"

    response5 = requests.post(f"http://localhost:6767/create/node", json=invalid_node_payloads[4], timeout=3)
    assert response5.status_code == 400, F"Expected 400 but got {response5.status_code}"

    response6 = requests.post(f"http://localhost:6767/create/node", json=invalid_node_payloads[5], timeout=3)
    assert response6.status_code == 400, F"Expected 400 but got {response6.status_code}"

    response7= requests.post(f"http://localhost:6767/create/node", json=invalid_node_payloads[6], timeout=3)
    assert response7.status_code == 400, F"Expected 400 but got {response7.status_code}"

def test_debug_response(add_node_response):
    print(add_node_response)