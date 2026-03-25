package handler

import (
	"buku-kotak-api/handlers"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func init() {
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
	api := app.Group("/api")
	// routes
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173", os.Getenv("BUKU_KOTAK_URL")},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
	}))
	api.GET("/history", handlers.PaperHistory)
	api.PUT("/history", handlers.PaperDeleteHistory)
	api.POST("/upload", handlers.PaperUpload)

	if err := router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
