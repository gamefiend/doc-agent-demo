# API Documentation

> **Last Updated:** 2025-11-06
>
> This document describes all API endpoints for the doc-agent-demo Go API.

## Overview

The API provides endpoints for managing users and products, with comprehensive health check functionality including detailed system information.

**Base URL:** `http://localhost:8080/api/v1`

## Endpoints

### Health Check

#### `GET /health`

Returns the basic health status of the API. This is a lightweight endpoint suitable for load balancer health checks and basic monitoring.

**Response (200 OK):**
```json
{
  "status": "healthy",
  "timestamp": "2025-11-05T12:00:00Z",
  "version": "1.0.0"
}
```

**Example Request (cURL):**
```bash
curl -X GET http://localhost:8080/api/v1/health
```

**Example Request (Go):**
```go
resp, err := http.Get("http://localhost:8080/api/v1/health")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

var health HealthResponse
json.NewDecoder(resp.Body).Decode(&health)
fmt.Printf("Status: %s\n", health.Status)
```

**Example Request (Python):**
```python
import requests

response = requests.get('http://localhost:8080/api/v1/health')
health = response.json()
print(f"Status: {health['status']}")
```

---

#### `GET /health/details`

Returns detailed system health information including runtime metrics, system information, and dependency versions. This endpoint is useful for debugging, monitoring dashboards, and detailed system status checks.

**Response (200 OK):**
```json
{
  "status": "healthy",
  "timestamp": "2025-11-05T12:00:00Z",
  "version": "1.0.0",
  "uptime": "2h15m30.5s",
  "system": {
    "go_version": "go1.23.1",
    "num_cpu": 8,
    "num_goroutine": 15,
    "os": "linux",
    "arch": "amd64"
  },
  "dependencies": {
    "gin": "v1.9.x",
    "runtime": "go1.23.1"
  }
}
```

**Response Fields:**

| Field | Type | Description |
|-------|------|-------------|
| `status` | `string` | Current health status (e.g., "healthy") |
| `timestamp` | `string` | ISO 8601 timestamp of the response |
| `version` | `string` | API version |
| `uptime` | `string` | Duration since application started (e.g., "2h15m30.5s") |
| `system` | `object` | System runtime information |
| `system.go_version` | `string` | Go runtime version |
| `system.num_cpu` | `int` | Number of logical CPUs |
| `system.num_goroutine` | `int` | Current number of goroutines |
| `system.os` | `string` | Operating system (e.g., "linux", "darwin", "windows") |
| `system.arch` | `string` | System architecture (e.g., "amd64", "arm64") |
| `dependencies` | `object` | Map of dependency names to versions |

**Example Request (cURL):**
```bash
curl -X GET http://localhost:8080/api/v1/health/details
```

**Example Request (Go):**
```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type HealthDetailsResponse struct {
    Status       string            `json:"status"`
    Timestamp    string            `json:"timestamp"`
    Version      string            `json:"version"`
    Uptime       string            `json:"uptime"`
    System       SystemInfo        `json:"system"`
    Dependencies map[string]string `json:"dependencies"`
}

type SystemInfo struct {
    GoVersion    string `json:"go_version"`
    NumCPU       int    `json:"num_cpu"`
    NumGoroutine int    `json:"num_goroutine"`
    OS           string `json:"os"`
    Arch         string `json:"arch"`
}

func main() {
    resp, err := http.Get("http://localhost:8080/api/v1/health/details")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    var details HealthDetailsResponse
    if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Status: %s\n", details.Status)
    fmt.Printf("Uptime: %s\n", details.Uptime)
    fmt.Printf("Go Version: %s\n", details.System.GoVersion)
    fmt.Printf("CPUs: %d\n", details.System.NumCPU)
    fmt.Printf("Goroutines: %d\n", details.System.NumGoroutine)
}
```

**Example Request (Python):**
```python
import requests
from datetime import datetime

response = requests.get('http://localhost:8080/api/v1/health/details')
details = response.json()

print(f"Status: {details['status']}")
print(f"Uptime: {details['uptime']}")
print(f"Go Version: {details['system']['go_version']}")
print(f"CPUs: {details['system']['num_cpu']}")
print(f"Goroutines: {details['system']['num_goroutine']}")
print(f"OS: {details['system']['os']} ({details['system']['arch']})")
print(f"\nDependencies:")
for dep, version in details['dependencies'].items():
    print(f"  - {dep}: {version}")
```

**Use Cases:**

- **Monitoring Dashboards**: Display real-time system metrics and uptime
- **Debugging**: Check goroutine counts and resource usage during troubleshooting
- **Capacity Planning**: Monitor CPU count and system architecture
- **Dependency Auditing**: Track deployed dependency versions
- **Health Checks**: More comprehensive health validation beyond basic liveness

**Status Codes:**

| Code | Description |
|------|-------------|
| 200 | System is healthy and operational |

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
