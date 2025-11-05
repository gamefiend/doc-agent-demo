package pokedex

import (
	"errors"
	"strconv"
	"sync"
)

// Pokemon represents a single entry in the personal Pokédex.
type Pokemon struct {
	ID     string `json:"id"`     // Unique identifier for the captured Pokemon
	Name   string `json:"name"`   // The species name (e.g., "Pikachu")
	CP     int    `json:"cp"`     // The Combat Power of the captured Pokemon
	IsShiny bool  `json:"isShiny"` // Whether the Pokemon is a shiny variant
}

// Store holds the in-memory state of the Pokédex.
type Store struct {
	pokemon map[string]Pokemon
	mu      sync.RWMutex // Mutex to safely handle concurrent access
	nextID  int          // Counter for simple ID generation
}

// NewStore initializes a new in-memory Pokédex store.
func NewStore() *Store {
	return &Store{
		pokemon: make(map[string]Pokemon),
		nextID:  1,
	}
}

// --- CRUD Operations (Core Logic) ---

// CreatePokemon adds a new Pokemon to the store and assigns an ID.
func (s *Store) CreatePokemon(p Pokemon) Pokemon {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Simple ID generation
	p.ID = strconv.Itoa(s.nextID)
	s.pokemon[p.ID] = p
	s.nextID++
	return p
}

// GetPokemonByID retrieves a Pokemon by its ID.
func (s *Store) GetPokemonByID(id string) (Pokemon, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	p, exists := s.pokemon[id]
	if !exists {
		return Pokemon{}, errors.New("pokemon not found")
	}
	return p, nil
}

// GetAllPokemon retrieves all Pokemon in the store.
func (s *Store) GetAllPokemon() []Pokemon {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Convert map to slice for response
	list := make([]Pokemon, 0, len(s.pokemon))
	for _, p := range s.pokemon {
		list = append(list, p)
	}
	return list
}

// UpdatePokemon updates an existing Pokemon entry.
func (s *Store) UpdatePokemon(id string, updatedP Pokemon) (Pokemon, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.pokemon[id]; !exists {
		return Pokemon{}, errors.New("pokemon not found")
	}

	updatedP.ID = id // Ensure the ID remains the same
	s.pokemon[id] = updatedP
	return updatedP, nil
}

// DeletePokemon removes a Pokemon from the store by its ID.
func (s *Store) DeletePokemon(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.pokemon[id]; !exists {
		return errors.New("pokemon not found")
	}

	delete(s.pokemon, id)
	return nil
}