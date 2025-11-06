package handlers

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

type HealthDetailsResponse struct {
	Status       string            `json:"status"`
	Timestamp    time.Time         `json:"timestamp"`
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

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Version:   "1.0.0",
	})
}

// HealthDetails returns detailed system health information
func HealthDetails(c *gin.Context) {
	uptime := time.Since(startTime)

	c.JSON(http.StatusOK, HealthDetailsResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Version:   "1.0.0",
		Uptime:    uptime.String(),
		System: SystemInfo{
			GoVersion:    runtime.Version(),
			NumCPU:       runtime.NumCPU(),
			NumGoroutine: runtime.NumGoroutine(),
			OS:           runtime.GOOS,
			Arch:         runtime.GOARCH,
		},
		Dependencies: map[string]string{
			"gin":     "v1.9.x",
			"runtime": "go" + runtime.Version(),
		},
	})
}
