package main

import (
	"buku-kotak-api/api"
	"log"
	"os"
	"regexp"

	"github.com/gin-contrib/cors"
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
		os.Setenv("BUKU_KOTAK_URL", "")
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// routes
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173", os.Getenv("BUKU_KOTAK_URL")},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
	}))
	router.GET("/api/history", api.PaperHistory)
	router.PUT("/api/history", api.PaperDeleteHistory)
	router.POST("/api/upload", api.PaperUpload)

	if err := router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
