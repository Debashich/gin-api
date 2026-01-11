# go-albums-api

A minimal REST API for managing music albums, built with Go and Gin.

## Project Structure

```
go-albums-api/
  cmd/
    server/
      main.go        // entrypoint: starts Gin and registers routes
  internal/
    albums/
      handler.go     // Gin handlers (getAlbums, getAlbumByID, postAlbums)
      model.go       // album struct
      store.go       // in-memory data (later DB)
  go.mod
  README.md
```

## Usage

1. Clone the repo
2. Run `go mod tidy`
3. Start the server:
   ```sh
   go run ./cmd/server
   ```
4. API Endpoints:
   - `GET /albums` — list all albums
   - `GET /albums/:id` — get album by ID
   - `POST /albums` — add a new album

## cURL examples

List all albums:

```sh
curl -X GET http://localhost:8080/albums
```

Get a single album by ID:

```sh
curl -X GET http://localhost:8080/albums/1
```

Create a new album:

```sh
curl -X POST http://localhost:8080/albums \
  -H "Content-Type: application/json" \
  -d '{
    "id": "4",
    "title": "New Album",
    "artist": "Some Artist",
    "price": 25.50
  }'
```

## License

MIT
