# `main.go` — Entrypoint for the Pokédex API

Purpose

`main.go` wires the application together and starts the HTTP server. It performs three responsibilities:

- Initialize the in-memory `Store` (business logic).
- Create a `HandlerContext` that injects the store into HTTP handlers.
- Build the router and start the HTTP server on `:8080`.

Key behavior

- Uses the local package `pokedex` to obtain `NewStore()`, `HandlerContext`, and `NewRouter()`.
- Starts the server via `http.ListenAndServe(":8080", router)`; logs fatal on error.

Run / Example

From the `pokedex-api` directory:

```bash
# Run the API server
go run main.go

# Server will be available at http://localhost:8080
```

Success criteria / Exit behavior

- The program logs a start message and then blocks running the HTTP server.
- If the server fails to start, `log.Fatal` will print the error and exit.

Notes / Next steps

- For local development you can add a simple `Makefile` or a `go run ./...` script.
- Consider adding graceful shutdown handling (context cancellation and http.Server) for production-ready behavior.
