package main

import (
	"buku-kotak-api/api"
	"log"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
)

func main() {
	host, _ := os.Hostname()
	// set env for local
	// The `(?i)` flag enables case-insensitive matching
	pattern := regexp.MustCompile(`(?i)desktop`)
	if pattern.MatchString(host) {
		os.Setenv("CLOUDINARY_API_NAME", "")
		os.Setenv("CLOUDINARY_API_KEY", "")
		os.Setenv("CLOUDINARY_API_SECRET", "")
		os.Setenv("CLOUDINARY_URL", "")
		os.Setenv("UPSTASH_REDIS_REST_TCP", "")
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// routes
	router.GET("/api/history", api.PaperHistory)
	router.POST("/api/upload", api.PaperUpload)

	if err := router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
