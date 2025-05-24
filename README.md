# Address Book Service

A modular address book service with in-memory storage.

## Requirements
- Go 1.21 or higher

## Installation
1. Clone the repository
2. Run `go mod download` to install dependencies
3. Run `go run main.go` to start the server

## API Endpoints
- POST `/create` - Create a new contact

    Sample Curl - `curl -X POST http://localhost:5000/create \
  -H "Content-Type: application/json" \
  -d '[{"name":"Alice Smith","phone":"1234567890","email":"alice@example.com"},
       {"name":"Bob Jones","phone":"2345678901","email":"bob@example.com"}]'`
- PUT `/update` - Update an existing contact

    Sample Curl - `curl -X PUT http://localhost:5000/update \
  -H "Content-Type: application/json" \
  -d '[{"id":"<UUID1>","phone":"9999999999"},
       {"id":"<UUID2>","email":"newbob@example.com"}]'`
- DELETE `/delete` - Delete one or more contacts

    Sample Curl - `curl --location --request DELETE 'http://localhost:5000/delete' \
    --header 'Content-Type: application/json' \
    --data '["4d696b62-5ed6-4b58-bb3d-fde66a694184"]'`
- POST `/search` - Search contacts by query

    Sample Curl - `curl --location 'http://localhost:5000/search' \
    --header 'Content-Type: application/json' \
    --data '{"query": "jones"}'`

The server will run on `http://localhost:5000`