# `pokedex/pokemon.go` — Data model and in-memory store

Overview

This file implements the core data model (`Pokemon`) and an in-memory thread-safe `Store` that provides CRUD operations used by the HTTP handlers.

Types

- `Pokemon`
  - Fields:
    - `ID string` — auto-assigned identifier (simple integer as string).
    - `Name string` — species name, e.g. "Pikachu".
    - `CP int` — Combat Power (integer).
    - `IsShiny bool` — whether the Pokemon is a shiny variant.

- `Store`
  - `pokemon map[string]Pokemon` — holds entries keyed by `ID`.
  - `mu sync.RWMutex` — protects concurrent access.
  - `nextID int` — monotonic counter for generating IDs.

Constructor

- `NewStore() *Store`
  - Returns an initialized `Store` with an empty map and `nextID` set to 1.

Methods (API)

- `CreatePokemon(p Pokemon) Pokemon`
  - Locks the store for writing, assigns `p.ID` using `nextID`, stores the Pokemon and increments `nextID`. Returns the created `Pokemon` (including the assigned ID).

- `GetPokemonByID(id string) (Pokemon, error)`
  - RLock; returns the pokemon if found, otherwise returns an error `"pokemon not found"`.

- `GetAllPokemon() []Pokemon`
  - RLock; returns a slice containing all Pokemon entries in the store.

- `UpdatePokemon(id string, updatedP Pokemon) (Pokemon, error)`
  - Lock; if the ID exists, updates the entry (ensures `ID` field stays the same) and returns the updated value; otherwise returns `"pokemon not found"`.

- `DeletePokemon(id string) error`
  - Lock; deletes entry if it exists; otherwise returns `"pokemon not found"`.

Concurrency and thread-safety

- The `Store` uses `sync.RWMutex` to allow multiple concurrent readers while serializing writers. All public methods appropriately lock/unlock the mutex.

ID generation and limitations

- ID generation is a simple integer counter converted to string. It is fine for local/testing usage but not globally unique in distributed or persistent systems.

Edge cases / Errors

- Methods return `errors.New("pokemon not found")` when the target ID is missing. Handlers map this to HTTP 404.
- There is no validation on the fields (e.g., `Name` non-empty or `CP` non-negative) — consider adding validation where entries are created/updated.

Possible Improvements

- Swap the in-memory store for a persistent database (SQLite/Postgres) with a data access layer.
- Add field validation and structured error types.
- Use UUIDs for IDs if external uniqueness is required.
