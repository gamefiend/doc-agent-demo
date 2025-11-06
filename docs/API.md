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

List all users.

**Response:**
```json
{
  "users": [
    {
      "id": "usr_001",
      "name": "Alice Johnson",
      "email": "alice@example.com",
      "role": "admin",
      "phone_number": "+1-555-0100",
      "avatar": "https://example.com/avatars/alice.jpg",
      "created_at": "2025-11-05T10:00:00Z",
      "updated_at": "2025-11-05T10:00:00Z"
    },
    {
      "id": "usr_002",
      "name": "Bob Smith",
      "email": "bob@example.com",
      "role": "user",
      "phone_number": "",
      "avatar": "",
      "created_at": "2025-11-05T22:00:00Z",
      "updated_at": "2025-11-05T22:00:00Z"
    }
  ],
  "count": 2
}
```

#### `GET /users/:id`

Get a specific user by ID.

**Parameters:**
- `id` (path) - User ID (string format, e.g., "usr_001")

**Response:**
```json
{
  "id": "usr_001",
  "name": "Alice Johnson",
  "email": "alice@example.com",
  "role": "admin",
  "phone_number": "+1-555-0100",
  "avatar": "https://example.com/avatars/alice.jpg",
  "created_at": "2025-11-05T10:00:00Z",
  "updated_at": "2025-11-05T10:00:00Z"
}
```

**Error Responses:**

**404 Not Found:**
```json
{
  "error": "user not found"
}
```

#### `GET /users/:id/profile`

Get enhanced user profile with computed fields and additional metadata.

**Parameters:**
- `id` (path) - User ID (string format, e.g., "usr_001")

**Response:**
```json
{
  "user": {
    "id": "usr_001",
    "name": "Alice Johnson",
    "email": "alice@example.com",
    "role": "admin",
    "phone_number": "+1-555-0100",
    "avatar": "https://example.com/avatars/alice.jpg",
    "created_at": "2025-11-05T10:00:00Z",
    "updated_at": "2025-11-05T10:00:00Z"
  },
  "profile": {
    "has_avatar": true,
    "has_phone_number": true,
    "is_admin": true,
    "account_age_days": 1
  }
}
```

**Response Fields:**

| Field | Type | Description |
|-------|------|-------------|
| `user` | `User` | Complete user object with all fields |
| `profile.has_avatar` | `boolean` | Whether the user has an avatar URL set |
| `profile.has_phone_number` | `boolean` | Whether the user has a phone number set |
| `profile.is_admin` | `boolean` | Whether the user has admin role |
| `profile.account_age_days` | `integer` | Number of days since account creation |

**Example with cURL:**
```bash
curl -X GET http://localhost:8080/api/v1/users/usr_001/profile
```

**Example with Python:**
```python
import requests

response = requests.get('http://localhost:8080/api/v1/users/usr_001/profile')
profile_data = response.json()

user = profile_data['user']
profile = profile_data['profile']

print(f"User: {user['name']}")
print(f"Has avatar: {profile['has_avatar']}")
print(f"Is admin: {profile['is_admin']}")
print(f"Account age: {profile['account_age_days']} days")
```

**Example with Go:**
```go
import (
    "encoding/json"
    "fmt"
    "net/http"
)

type ProfileResponse struct {
    User    User   `json:"user"`
    Profile struct {
        HasAvatar      bool `json:"has_avatar"`
        HasPhoneNumber bool `json:"has_phone_number"`
        IsAdmin        bool `json:"is_admin"`
        AccountAgeDays int  `json:"account_age_days"`
    } `json:"profile"`
}

func getUserProfile(userID string) (*ProfileResponse, error) {
    resp, err := http.Get(fmt.Sprintf("http://localhost:8080/api/v1/users/%s/profile", userID))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var profileResp ProfileResponse
    if err := json.NewDecoder(resp.Body).Decode(&profileResp); err != nil {
        return nil, err
    }

    return &profileResp, nil
}
```

**Error Responses:**

**404 Not Found:**
```json
{
  "error": "user not found"
}
```

**Use Cases:**
- Displaying complete user profile pages
- Checking if user has completed profile setup (avatar, phone)
- Conditional UI rendering based on admin status
- Showing account age/tenure information

**Handler Location:** `internal/handlers/user.go:88`

---

#### `POST /users`

Create a new user.

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "role": "user",
  "phone_number": "+1-555-0150",
  "avatar": "https://example.com/avatars/john.jpg"
}
```

**Required Fields:**
- `name` - User's full name
- `email` - User's email address
- `role` - User's role ("user" or "admin")

**Optional Fields:**
- `phone_number` - User's contact phone number
- `avatar` - URL to user's avatar image

**Response:**
```json
{
  "id": "usr_003",
  "name": "John Doe",
  "email": "john@example.com",
  "role": "user",
  "phone_number": "+1-555-0150",
  "avatar": "https://example.com/avatars/john.jpg",
  "created_at": "2025-11-06T10:30:00Z",
  "updated_at": "2025-11-06T10:30:00Z"
}
```

**Status Code:** `201 Created`

**Error Responses:**

**400 Bad Request:**
```json
{
  "error": "invalid request body"
}
```

#### `PUT /users/:id`

Update an existing user.

**Parameters:**
- `id` (path) - User ID (string format, e.g., "usr_001")

**Request Body:**
```json
{
  "name": "John Updated",
  "email": "john.updated@example.com",
  "role": "admin",
  "phone_number": "+1-555-9999",
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
  "phone_number": "+1-555-9999",
  "avatar": "https://example.com/avatars/john-new.jpg",
  "created_at": "2025-11-06T10:30:00Z",
  "updated_at": "2025-11-06T12:00:00Z"
}
```

**Status Code:** `200 OK`

**Error Responses:**

**400 Bad Request:**
```json
{
  "error": "invalid request body"
}
```

**404 Not Found:**
```json
{
  "error": "user not found"
}
```

#### `DELETE /users/:id`

Delete a user.

**Parameters:**
- `id` (path) - User ID (string format, e.g., "usr_001")

**Response:**
```json
{
  "message": "user deleted successfully"
}
```

**Status Code:** `200 OK`

**Error Responses:**

**404 Not Found:**
```json
{
  "error": "user not found"
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
