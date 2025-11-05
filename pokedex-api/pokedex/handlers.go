package pokedex

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// HandlerContext holds the application's dependencies (the Store).
// This uses Dependency Injection to provide the store to the handlers.
type HandlerContext struct {
	Store *Store
}

// respondJSON is a utility function to send a JSON response.
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// respondError is a utility function to send an error response.
func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

// --- HTTP Handlers ---

// CreatePokemonHandler handles POST /pokemon
func (hc *HandlerContext) CreatePokemonHandler(w http.ResponseWriter, r *http.Request) {
	var p Pokemon
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	createdP := hc.Store.CreatePokemon(p)
	respondJSON(w, http.StatusCreated, createdP)
}

// GetPokemonHandler handles GET /pokemon/{id}
func (hc *HandlerContext) GetPokemonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	p, err := hc.Store.GetPokemonByID(id)
	if err != nil {
		respondError(w, http.StatusNotFound, "Pokemon not found: "+id)
		return
	}
	respondJSON(w, http.StatusOK, p)
}

// GetAllPokemonHandler handles GET /pokemon
func (hc *HandlerContext) GetAllPokemonHandler(w http.ResponseWriter, r *http.Request) {
	pokemonList := hc.Store.GetAllPokemon()
	respondJSON(w, http.StatusOK, pokemonList)
}

// UpdatePokemonHandler handles PUT /pokemon/{id}
func (hc *HandlerContext) UpdatePokemonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p Pokemon
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	updatedP, err := hc.Store.UpdatePokemon(id, p)
	if err != nil {
		respondError(w, http.StatusNotFound, "Pokemon not found: "+id)
		return
	}
	respondJSON(w, http.StatusOK, updatedP)
}

// DeletePokemonHandler handles DELETE /pokemon/{id}
func (hc *HandlerContext) DeletePokemonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := hc.Store.DeletePokemon(id)
	if err != nil {
		respondError(w, http.StatusNotFound, "Pokemon not found: "+id)
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"result": "success", "id": id})
}

// NewRouter creates and configures the HTTP router.
func NewRouter(hc *HandlerContext) *mux.Router {
	r := mux.NewRouter()

	// CRUD Routes for /pokemon
	r.HandleFunc("/pokemon", hc.GetAllPokemonHandler).Methods("GET")
	r.HandleFunc("/pokemon", hc.CreatePokemonHandler).Methods("POST")
	r.HandleFunc("/pokemon/{id}", hc.GetPokemonHandler).Methods("GET")
	r.HandleFunc("/pokemon/{id}", hc.UpdatePokemonHandler).Methods("PUT")
	r.HandleFunc("/pokemon/{id}", hc.DeletePokemonHandler).Methods("DELETE")

	return r
}