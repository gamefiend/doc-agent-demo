package models

import (
	"sync"
	"time"
)

type User struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	PhoneNumber string    `json:"phone_number"` // NEW: User's phone number
	Avatar      string    `json:"avatar"`       // NEW: URL to user's avatar image
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// In-memory storage (for demo purposes)
var (
	users    = make(map[string]*User)
	products = make(map[string]*Product)
	mu       sync.RWMutex
)

func InitSampleData() {
	mu.Lock()
	defer mu.Unlock()

	// Sample users
	users["usr_001"] = &User{
		ID:        "usr_001",
		Name:      "Alice Johnson",
		Email:     "alice@example.com",
		Role:      "admin",
		CreatedAt: time.Now().Add(-24 * time.Hour),
		UpdatedAt: time.Now().Add(-24 * time.Hour),
	}

	users["usr_002"] = &User{
		ID:        "usr_002",
		Name:      "Bob Smith",
		Email:     "bob@example.com",
		Role:      "user",
		CreatedAt: time.Now().Add(-12 * time.Hour),
		UpdatedAt: time.Now().Add(-12 * time.Hour),
	}

	// Sample products
	products["prd_001"] = &Product{
		ID:          "prd_001",
		Name:        "Laptop",
		Description: "High-performance laptop",
		Price:       999.99,
		Stock:       10,
		CreatedAt:   time.Now().Add(-48 * time.Hour),
		UpdatedAt:   time.Now().Add(-48 * time.Hour),
	}

	products["prd_002"] = &Product{
		ID:          "prd_002",
		Name:        "Mouse",
		Description: "Wireless mouse",
		Price:       29.99,
		Stock:       50,
		CreatedAt:   time.Now().Add(-36 * time.Hour),
		UpdatedAt:   time.Now().Add(-36 * time.Hour),
	}
}

func GetAllUsers() []*User {
	mu.RLock()
	defer mu.RUnlock()

	result := make([]*User, 0, len(users))
	for _, user := range users {
		result = append(result, user)
	}
	return result
}

func GetUserByID(id string) (*User, bool) {
	mu.RLock()
	defer mu.RUnlock()

	user, exists := users[id]
	return user, exists
}

func CreateUser(user *User) {
	mu.Lock()
	defer mu.Unlock()

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	users[user.ID] = user
}

func UpdateUser(id string, user *User) bool {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[id]; !exists {
		return false
	}

	user.ID = id
	user.UpdatedAt = time.Now()
	users[id] = user
	return true
}

func DeleteUser(id string) bool {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[id]; !exists {
		return false
	}

	delete(users, id)
	return true
}

func GetAllProducts() []*Product {
	mu.RLock()
	defer mu.RUnlock()

	result := make([]*Product, 0, len(products))
	for _, product := range products {
		result = append(result, product)
	}
	return result
}

func GetProductByID(id string) (*Product, bool) {
	mu.RLock()
	defer mu.RUnlock()

	product, exists := products[id]
	return product, exists
}

func CreateProduct(product *Product) {
	mu.Lock()
	defer mu.Unlock()

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	products[product.ID] = product
}
