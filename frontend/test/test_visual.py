from pyvis.network import Network
import networkx as nx
import pytest
import os
import webbrowser
import tempfile
from unittest.mock import patch

def build_graph(graph_json):
    graph = nx.DiGraph()

    for node in graph_json.get("nodes", []):
        if 'user_id' not in node: continue
        
        graph.add_node(
            node["user_id"],
            label=f"{node.get('First_name', '')} {node.get('Last_name', '')}",
            email = node.get("Email", ''),
            created_at = node.get("Created_at", ''),
            profile_picture_url = node.get("Profile_picture_url", ''))
    
    for edge in graph_json.get("edges", []):
        source = edge.get("Source")
        target = edge.get("Target")

        if source == target: continue

        graph.add_edge(source, target,
            edge_tag = edge.get("Edge_tag"),
            edge_desc = edge.get("Edge_desc"))
    return graph


def render_graph(graph_json):
    net = Network(directed=True, height="750px", width="100%", bgcolor="#ffffff")

    for node in graph_json.get("nodes", []):
        if "user_id" not in node:
            continue
        net.add_node(
            node["user_id"],
            label=f"{node['First_name']} {node['Last_name']}",
            title = node["Email"]
        )
    
    for edge in graph_json.get("edges", []):
        source = edge.get("Source")
        target = edge.get("Target")

        if source == target: continue

        net.add_edge(source, target,
            edge_tag = edge.get("Edge_tag"),
            edge_desc = edge.get("Edge_desc"))
    with tempfile.NamedTemporaryFile(suffix='.html', delete=False, mode='w') as f:
        f.write(net.generate_html())
        temp_path = f.name
    webbrowser.open(f"file://{os.path.abspath(temp_path)}")


# Edge testcases

def test_successfulGraphRender(graph_withnodes_payload):
    with patch("webbrowser.open"):
        with patch("pyvis.network.Network.generate_html", return_value="<html></html>"):
            try:
                render_graph(graph_withnodes_payload["SJ"])
            except Exception as e:
                pytest.fail(f"The rendered graph crashed see: {e}")

def test_emptyGraphRender(add_graph_response):
    with patch("webbrowser.open"):
        with patch("pyvis.network.Network.generate_html", return_value="<html></html>"):
            try:
                render_graph(add_graph_response)
            except Exception as e:
                pytest.fail(f"The rendered empty graph crashed see: {e}")

def test_graphNetworkCreation(graph_withnodes_payload):
    with patch("pyvis.network.Network.show"):
        net = Network(directed=True)
        for node in graph_withnodes_payload["SJ"]["nodes"]:
            net.add_node( node["user_id"], label=f"{node['First_name']} {node['Last_name']}", title = node["Email"])
        assert len(net.nodes) == len(graph_withnodes_payload["SJ"]["nodes"])

def test_correctNodeLabels(graph_withnodes_payload):
    with patch("pyvis.network.Network.show"):
        net = Network(directed=True)
        for node in graph_withnodes_payload["SJ"]["nodes"]:
            net.add_node( node["user_id"], label=f"{node['First_name']} {node['Last_name']}", title = node["Email"])
        labels = [n['label'] for n in net.nodes]
        assert "V Rao" in labels
        assert "Jay Lo" in labels

def test_correctNodeTitles(graph_withnodes_payload):
    with patch("pyvis.network.Network.show"):
        net = Network(directed=True)
        for node in graph_withnodes_payload["SJ"]["nodes"]:
            net.add_node( node["user_id"], label=f"{node['First_name']} {node['Last_name']}", title = node["Email"])
        titles = [n["title"] for n in net.nodes]
        assert "VRao@outlook.com" in titles
        assert "JLo@gmail.com" in titles

def test_correctGraphFilename(graph_withnodes_payload):
    with patch("webbrowser.open") as mock_open:
        with patch("pyvis.network.Network.generate_html", return_value="<html></html>"):
            render_graph(graph_withnodes_payload["SJ"])
            args, _ = mock_open.call_args
            assert args[0].startswith("file://")
            assert args[0].endswith(".html")


def test_graphJsonCreation(graph_withnodes_payload):
    graph = build_graph(graph_withnodes_payload["SJ"])
    assert isinstance(graph, nx.DiGraph)

def test_graphNodeCount(graph_withnodes_payload):
    graph = build_graph(graph_withnodes_payload["SJ"])
    assert graph.number_of_nodes() == len(graph_withnodes_payload["SJ"]["nodes"])

def test_graphEdgeCount(graph_withnodes_payload):
    graph = build_graph(graph_withnodes_payload["SJ"])
    assert graph.number_of_edges() == len(graph_withnodes_payload["SJ"]["edges"])

def test_emptyGraphRender(add_graph_response):
    graph = build_graph(add_graph_response)
    assert graph.number_of_nodes() == 0, f"Expected 0 but got {graph.number_of_nodes()}"
    assert graph.number_of_edges() == 0, f"Expected 0 but got {graph.number_of_edges()}"

def test_graphOnlyNodesRender(graph_withnodes_payload):
    graph = build_graph(graph_withnodes_payload["SJ"])
    assert graph.number_of_nodes() == len(graph_withnodes_payload["SJ"]["nodes"])
    assert graph.number_of_edges() == 0

def test_nodesStoredCorrectlyOnNx(graph_withnodes_payload):
    graph = build_graph(graph_withnodes_payload["SJ"])
    node_id = {node["user_id"] for node in graph_withnodes_payload["SJ"]["nodes"]}
    assert set(graph.nodes()) == node_id

def test_correctNodeLabels(graph_withnodes_payload):
    graph = build_graph(graph_withnodes_payload["SJ"])
    labels = nx.get_node_attributes(graph, "label")
    assert labels['f1d71d4b-a88c-4170-aa6b-eb8753bc2e89'] == "V Rao"
    assert labels['c11fc4fb-266a-4ae3-acdc-49fec05b85d9'] == "Jay Lo"

def test_correctlyStoredNodeData(graph_withnodes_payload):
    graph = build_graph(graph_withnodes_payload["SJ"])
    data = graph.nodes['f1d71d4b-a88c-4170-aa6b-eb8753bc2e89']
    assert data["email"] == "VRao@outlook.com"
    assert data["created_at"] == '2026-07-09T09:08:49Z'
    assert data['profile_picture_url'] == ""

def test_nodeMap_matchesNxGraphNodes(graph_withnodes_payload):
    graph = build_graph(graph_withnodes_payload["SJ"])
    for node_id in graph_withnodes_payload["SJ"]["node_map"]:
        assert graph.has_node(node_id)


def test_invalidpayload():
    with pytest.raises((KeyError, TypeError, AttributeError)):
        graph = nx.DiGraph()
        invalid_payload = "invalid"
        for node in invalid_payload['nodes']:
            graph.add_node(node['user_id'])