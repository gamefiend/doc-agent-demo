# Documentation — Pokédex API

This folder contains generated documentation for the mini Pokédex API project. Files:

- `main.md` — entrypoint explanation and how to run the server.
- `pokedex_store.md` — data model and in-memory store (`pokedex/pokemon.go`).
- `pokedex_handlers.md` — HTTP handlers, routes, and request/response shapes (`pokedex/handlers.go`).

Quick start

1. From the `pokedex-api` folder, run the server:

```bash
# From the project root
cd pokedex-api
go run main.go
```

2. The server listens on `http://localhost:8080` with CRUD endpoints under `/pokemon`.

Notes

- The implementation uses an in-memory store; data is not persisted between runs.
- The router uses `github.com/gorilla/mux`.
