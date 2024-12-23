This is the backend exercise on creating a receipt api for Fetch Rewards.

To run this project first clone the project into a directory of your choice.
Then from a terminal, cd into cmd/api:
```bash
cd cmd/api
```

Then run:
```
go build
go run main.go
```

This will build the project and install dependencies, then start the server.
After the server is running, you can make a POST request to the api at http://localhost:8080/receipts/process to add a receipt.
You can also make a GET request to the API at http://localhost:8080/receipts/{id}/points to get the points based on receipt id.
