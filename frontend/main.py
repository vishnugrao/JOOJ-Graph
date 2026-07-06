import requests
import os
BACKEND_URL = os.environ.get("BACKEND_URL", "http://localhost:6767")


def createExampleNode():
    first = input("\nEnter first name: ").strip()
    last = input("\nEnter last name: ").strip()
    email = input("\nEnter email: ").strip()

    payload = {"first_name": first, "last_name": last, "email": email}
    try:
        response = requests.post(f"{BACKEND_URL}/input/dump", json=payload, timeout=5)
        response.raise_for_status()
        node = response.json()
        print(f"\nHere is your node: {node}")
    except requests.exceptions.RequestException as e:
        print(f"\nFailed to create node see: {e}")

def tuiScreen():

    print("\n---------------------------------------------------\nPress 1 to print Hello to console\nPress 0 to make a Node\nPress control+c to quit\n---------------------------------------------------\n")
    user = input("\n1 or 0:")
    if user == "1": print("\nHello User")
    elif user =="0": 
        createExampleNode()

def main():
    while True:
        tuiScreen()



if __name__ == "__main__":
    main()

    