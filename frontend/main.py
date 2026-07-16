import requests
import os
import networkx as nx
from pyvis.network import Network
import webbrowser
import tempfile
BACKEND_URL = os.environ.get("BACKEND_URL", "http://localhost:6767")

# def render_graph(graph_json):
#     net = Network(directed=True, height="750px", width="100%", bgcolor="#ffffff")

#     for node in graph_json.get("nodes", []):
#         if "user_id" not in node:
#             continue
#         net.add_node(
#             node["user_id"],
#             label=f"{node['First_name']} {node['Last_name']}",
#             title = node["Email"]
#         )
    
#     for edge in graph_json.get("edges", []):
#         source = edge.get("Source")
#         target = edge.get("Target")

#         if source == target: continue

#         net.add_edge(source, target,
#             edge_tag = edge.get("Edge_tag"),
#             edge_desc = edge.get("Edge_desc"))
#     with tempfile.NamedTemporaryFile(suffix='.html', delete=False, mode='w') as f:
#         f.write(net.generate_html())
#         temp_path = f.name
#     print(f"Opening: file://{os.path.abspath(temp_path)}")
#     webbrowser.open(f"file://{os.path.abspath(temp_path)}")

# def show_graph(local_store):
#     print("\nAvailable Graphs:")
#     for key in local_store:
#         print(key)
#     name = input("\nWhich graph do you want to see? ")
#     if name not in local_store:
#         print(f"\nGraph {name} does not exist, please try again...")
#         return
#     if len(local_store[name]["nodes"]) == 0:
#         print(f"Graph {name} has no nodes")
#         return
#     render_graph(local_store[name])

def createGraph():
    name = input("\nEnter a name: ").strip()
    payload = {"name": name}
    try:
        response = requests.post(f"{BACKEND_URL}/create/graph", json=payload, timeout=5)
        response.raise_for_status()
        graph = response.json()
        return graph
    except requests.exceptions.RequestException as e:
        if response.status_code == 409: print("Graph with this name already exists, please use another name...")
        elif response.status_code == 400: print("Unexpected error occurred, please try again...")
        else: print(f"\nFailed to create graph see: {e}")

def createNode():
    first = input("\nEnter first name: ").strip()
    last = input("\nEnter last name: ").strip()
    email = input("\nEnter email name: ").strip()
    payload = {"first_name": first, "last_name": last, "email":email}
    try:
        response = requests.post(f"{BACKEND_URL}/create/node", json=payload, timeout=5)
        response.raise_for_status()
        node = response.json()
        return node
    except requests.exceptions.RequestException as e:
        if response.status_code == 400: print("Unexpected error occurred, please try again...")
        else: print(f"\nFailed to create node see: {e}")

def returnAllGraphs():
    try:
        response = requests.get(f"{BACKEND_URL}/get/allgraphs", timeout=5)
        return response.json()
    except requests.exceptions.RequestException as e:
        if response.status_code == 400: print("try again...")
        else: print(f"\n Failed to fetch all graphs see: {e}")

def selectGraph():
    graphlist = returnAllGraphs()
    try:
        print("\nAvailable Graphs")
        for key in graphlist:
            print(key)
        name = input("\nWhich graph do you want to select?: ").strip()

        response = requests.get(f"{BACKEND_URL}/get/graph?name={name}", timeout=5)
        graph = response.json()
        return graph
    except requests.exceptions.RequestException as e:
        if response.status_code == 400: print("try again...")
        elif response.status_code == 404: print("Graph does not exist, please try again...")
        else: print(f"\n Failed to select graph see: {e}")

def addNodeToGraph(name, node):
    payload = {"graphname" : name , "node" : node}
    try:
        response = requests.post(f"{BACKEND_URL}/graph/add/node", json=payload,  timeout=5)
    except requests.exceptions.RequestException as e:
        if response.status_code == 201: print(f"Node has been added to Graph {name}")
        elif response.status_code == 400: print(f"Failed to add Node to Graph {name}")
        elif response.status_code == 404: print(f"Graph {name} does not exist")
        else: print(f"Unexpected error see: {e}")

