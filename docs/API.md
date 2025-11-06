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
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  }
]
```

#### `GET /users/:id`

Get a specific user by ID.

**Parameters:**
- `id` (path) - User ID

**Response:**
```json
{
  "id": "usr_123456",
  "name": "John Doe",
  "email": "john@example.com",
  "role": "user",
  "phone_number": "",
  "avatar": "",
  "created_at": "2025-11-06T10:30:00Z",
  "updated_at": "2025-11-06T10:30:00Z"
}
```

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/v1/users/usr_123456 \
  -H "Content-Type: application/json"
```

**Python Example:**
```python
import requests

response = requests.get('http://localhost:8080/api/v1/users/usr_123456')
user = response.json()
print(f"User: {user['name']} ({user['email']})")
```

**Error Responses:**
- `404 Not Found`: User does not exist
```json
{
  "error": "user not found"
}
```

---

#### `GET /users/:id/profile`

Get enhanced user profile with computed fields and additional metadata.

**Added in:** PR #5

**Parameters:**
- `id` (path) - User ID

**Response:**
```json
{
  "user": {
    "id": "usr_123456",
    "name": "Jane Smith",
    "email": "jane@example.com",
    "role": "admin",
    "phone_number": "+1-555-987-6543",
    "avatar": "https://example.com/avatars/jane.jpg",
    "created_at": "2025-11-01T08:00:00Z",
    "updated_at": "2025-11-06T10:30:00Z"
  },
  "profile": {
    "has_avatar": true,
    "has_phone_number": true,
    "is_admin": true,
    "account_age_days": 5
  }
}
```

**Response Fields:**

| Field | Type | Description |
|-------|------|-------------|
| `user` | `object` | Complete user object with all fields |
| `profile.has_avatar` | `boolean` | Whether user has uploaded an avatar image |
| `profile.has_phone_number` | `boolean` | Whether user has provided a phone number |
| `profile.is_admin` | `boolean` | Whether user has admin role |
| `profile.account_age_days` | `integer` | Number of days since account creation |

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/v1/users/usr_123456/profile \
  -H "Content-Type: application/json"
```

**Python Example:**
```python
import requests

response = requests.get('http://localhost:8080/api/v1/users/usr_123456/profile')
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
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type ProfileResponse struct {
    User    User    `json:"user"`
    Profile Profile `json:"profile"`
}

type Profile struct {
    HasAvatar      bool `json:"has_avatar"`
    HasPhoneNumber bool `json:"has_phone_number"`
    IsAdmin        bool `json:"is_admin"`
    AccountAgeDays int  `json:"account_age_days"`
}

func main() {
    resp, _ := http.Get("http://localhost:8080/api/v1/users/usr_123456/profile")
    defer resp.Body.Close()

    var data ProfileResponse
    json.NewDecoder(resp.Body).Decode(&data)

    fmt.Printf("User: %s\n", data.User.Name)
    fmt.Printf("Has Avatar: %v\n", data.Profile.HasAvatar)
}
```

**Error Responses:**
- `404 Not Found`: User does not exist
```json
{
  "error": "user not found"
}
```

**Use Cases:**
- Display rich user profiles in frontend applications
- Check user completion status (avatar, phone number)
- Implement role-based access control checks
- Show account age for gamification features
- Verify user profile completeness before certain actions

---

#### `POST /users`

Create a new user.

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "role": "user",
  "phone_number": "+1-555-123-4567",
  "avatar": "https://example.com/avatars/john.jpg"
}
```

**Request Body Fields:**

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | `string` | Yes | Full name of the user |
| `email` | `string` | Yes | Email address (must be unique) |
| `role` | `string` | Yes | User role ("user" or "admin") |
| `phone_number` | `string` | No | Phone number in international format |
| `avatar` | `string` | No | URL to avatar image |

**Response:**
```json
{
  "id": "usr_123456",
  "name": "John Doe",
  "email": "john@example.com",
  "role": "user",
  "phone_number": "+1-555-123-4567",
  "avatar": "https://example.com/avatars/john.jpg",
  "created_at": "2025-11-06T10:30:00Z",
  "updated_at": "2025-11-06T10:30:00Z"
}
```

**cURL Example:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "role": "user",
    "phone_number": "+1-555-123-4567",
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
    "phone_number": "+1-555-123-4567",
    "avatar": "https://example.com/avatars/john.jpg"
}

response = requests.post(
    'http://localhost:8080/api/v1/users',
    json=user_data
)
new_user = response.json()
print(f"Created user: {new_user['id']}")
```

**Error Responses:**
- `400 Bad Request`: Invalid request body or missing required fields
```json
{
  "error": "Invalid request format"
}
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
  "phone_number": "+1-555-999-8888",
  "avatar": "https://example.com/avatars/john-new.jpg"
}
```

**Request Body Fields:**

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | `string` | No | Updated full name |
| `email` | `string` | No | Updated email address |
| `role` | `string` | No | Updated user role |
| `phone_number` | `string` | No | Updated phone number |
| `avatar` | `string` | No | Updated avatar URL |

**Response:**
```json
{
  "id": "usr_123456",
  "name": "John Updated",
  "email": "john.updated@example.com",
  "role": "admin",
  "phone_number": "+1-555-999-8888",
  "avatar": "https://example.com/avatars/john-new.jpg",
  "created_at": "2025-11-06T10:30:00Z",
  "updated_at": "2025-11-06T15:45:00Z"
}
```

**cURL Example:**
```bash
curl -X PUT http://localhost:8080/api/v1/users/usr_123456 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Updated",
    "email": "john.updated@example.com",
    "phone_number": "+1-555-999-8888"
  }'
```

**Error Responses:**
- `400 Bad Request`: Invalid request body
- `404 Not Found`: User does not exist
```json
{
  "error": "user not found"
}
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
