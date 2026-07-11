import requests
import os
import networkx as nx
from pyvis.network import Network
import webbrowser
import tempfile
BACKEND_URL = os.environ.get("BACKEND_URL", "http://localhost:6767")


def createGraph():
    name = input("\nEnter a name: ").strip()
    payload = {"name": name}
    try:
        response = requests.post(f"{BACKEND_URL}/create/graph", json=payload, timeout=5)
        response.raise_for_status()
        graph = response.json()
        return graph
    except requests.exceptions.RequestException as e:
        print(f"\nFailed to create graph see: {e}")

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
        print(f"\nFailed to create node see: {e}")


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
    print(f"Opening: file://{os.path.abspath(temp_path)}")
    webbrowser.open(f"file://{os.path.abspath(temp_path)}")

def show_graph(local_store):
    print("\nAvailable Graphs:")
    for key in local_store:
        print(key)
    name = input("\nWhich graph do you want to see? ")
    if name not in local_store:
        print(f"\nGraph {name} does not exist, please try again...")
        return
    if len(local_store[name]["nodes"]) == 0:
        print(f"Graph {name} has no nodes")
        return
    render_graph(local_store[name])

def tuiScreen():
    local_store = {}
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
                    local_store[graph["name"]] = graph
                    print("\nGraph has been created", local_store)

            elif opt == "node":
                if len(local_store) == 0:
                    print("Graph does not exist, please create a graph first and try again...")
                    continue
                else:
                    print("\nAvailable Graphs:")
                    for key in local_store:
                        print(key)

                    name = input("\nWhich graph do you want to add a node to?: ")
                    if name in local_store:
                        graph = local_store[name]
                        node = createNode()
                    
                        existing_emails = {n['Email'] for n in graph["nodes"]}
                        if node["Email"] in existing_emails:
                            print("\nThis node already exists...")
                            continue

                        graph["nodes"].append(node)
                        graph["node_map"][node["user_id"]] = True
                        print("\nNode has been created", local_store)
                    else:
                        print(f"\n Graph {name} does not exist...")


            elif opt == "edge": print("\nEdge has been created")
            elif opt == "back": print("\nGoing back to menu...")
            else: print("\n Try again...")

        elif user =="read":
            opt = input("\nWhat do you want to read? enter 'graph'/'node'/'edge' or 'back' to go back to menu: ")
            if opt == "graph":
                show_graph(local_store) 
                print("\nGraph has been read")

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