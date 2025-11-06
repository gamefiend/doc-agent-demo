# `main.go` — Entrypoint for the API Server

> **Last Updated:** 2025-11-06

## Purpose

`main.go` wires the application together and starts the HTTP server. It performs the following responsibilities:

- Initialize the Gin router with default middleware
- Register custom middleware (RequestLoggerMiddleware)
- Set up API routes under `/api/v1` group
- Initialize sample data for users and products
- Start the HTTP server on port `:8080`

## Key Components

### Router Initialization

```go
r := gin.Default()
```

Creates a Gin router with default middleware (Logger and Recovery).

### Custom Middleware

#### RequestLoggerMiddleware

**Purpose:** Logs all API requests with timing information for monitoring and debugging.

**Implementation:**
```go
func RequestLoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path
        method := c.Request.Method

        c.Next() // Process request

        duration := time.Since(start)
        status := c.Writer.Status()
        log.Printf("[%s] %s - Status: %d - Duration: %v", method, path, status, duration)
    }
}
```

**What it logs:**
- HTTP method (GET, POST, PUT, DELETE)
- Request path
- Response status code
- Request duration

**Example output:**
```
[GET] /api/v1/users/usr_001 - Status: 200 - Duration: 2.345ms
[POST] /api/v1/users - Status: 201 - Duration: 5.123ms
```

### API Routes

All routes are organized under the `/api/v1` group:

**Health Check:**
- `GET /api/v1/health` - Health check endpoint

**User Endpoints:**
- `GET /api/v1/users` - List all users
- `GET /api/v1/users/:id` - Get user by ID
- `GET /api/v1/users/:id/profile` - Get enhanced user profile (NEW in PR #5)
- `POST /api/v1/users` - Create new user
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

**Product Endpoints:**
- `GET /api/v1/products` - List all products
- `GET /api/v1/products/:id` - Get product by ID
- `POST /api/v1/products` - Create new product

## Server Initialization

The server initialization process:

1. Creates Gin router with `gin.Default()`
2. Registers `RequestLoggerMiddleware()` for request logging
3. Creates `/api/v1` route group
4. Registers all endpoint handlers
5. Calls `models.InitSampleData()` to populate demo data
6. Starts HTTP server on port 8080

```go
func main() {
    r := gin.Default()
    r.Use(RequestLoggerMiddleware())

    v1 := r.Group("/api/v1")
    {
        v1.GET("/health", handlers.HealthCheck)
        // ... register all routes ...
    }

    models.InitSampleData()

    log.Println("Starting server on :8080")
    if err := r.Run(":8080"); err != nil && err != http.ErrServerClosed {
        log.Fatalf("Failed to start server: %v", err)
    }
}
```

## Run / Example

From the project root directory:

```bash
# Run the API server
go run cmd/api/main.go

# Server will be available at http://localhost:8080
```

**Testing the server:**

```bash
# Health check
curl http://localhost:8080/api/v1/health

# List users
curl http://localhost:8080/api/v1/users

# Get user profile (NEW in PR #5)
curl http://localhost:8080/api/v1/users/usr_001/profile
```

## Success Criteria / Exit Behavior

- The program logs "Starting server on :8080" and blocks running the HTTP server
- Each request is logged with method, path, status, and duration
- If the server fails to start, `log.Fatal` prints the error and exits with non-zero status

## Sample Data

The application initializes with sample data for demonstration:

**Users:**
- `usr_001` - Alice Johnson (admin)
- `usr_002` - Bob Smith (user)

**Products:**
- `prd_001` - Laptop ($999.99)
- `prd_002` - Mouse ($29.99)

## Notes / Next Steps

**Current Features:**
- Custom request logging middleware for monitoring
- Enhanced user profile endpoint with computed fields
- User model supports phone numbers and avatars
- Thread-safe in-memory data storage

**Potential Improvements:**
- Add graceful shutdown handling (context cancellation and http.Server.Shutdown)
- Implement persistent storage (database connection)
- Add authentication/authorization middleware
- Implement rate limiting
- Add CORS support for frontend integration
- Add request validation middleware
- Implement structured logging (e.g., logrus, zap)
