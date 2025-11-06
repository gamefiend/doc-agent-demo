# API Documentation

> **Last Updated:** 2025-11-06
>
> This document describes all API endpoints for the doc-agent-demo Go API.

## Overview

The API provides endpoints for managing users and products, with health check functionality.

**Base URL:** `http://localhost:8080/api/v1`

## Endpoints

### Health Check

#### `GET /health`

Returns the health status of the API.

**Response:**
```json
{
  "status": "healthy",
  "timestamp": "2025-11-05T12:00:00Z"
}
```

---

### Users

User management endpoints provide full CRUD operations for user entities. All user responses include the `last_login` field for tracking user engagement.

#### `GET /users`

Retrieve a list of all users in the system.

**Response: `200 OK`**
```json
[
  {
    "id": "usr_001",
    "name": "Alice Johnson",
    "email": "alice@example.com",
    "role": "admin",
    "phone_number": "+1-555-0100",
    "avatar": "https://example.com/avatars/alice.jpg",
    "last_login": "2025-11-06T14:30:00Z",
    "created_at": "2025-11-05T08:00:00Z",
    "updated_at": "2025-11-06T14:30:00Z"
  },
  {
    "id": "usr_002",
    "name": "Bob Smith",
    "email": "bob@example.com",
    "role": "user",
    "phone_number": "+1-555-0101",
    "avatar": "https://example.com/avatars/bob.jpg",
    "last_login": null,
    "created_at": "2025-11-05T12:00:00Z",
    "updated_at": "2025-11-05T12:00:00Z"
  }
]
```

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json"
```

**Python Example:**
```python
import requests

response = requests.get('http://localhost:8080/api/v1/users')
users = response.json()

for user in users:
    last_login = user.get('last_login', 'Never')
    print(f"{user['name']} - Last login: {last_login}")
```

---

#### `GET /users/:id`

Get a specific user by their unique identifier.

**Parameters:**
- `id` (path, required) - User ID (e.g., "usr_001")

**Response: `200 OK`**
```json
{
  "id": "usr_001",
  "name": "Alice Johnson",
  "email": "alice@example.com",
  "role": "admin",
  "phone_number": "+1-555-0100",
  "avatar": "https://example.com/avatars/alice.jpg",
  "last_login": "2025-11-06T14:30:00Z",
  "created_at": "2025-11-05T08:00:00Z",
  "updated_at": "2025-11-06T14:30:00Z"
}
```

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/v1/users/usr_001 \
  -H "Content-Type: application/json"
```

**Go Example:**
```go
import (
    "encoding/json"
    "net/http"
)

resp, err := http.Get("http://localhost:8080/api/v1/users/usr_001")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

var user models.User
json.NewDecoder(resp.Body).Decode(&user)

if user.LastLogin != nil {
    fmt.Printf("User last logged in: %s\n", user.LastLogin.Format(time.RFC3339))
} else {
    fmt.Println("User has never logged in")
}
```

---

#### `POST /users`

Create a new user. The `last_login` field will be `null` by default and should be updated by the authentication system upon successful login.

**Request Body:**
```json
{
  "id": "usr_003",
  "name": "Charlie Brown",
  "email": "charlie@example.com",
  "role": "user",
  "phone_number": "+1-555-0102",
  "avatar": "https://example.com/avatars/charlie.jpg"
}
```

**Response: `201 Created`**
```json
{
  "id": "usr_003",
  "name": "Charlie Brown",
  "email": "charlie@example.com",
  "role": "user",
  "phone_number": "+1-555-0102",
  "avatar": "https://example.com/avatars/charlie.jpg",
  "last_login": null,
  "created_at": "2025-11-06T15:00:00Z",
  "updated_at": "2025-11-06T15:00:00Z"
}
```

**cURL Example:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "id": "usr_003",
    "name": "Charlie Brown",
    "email": "charlie@example.com",
    "role": "user",
    "phone_number": "+1-555-0102",
    "avatar": "https://example.com/avatars/charlie.jpg"
  }'
```

**Python Example:**
```python
import requests

new_user = {
    "id": "usr_003",
    "name": "Charlie Brown",
    "email": "charlie@example.com",
    "role": "user",
    "phone_number": "+1-555-0102",
    "avatar": "https://example.com/avatars/charlie.jpg"
}

response = requests.post(
    'http://localhost:8080/api/v1/users',
    json=new_user
)

if response.status_code == 201:
    user = response.json()
    print(f"Created user: {user['id']}")
```

---

#### `PUT /users/:id`

Update an existing user. The `last_login` field is typically managed by the authentication system and should not be manually updated via this endpoint.

**Parameters:**
- `id` (path, required) - User ID (e.g., "usr_001")

**Request Body:**
```json
{
  "name": "Alice Johnson-Smith",
  "email": "alice.smith@example.com",
  "role": "admin",
  "phone_number": "+1-555-0199",
  "avatar": "https://example.com/avatars/alice-new.jpg"
}
```

**Response: `200 OK`**
```json
{
  "id": "usr_001",
  "name": "Alice Johnson-Smith",
  "email": "alice.smith@example.com",
  "role": "admin",
  "phone_number": "+1-555-0199",
  "avatar": "https://example.com/avatars/alice-new.jpg",
  "last_login": "2025-11-06T14:30:00Z",
  "created_at": "2025-11-05T08:00:00Z",
  "updated_at": "2025-11-06T15:30:00Z"
}
```

**cURL Example:**
```bash
curl -X PUT http://localhost:8080/api/v1/users/usr_001 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice Johnson-Smith",
    "email": "alice.smith@example.com",
    "role": "admin",
    "phone_number": "+1-555-0199",
    "avatar": "https://example.com/avatars/alice-new.jpg"
  }'
```

---

#### `DELETE /users/:id`

Delete a user from the system. This will permanently remove the user and their associated data.

**Parameters:**
- `id` (path, required) - User ID (e.g., "usr_001")

**Response: `200 OK`**
```json
{
  "message": "User deleted successfully"
}
```

**cURL Example:**
```bash
curl -X DELETE http://localhost:8080/api/v1/users/usr_001 \
  -H "Content-Type: application/json"
```

---

### Products

#### `GET /products`

List all products.

**Response:**
```json
[
  {
    "id": 1,
    "name": "Product Name",
    "price": 99.99
  }
]
```

#### `GET /products/:id`

Get a specific product by ID.

**Parameters:**
- `id` (path) - Product ID

**Response:**
```json
{
  "id": 1,
  "name": "Product Name",
  "price": 99.99
}
```

#### `POST /products`

Create a new product.

**Request Body:**
```json
{
  "name": "Product Name",
  "price": 99.99
}
```

**Response:**
```json
{
  "id": 1,
  "name": "Product Name",
  "price": 99.99
}
```

---

## Error Responses

All endpoints may return the following error responses:

**400 Bad Request:**
```json
{
  "error": "Invalid request format"
}
```

**404 Not Found:**
```json
{
  "error": "Resource not found"
}
```

**500 Internal Server Error:**
```json
{
  "error": "Internal server error"
}
```

---

> **Note:** This documentation is maintained by the automated documentation bot.
> When code changes are merged, the bot analyzes the changes and updates this file accordingly.
