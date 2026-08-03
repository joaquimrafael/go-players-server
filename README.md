# Go Players Server

A small Texas Hold'em application written in Go. It includes a command-line game and an HTTP server that stores player wins in a JSON file.

## Requirements

- Go 1.26.4 or later

## Run the project

Start the web server on `http://localhost:5500`:

```sh
go run ./cmd/webserver
```

Or play from the command line:

```sh
go run ./cmd/cli
```

## API

- `GET /league` lists players and scores.
- `GET /players/{name}` returns a player's score.
- `POST /players/{name}` records a win.
- `GET /game` opens the game page.

## Tests

```sh
go test ./...
```
