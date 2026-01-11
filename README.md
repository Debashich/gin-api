# go-albums-api

A minimal REST API for managing music albums, built with Go and the Gin framework.[web:1][web:11]

> Repo: https://github.com/Debashich/gin-api

---

## Features

- JSON-based REST API for albums (list, retrieve by ID, create).[web:1]
- In-memory data store with a simple `Album` model.
- Layered structure: HTTP handlers, model, and store separated for clarity.[web:1][web:14]
- Ready to extend with a real database or authentication.

---

## Project structure

```txt
gin-api/
  cmd/
    server/
      main.go        # Entrypoint: starts Gin and registers routes
  internal/
    albums/
      handler.go     # Gin handlers (GetAll, GetByID, Add)
      model.go       # Album struct definition
      store.go       # In-memory data store (seed albums)
  go.mod
  go.sum
  README.md

```

Getting started
Prerequisites
Go 1.21 or later installed.[web:71]

Git installed.

Setup
```
git clone https://github.com/Debashich/gin-api.git
cd gin-api

go mod tidy
```
Run the server
From the project root:

```
go run ./cmd/server
```
By default the API listens on http://localhost:8080.[web:1]

API
Endpoints
``GET /albums`` – return all albums.

``GET /albums/:id`` – return a single album by ID.

``POST /albums`` – create a new album.[web:1]

Example requests (cURL)
List all albums
text
curl http://localhost:8080/albums
Get a single album by ID
text
curl http://localhost:8080/albums/1
If the album does not exist, the API returns a 404 with a JSON error.

Create a new album
```
curl -X POST http://localhost:8080/albums \
  -H "Content-Type: application/json" \
  -d '{
    "id": "4",
    "title": "New Album",
    "artist": "Some Artist",
    "price": 25.50
  }'
```
Then verify it was added:

```
curl http://localhost:8080/albums
```
