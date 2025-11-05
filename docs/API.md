# API Documentation

> **Last Updated:** 2025-11-05
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
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com"
}
```

#### `POST /users`

Create a new user.

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com"
}
```

**Response:**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com"
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
  "email": "john.updated@example.com"
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
