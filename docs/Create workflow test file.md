## `Pokemon` Struct and CRUD Operations

This section details the `Pokemon` struct and its associated CRUD (Create, Read, Update, Delete) methods.

### `Pokemon` Struct

The `Pokemon` struct represents a Pokémon with its core attributes.

```go
type Pokemon struct {
	ID      int    `json:"id"`      // Unique identifier for the Pokémon.
	Name    string `json:"name"`    // The name of the Pokémon.
	Type    string `json:"type"`    // The primary type of the Pokémon.
	HP      int    `json:"hp"`      // Hit Points.
	Attack  int    `json:"attack"`  // Attack stat.
	Defense int    `json:"defense"` // Defense stat.
}
```

### CRUD Methods

#### `CreatePokemon`

Creates a new Pokémon entry in the database.

-   **Purpose:** Inserts a new Pokémon record into the database.
-   **Parameters:**
    -   `db *sql.DB`: A database connection.
    -   `pokemon Pokemon`: The `Pokemon` struct containing the Pokémon's data.
-   **Return Values:**
    -   `int`: The ID of the newly created Pokémon, or 0 if an error occurred.
    -   `error`: An error object if the operation fails, or `nil` on success.
-   **Example:**

    ```go
    newPokemon := Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, Attack: 55, Defense: 40}
    pokemonID, err := CreatePokemon(db, newPokemon)
    if err != nil {
    	// Handle error
    }
    fmt.Printf("Created Pokemon with ID: %d\n", pokemonID)
    ```

#### `ReadPokemon`

Retrieves a Pokémon's information from the database by its ID.

-   **Purpose:** Fetches a Pokémon's data based on its ID.
-   **Parameters:**
    -   `db *sql.DB`: A database connection.
    -   `id int`: The ID of the Pokémon to retrieve.
-   **Return Values:**
    -   `Pokemon`: The `Pokemon` struct if found, or an empty `Pokemon` struct if not found.
    -   `error`: An error object if the operation fails, or `nil` on success.
-   **Example:**

    ```go
    pokemon, err := ReadPokemon(db, 1)
    if err != nil {
    	// Handle error
    }
    fmt.Printf("Pokemon Name: %s\n", pokemon.Name)
    ```

#### `UpdatePokemon`

Updates an existing Pokémon's information in the database.

-   **Purpose:** Modifies an existing Pokémon record in the database.
-   **Parameters:**
    -   `db *sql.DB`: A database connection.
    -   `pokemon Pokemon`: The `Pokemon` struct containing the updated Pokémon data. The `ID` field is used to identify the record to update.
-   **Return Values:**
    -   `error`: An error object if the operation fails, or `nil` on success.
-   **Example:**

    ```go
    updatedPokemon := Pokemon{ID: 1, Name: "Raichu", Type: "Electric", HP: 60, Attack: 90, Defense: 55}
    err := UpdatePokemon(db, updatedPokemon)
    if err != nil {
    	// Handle error
    }
    fmt.Println("Pokemon updated successfully")
    ```

#### `DeletePokemon`

Deletes a Pokémon from the database by its ID.

-   **Purpose:** Removes a Pokémon record from the database.
-   **Parameters:**
    -   `db *sql.DB`: A database connection.
    -   `id int`: The ID of the Pokémon to delete.
-   **Return Values:**
    -   `error`: An error object if the operation fails, or `nil` on success.
-   **Example:**

    ```go
    err := DeletePokemon(db, 1)
    if err != nil {
    	// Handle error
    }
    fmt.Println("Pokemon deleted successfully")
    ```