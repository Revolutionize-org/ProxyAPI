import requests
import json

headers = {
    "X-API-KEY": "key",
}

payload = {
    "Url": "https://reqres.in/api/users?page=2",
    "Method": "GET",
    "Header": {
        "Content-Type": "application/json"
    }
}


response = requests.post(
    "http://localhost:8899/v1/send-request", headers=headers, json=payload)
if (response.status_code == 200):
    resp = json.loads(response.text)
    print(resp.get("Data").get("Body"))
else:
    print(response.text)

    print(response.status_code)
