import pytest
import requests
import json

url = "http://localhost:6767"


def test_ServerReachability():
    payload = {"first_name": "John", "last_name": "Smith", "email": "Jsmith@gmail.com"}
    response = requests.post(f"{url}/input/dump", json=payload, timeout=5)
    assert response.status_code != 404, f"Server not reachable, got {response.status_code}"

def test_CorrectPort():
    try:
        payload = {"first_name": "John", "last_name": "Smith", "email": "Jsmith@gmail.com"}
        response = requests.post(f"{url}/input/dump", json=payload, timeout=5)
        assert response.status_code != None
    except requests.exceptions.ConnectionError:
        pytest.fail("Could not connect to the port 6767")


def test_TimeOutHandling():
    try:
        payload = {"first_name": "John", "last_name": "Smith", "email": "Jsmith@gmail.com"}
        response = requests.post(f"{url}/input/dump", json=payload, timeout=5)
        assert response.status_code != None
    except requests.exceptions.Timeout:
        pytest.fail("Server timed out after 5 seconds")

def test_ServerDown():
    try:
        response = requests.get("http://localhost:1403/input/dump", timeout=3)
        pytest.fail("Expected connection error but got a response")
    except requests.exceptions.ConnectionError:
        pass

def test_JsonPayloadPost():
    payload = {"first_name": "John", "last_name": "Smith", "email": "Jsmith@gmail.com"}
    response = requests.post(f"{url}/input/dump", json=payload, timeout=5)
    assert response.status_code == 201, f"Expected 201 but got {response.status_code}"

def test_correctHeaders():
    payload = {"first_name": "John", "last_name": "Smith", "email": "Jsmith@gmail.com"}
    response = requests.post(f"{url}/input/dump", json=payload, timeout=5)
    assert response.status_code == 201, f"Expected 201 but got {response.status_code}"
    assert "Content-Type" in response.request.headers
    assert response.request.headers["Content-Type"] == "application/json"


# Empty body POST — what happens if you send an empty payload
def test_emptyPayload():
    response = requests.post(f"{url}/input/dump", data="",headers={"Content-Type":"application/json"} ,timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"

# Invalid JSON POST — what happens if the JSON is malformed
def test_invalidJson():
    response = requests.post(f"{url}/input/dump", data="Invalid Json", headers={"Content-Type":"application/json"}, timeout=5)
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"
    

# Response Tests

# Status codes — is the server returning the right codes e.g. 200, 400, 404, 500
def test_statusCodes():
    payload = {"first_name": "John", "last_name": "Smith", "email": "Jsmith@gmail.com"}
    response = requests.post(f"{url}/input/dump", json=payload, timeout=5)
    assert response.status_code == 201, f"Expected 201 but got {response.status_code}"

    response2 = requests.post(f"{url}/input/dummy2/dump", json=payload, timeout=3)
    assert response2.status_code == 404, f"Expected 404 but got {response2.status_code}"

    response3 = requests.get(f"{url}/input/dump")
    assert response3.status_code == 400, f"Expected 400 but got {response3.status_code}"


# Response is a string — is the GET response returning readable text
def test_isString():
    payload = {"first_name": "John", "last_name": "Smith", "email": "Jsmith@gmail.com"}
    response = requests.post(f"{url}/input/dump", json=payload, timeout=5)
    data = response.text
    assert response.status_code == 201, f"Expected 201 but got {response.status_code}"
    assert isinstance(data, str), f"Expected string but got {type(data)}"

# Error Handling Tests

# 404 on wrong route — does hitting a non existent route return 404
def test_wrongRoute():
    response = requests.post(f"{url}/input/dummy/dump")
    assert response.status_code == 404, f"Expected 404 but got: {response.status_code}"

# Wrong method — what happens if you send a GET to a POST only endpoint
def test_wrongMethod():
    response = requests.get(f"{url}/input/dump")
    assert response.status_code == 400, f"Expected 400 but got {response.status_code}"