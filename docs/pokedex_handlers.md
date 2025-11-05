# `pokedex/handlers.go` — HTTP handlers and router

Overview

This file exposes the HTTP layer for the Pokédex API. It provides request/response utilities, an injected `HandlerContext` that contains the `Store`, concrete HTTP handlers for CRUD operations on `/pokemon`, and a `NewRouter` function that wires endpoints to handlers.

Key types

- `HandlerContext`
  - Fields:
    - `Store *Store` — dependency-injected reference to the in-memory store.

Utility functions

- `respondJSON(w http.ResponseWriter, status int, payload interface{})`
  - Serializes `payload` to JSON and writes it with the given HTTP status.
  - If serialization fails, writes HTTP 500 and the serialization error message.

- `respondError(w http.ResponseWriter, status int, message string)`
  - Helper which sends a JSON object `{ "error": "message" }` with the given status.

HTTP Handlers (endpoints)

All handlers decode/encode JSON and map store errors to appropriate HTTP statuses.

- `CreatePokemonHandler` — POST /pokemon
  - Request body: JSON object matching `Pokemon` (excluding ID).
  - Behavior: decodes payload, calls `Store.CreatePokemon`, returns `201 Created` and the created resource (with assigned `id`).
  - Error cases: invalid JSON -> `400 Bad Request`.

- `GetPokemonHandler` — GET /pokemon/{id}
  - Path parameter: `id`.
  - Behavior: fetches by ID via `Store.GetPokemonByID`; returns `200 OK` with the Pokemon JSON.
  - Error cases: not found -> `404 Not Found`.

- `GetAllPokemonHandler` — GET /pokemon
  - Behavior: returns all stored Pokemon as a JSON array with `200 OK`.

- `UpdatePokemonHandler` — PUT /pokemon/{id}
  - Path parameter: `id`.
  - Request body: JSON object representing updated `Pokemon` fields (ID will be set by the handler/store).
  - Behavior: attempts `Store.UpdatePokemon`; on success returns `200 OK` with updated resource.
  - Error cases: invalid JSON -> `400 Bad Request`; not found -> `404 Not Found`.

- `DeletePokemonHandler` — DELETE /pokemon/{id}
  - Path parameter: `id`.
  - Behavior: attempts `Store.DeletePokemon`; on success returns `200 OK` and a JSON success message.
  - Error cases: not found -> `404 Not Found`.

Router

- `NewRouter(hc *HandlerContext) *mux.Router`
  - Registers all routes using `github.com/gorilla/mux`:
    - `GET /pokemon` -> `GetAllPokemonHandler`
    - `POST /pokemon` -> `CreatePokemonHandler`
    - `GET /pokemon/{id}` -> `GetPokemonHandler`
    - `PUT /pokemon/{id}` -> `UpdatePokemonHandler`
    - `DELETE /pokemon/{id}` -> `DeletePokemonHandler`

Request/Response JSON shapes

- Pokemon (example):

```json
{
  "id": "1",
  "name": "Pikachu",
  "cp": 420,
  "isShiny": false
}
```

- Error response example:

```json
{ "error": "Pokemon not found: 1" }
```

Examples (curl)

- Create:

```bash
curl -X POST http://localhost:8080/pokemon \
  -H "Content-Type: application/json" \
  -d '{"name":"Pikachu","cp":420,"isShiny":false}'
```

- Get all:

```bash
curl http://localhost:8080/pokemon
```

- Get by id:

```bash
curl http://localhost:8080/pokemon/1
```

Notes and Improvements

- Handlers trust the JSON structure; adding validation or using a request DTO would help catch malformed/invalid fields.
- `respondJSON` returns raw serialization errors to the client; consider logging the error and returning a more generic error message for security.
- Consider adding middleware for logging, request timeouts, and CORS as needed.