def createEdge(graph):
    nodes = graph['nodes']
    sourceIdx = input("\nEnter source node number: ").strip()
    targetIdx = input("\nEnter target node number: ").strip()

    try:
        source = nodes[int(sourceIdx)-1]
        target = nodes[int(targetIdx)-1]
    except (IndexError, ValueError):
        print("\nInvalid node selected")
        return None
    
    if source['user_id'] == target['user_id']:
        print("\nSource and Target cannot be the same node, please try again...")
        return None

    tag = input("\nEnter a tag (eg. Friend/Colleague, Family): ").strip()
    desc = input("\nEnter a description of your relationship: ").strip()

    payload = {"source": source, "target": target, "edge_tag": tag, "edge_desc": desc}
    try:
        response = requests.post(f"{BACKEND_URL}/create/edge", json=payload, timeout=5)
        response.raise_for_status()
        edge = response.json()
        return edge
    except requests.exceptions.RequestException as e:
        print(f"\nFailed to create edge see: {e}") 
        return None


def tuiScreen():
    while True:
        print("\n---------------------------------------------------\nWelcome to Jooj\nYour throwaway CLI\n---------------------------------------------------\n")
        user = input("\nActions are 'create'/'read'/'update'/'delete' or to quit enter 0: ")

        if user == "create":

            opt = input("\nWhat do you want to create? enter 'graph'/'node'/'edge' or 'back' to go back to menu: ")

            if opt == "back":
                continue

            elif opt == "graph":
                graph = createGraph() 
                if graph:
                    print("\nGraph has been created")
                    print(selectGraph())
            elif opt == "node":
                allgraph = returnAllGraphs()
                if len(allgraph) == 0:
                    print("Graph does not exist, please create a graph first and try again...")
                    continue
                else:
                    graph = selectGraph()
                    node = createNode()
                
                    existing_emails = {n['Email'] for n in graph["nodes"]}
                    if node["Email"] in existing_emails:
                        print("\nThis node already exists...")
                        continue
                    
                    addNodeToGraph(graph['name'], node)
                    print("\nNode has been created")
                    print(selectGraph())
                        

            elif opt == "edge":
                # graph = selectGraph()

                # if len(graph['nodes']) < 2:
                #     print("Graph must have 2 or more nodes, please try again...")
                #     continue

                # print("\nAvailable Nodes")
                # for i, node in enumerate(graph['nodes']):
                #     print(f"{i + 1}. {node['First_name']} {node['Last_name']} {node['Email']}")

                # edge = createEdge(graph)
                # edge_id = f"{edge['source']['user_id']}-{edge['target']['user_id']}"
                # if edge and edge_id not in graph['edge_map']:
                #     graph['edges'].append(edge)
                #     graph['edge_map'][edge_id] = True
                #     print("\nEdge has been created")
                print("\nEdge has been created")
            elif opt == "back": print("\nGoing back to menu...")
            else: print("\n Try again...")

        elif user =="read":
            opt = input("\nWhat do you want to read? enter 'graph'/'node'/'edge' or 'back' to go back to menu: ")
            if opt == "graph": print("\nGraph has been read")

            elif opt == "node": print("\nNode has been read")
            elif opt == "edge": print("\nEdge has been read")
            elif opt == "back": print("\nGoing back to menu...")
            else: print("\n Try again...")

        elif user =="update":
            opt = input("\nWhat do you want to update? enter 'graph'/'node'/'edge' or 'back' to go back to menu: ")
            if opt == "graph": print("\nGraph has been updated")
            elif opt == "node": print("\nNode has been updated")
            elif opt == "edge": print("\nEdge has been updated")
            elif opt == "back": print("\nGoing back to menu...")
            else: print("\n Try again...")

        elif user =="delete":
            opt = input("\nWhat do you want to delete? enter 'graph'/'node'/'edge' or 'back' to go back to menu: ")
            if opt == "graph": print("\nGraph has been deleted")
            elif opt == "node": print("\nNode has been deleted")
            elif opt == "edge": print("\nEdge has been deleted")
            elif opt == "back": print("\nGoing back to menu...")
            else: print("\n Try again...")

        elif user =="0": break
        else: print("\nPlease try again...")


def main():
    tuiScreen()


if __name__ == "__main__":
    main()