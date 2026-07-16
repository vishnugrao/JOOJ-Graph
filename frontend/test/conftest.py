import pytest
import requests
from unittest.mock import patch

BACKEND_URL = "http://localhost:6767"

@pytest.fixture
def node_payload():
    return  {"first_name": "John", "last_name": "Smith", "email": "Jsmith@gmail.com"}
@pytest.fixture
def invalid_node_payloads():
    return  [
        {"first_name": "J0hn", "last_name": "Smith", "email": "Jsmith@gmail.com"},
        {"first_name": "John", "last_name": "Sm1th", "email": "Jsmith@gmail.com"},
        {"first_name": "John", "last_name": "Smith", "email": "invalidemail"},
        {"first_name": " ", "last_name": "Smith", "email": "Jsmith@gmail.com"},
        {"first_name": "John", "last_name": "Smith", "email": " "},
        {"first_name": "", "last_name": "Smith", "email": "Jsmith@gmail.com"},
        {"first_name": "John", "last_name": "Smith", "email": ""}
    ]

@pytest.fixture
def invalid_edge_payloads(graph_withnodes_payload):
    return [
        {"source" : "", "target" : graph_withnodes_payload['SJ']['nodes'][1], "edge_tag": "Friend", "edge_desc" : "From Highschool"},
        {"source" : graph_withnodes_payload['SJ']['nodes'][0], "target" : "", "edge_tag": "Friend", "edge_desc" : "From Highschool"},
        {"source" : "graph_withnodes_payload['SJ']['nodes'][0]", "target" : graph_withnodes_payload['SJ']['nodes'][1], "edge_tag": "Friend", "edge_desc" : "From Highschool"},
        {"source" : graph_withnodes_payload['SJ']['nodes'][0], "target" : "graph_withnodes_payload['SJ']['nodes'][1]", "edge_tag": "Friend", "edge_desc" : "From Highschool"},
        {"source" : graph_withnodes_payload['SJ']['nodes'][0], "target" : graph_withnodes_payload['SJ']['nodes'][1], "edge_tag": "Fri092nd", "edge_desc" : "From Highschool"},
        {"source" : graph_withnodes_payload['SJ']['nodes'][0], "target" : graph_withnodes_payload['SJ']['nodes'][0], "edge_tag": "Friend", "edge_desc" : "From Highschool"},
        {"source" : graph_withnodes_payload['SJ']['nodes'][0], "target" : graph_withnodes_payload['SJ']['nodes'][1], "edge_tag": " ", "edge_desc" : "From Highschool"},
        {"source" : graph_withnodes_payload['SJ']['nodes'][0], "target" : graph_withnodes_payload['SJ']['nodes'][1], "edge_tag": "Friend", "edge_desc" : " "}
    ]

@pytest.fixture
def graph_payload():
    return  {"name": "SJ"}

@pytest.fixture
def graph_withnodes_payload():
    return {'SJ': {'name': 'SJ', 'nodes': [{'user_id': 'f1d71d4b-a88c-4170-aa6b-eb8753bc2e89', 'First_name': 'V', 'Last_name': 'Rao', 'Profile_picture_url': '', 'Email': 'VRao@outlook.com', 'Created_at': '2026-07-09T09:08:49Z'}, 
                                           {'user_id': 'c11fc4fb-266a-4ae3-acdc-49fec05b85d9', 'First_name': 'Jay', 'Last_name': 'Lo', 'Profile_picture_url': '', 'Email': 'JLo@gmail.com', 'Created_at': '2026-07-09T09:09:41Z'}], 
                    'edges': [], 'node_map': {'f1d71d4b-a88c-4170-aa6b-eb8753bc2e89': True, 'c11fc4fb-266a-4ae3-acdc-49fec05b85d9': True}, 'edge_map': {}}}

@pytest.fixture
def edge_payload(graph_withnodes_payload):
    return {"source" : graph_withnodes_payload['SJ']['nodes'][0], "target" : graph_withnodes_payload['SJ']['nodes'][1], "edge_tag": "Friend", "edge_desc" : "From Highschool"}


@pytest.fixture
def add_node_response(node_payload):
    # Returns: {'user_id': 'f1d71d4b-a88c-4170-aa6b-eb8753bc2e89', 'First_name': 'V', 'Last_name': 'Rao', 'Profile_picture_url': '', 'Email': 'VRao@outlook.com', 'Created_at': '2026-07-09T09:08:49Z'}
    response = requests.post(f"{BACKEND_URL}/create/node", json=node_payload, timeout=5)
    return response.json()

