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

#### `GET /users`

List all users with count information.

**Response:**
```json
{
  "users": [
    {
      "id": "usr_001",
      "name": "Alice Johnson",
      "email": "alice@example.com",
      "role": "admin",
      "phone_number": "+1-555-0123",
      "avatar": "https://example.com/avatars/alice.jpg",
      "created_at": "2025-11-05T12:00:00Z",
      "updated_at": "2025-11-05T12:00:00Z"
    },
    {
      "id": "usr_002",
      "name": "Bob Smith",
      "email": "bob@example.com",
      "role": "user",
      "phone_number": "",
      "avatar": "",
      "created_at": "2025-11-06T00:00:00Z",
      "updated_at": "2025-11-06T00:00:00Z"
    }
  ],
  "count": 2
}
```

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/v1/users
```

#### `GET /users/:id`

Get a specific user by ID.

**Parameters:**
- `id` (path) - User ID

**Response:**
```json
{
  "id": "usr_001",
  "name": "Alice Johnson",
  "email": "alice@example.com",
  "role": "admin",
  "phone_number": "+1-555-0123",
  "avatar": "https://example.com/avatars/alice.jpg",
  "created_at": "2025-11-05T12:00:00Z",
  "updated_at": "2025-11-05T12:00:00Z"
}
```

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/v1/users/usr_001
```

**Python Example:**
```python
import requests

response = requests.get('http://localhost:8080/api/v1/users/usr_001')
user = response.json()
print(f"User: {user['name']} ({user['email']})")
```

#### `GET /users/:id/profile`

Get enhanced user profile with computed fields.

**Parameters:**
- `id` (path) - User ID

**Response:**
```json
{
  "user": {
    "id": "usr_001",
    "name": "Alice Johnson",
    "email": "alice@example.com",
    "role": "admin",
    "phone_number": "+1-555-0123",
    "avatar": "https://example.com/avatars/alice.jpg",
    "created_at": "2025-11-05T12:00:00Z",
    "updated_at": "2025-11-05T12:00:00Z"
  },
  "profile": {
    "has_avatar": true,
    "has_phone_number": true,
    "is_admin": true,
    "account_age_days": 1
  }
}
```

**Profile Fields:**

| Field | Type | Description |
|-------|------|-------------|
| `has_avatar` | `boolean` | True if user has uploaded an avatar |
| `has_phone_number` | `boolean` | True if user has provided a phone number |
| `is_admin` | `boolean` | True if user has admin role |
| `account_age_days` | `integer` | Number of days since account creation |

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/v1/users/usr_001/profile
```

**Python Example:**
```python
import requests

response = requests.get('http://localhost:8080/api/v1/users/usr_001/profile')
data = response.json()
user = data['user']
profile = data['profile']

print(f"User: {user['name']}")
print(f"Has Avatar: {profile['has_avatar']}")
print(f"Is Admin: {profile['is_admin']}")
print(f"Account Age: {profile['account_age_days']} days")
```

**Go Example:**
```go
type ProfileResponse struct {
    User    models.User `json:"user"`
    Profile struct {
        HasAvatar      bool `json:"has_avatar"`
        HasPhoneNumber bool `json:"has_phone_number"`
        IsAdmin        bool `json:"is_admin"`
        AccountAgeDays int  `json:"account_age_days"`
    } `json:"profile"`
}

resp, err := http.Get("http://localhost:8080/api/v1/users/usr_001/profile")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

var profileResp ProfileResponse
json.NewDecoder(resp.Body).Decode(&profileResp)
fmt.Printf("User: %s, Is Admin: %v\n", profileResp.User.Name, profileResp.Profile.IsAdmin)
```

**Error Responses:**

**404 Not Found:**
```json
{
  "error": "user not found"
}
```

#### `POST /users`

Create a new user.

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "role": "user",
  "phone_number": "+1-555-9876",
  "avatar": "https://example.com/avatars/john.jpg"
}
```

**Required Fields:**
- `name` - User's full name
- `email` - User's email address

**Optional Fields:**
- `role` - User role (defaults to "user" if not provided)
- `phone_number` - User's phone number
- `avatar` - URL to user's avatar image

**Response:**
```json
{
  "id": "usr_003",
  "name": "John Doe",
  "email": "john@example.com",
  "role": "user",
  "phone_number": "+1-555-9876",
  "avatar": "https://example.com/avatars/john.jpg",
  "created_at": "2025-11-06T12:00:00Z",
  "updated_at": "2025-11-06T12:00:00Z"
}
```

**Status Codes:**
- `201 Created` - User created successfully
- `400 Bad Request` - Invalid request body

**cURL Example:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "role": "user",
    "phone_number": "+1-555-9876",
    "avatar": "https://example.com/avatars/john.jpg"
  }'
```

**Python Example:**
```python
import requests

user_data = {
    "name": "John Doe",
    "email": "john@example.com",
    "role": "user",
    "phone_number": "+1-555-9876",
    "avatar": "https://example.com/avatars/john.jpg"
}

response = requests.post(
    'http://localhost:8080/api/v1/users',
    json=user_data
)
new_user = response.json()
print(f"Created user with ID: {new_user['id']}")
```

#### `PUT /users/:id`

Update an existing user.

**Parameters:**
- `id` (path) - User ID

**Request Body:**
```json
{
  "name": "John Updated",
  "email": "john.updated@example.com",
  "role": "admin",
  "phone_number": "+1-555-1111",
  "avatar": "https://example.com/avatars/john-new.jpg"
}
```

**Response:**
```json
{
  "id": "usr_003",
  "name": "John Updated",
  "email": "john.updated@example.com",
  "role": "admin",
  "phone_number": "+1-555-1111",
  "avatar": "https://example.com/avatars/john-new.jpg",
  "created_at": "2025-11-06T12:00:00Z",
  "updated_at": "2025-11-06T14:30:00Z"
}
```

**Status Codes:**
- `200 OK` - User updated successfully
- `400 Bad Request` - Invalid request body
- `404 Not Found` - User not found

**cURL Example:**
```bash
curl -X PUT http://localhost:8080/api/v1/users/usr_003 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Updated",
    "email": "john.updated@example.com",
    "role": "admin",
    "phone_number": "+1-555-1111",
    "avatar": "https://example.com/avatars/john-new.jpg"
  }'
```

#### `DELETE /users/:id`

Delete a user.

**Parameters:**
- `id` (path) - User ID

**Response:**
```json
{
  "message": "User deleted successfully"
}
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