@pytest.fixture
def add_node_response_raw(node_payload):
    # Returns: {'user_id': 'f1d71d4b-a88c-4170-aa6b-eb8753bc2e89', 'First_name': 'V', 'Last_name': 'Rao', 'Profile_picture_url': '', 'Email': 'VRao@outlook.com', 'Created_at': '2026-07-09T09:08:49Z'}
    return requests.post(f"{BACKEND_URL}/create/node", json=node_payload, timeout=5)

@pytest.fixture
def add_graph_response_raw(graph_payload):
    # Returns: {'SJ': {'name': 'SJ', 'nodes': [], 'edges': [], 'node_map': {}, 'edge_map': {}}}
     return requests.post(f"{BACKEND_URL}/create/graph", json=graph_payload, timeout=5)

@pytest.fixture
def add_graph_response(graph_payload):
    # Returns: {'SJ': {'name': 'SJ', 'nodes': [], 'edges': [], 'node_map': {}, 'edge_map': {}}}
    response = requests.post(f"{BACKEND_URL}/create/graph", json=graph_payload, timeout=5)
    return response.json()

@pytest.fixture
def add_edge_response_raw(edge_payload):
    return requests.post(f"{BACKEND_URL}/create/edge", json=edge_payload, timeout=5)

@pytest.fixture
def add_edge_response(edge_payload):
    # Returns {'source': {'user_id': 'f1d71d4b-a88c-4170-aa6b-eb8753bc2e89', 'First_name': 'V', 'Last_name': 'Rao', 'Profile_picture_url': '', 'Email': 'VRao@outlook.com', 'Created_at': '2026-07-09T09:08:49Z'}, 'target': {'user_id': 'c11fc4fb-266a-4ae3-acdc-49fec05b85d9', 'First_name': 'Jay', 'Last_name': 'Lo', 'Profile_picture_url': '', 'Email': 'JLo@gmail.com', 'Created_at': '2026-07-09T09:09:41Z'}, 'edge_tag': 'Friend', 'edge_desc': 'From Highschool', 'created_at_edges': '2026-07-12T08:58:25Z'}
    response = requests.post(f"{BACKEND_URL}/create/edge", json=edge_payload, timeout=5)
    return response.json()

@pytest.fixture(autouse=True)
def reset_store():
    requests.post(f"{BACKEND_URL}/reset", timeout=3)
    yield

@pytest.fixture
def get_graph_response(add_graph_response):
    # Returns {'name': 'SJ', 'nodes': [], 'edges': [], 'node_map': {}, 'edge_map': {}}
    response = requests.get(f"{BACKEND_URL}/get/graph?name=SJ", timeout=5)
    return response.json()

@pytest.fixture
def get_graph_response_raw(add_graph_response):
    return requests.get(f"{BACKEND_URL}/get/graph?name=SJ", timeout=5)

@pytest.fixture
def get_all_graphs_response(add_graph_response):
    # Returns {'ADL': {'name': 'ADL', 'nodes': [], 'edges': [], 'node_map': {}, 'edge_map': {}}, 'SJ': {'name': 'SJ', 'nodes': [], 'edges': [], 'node_map': {}, 'edge_map': {}}}
    requests.post(f"{BACKEND_URL}/create/graph", json={"name":"ADL"}, timeout=5)
    response = requests.get(f"{BACKEND_URL}/get/allgraphs", timeout=5)
    return response.json()

@pytest.fixture
def get_all_graphs_response_raw(add_graph_response):
    requests.post(f"{BACKEND_URL}/create/graph", json={"name":"ADL"}, timeout=5)
    return requests.get(f"{BACKEND_URL}/get/allgraphs", timeout=5)

@pytest.fixture
def graph_with_one_node_response(get_graph_response, add_node_response):
    # Returns graph with 1 node fixture response: {'name': 'SJ', 'nodes': [{'user_id': '0ff6e83c-d9d5-42aa-b626-4668c0134757', 'First_name': 'John', 'Last_name': 'Smith', 'Profile_picture_url': '', 'Email': 'Jsmith@gmail.com', 'Created_at': '2026-07-16T07:15:40Z'}], 'edges': [], 'node_map': {'0ff6e83c-d9d5-42aa-b626-4668c0134757': True}, 'edge_map': {}}
    payload = {"graphname" : get_graph_response['name'] , "node" : add_node_response}
    addedNode = requests.post(f"{BACKEND_URL}/graph/add/node", json=payload,  timeout=5)
    if addedNode.status_code != 201: pytest.fail(f"Failed to add node to the graph")
    response = requests.get(f"{BACKEND_URL}/get/graph?name={get_graph_response['name']}",  timeout=5)
    return response.json()