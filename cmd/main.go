package main

import (
	"image-service/internal/api"
	"image-service/internal/app"
	"image-service/internal/db"
	"image-service/internal/service"
	"image-service/internal/storage"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/viebiz/lit/env"
)

func main() {
	gin.ForceConsoleColor()
	r := gin.Default()
	config, err := env.ReadAppConfig[app.Config]()
	if err != nil {
		log.Fatal("failed to read config:", err)
	}
	// DB
	database, err := db.Connect(config.PG.URL)
	if err != nil {
		log.Fatal("failed to connect DB:", err)
	}

	// Cloudinary
	cloudStore, err := storage.NewCloudinaryStorage(config.Cloudinary)
	if err != nil {
		log.Fatal("cloudinary init error:", err)
	}

	// Service
	imgService := service.NewImageService(cloudStore, database)
	api.RegisterRoutes(r, imgService)
	port := config.Port
	log.Printf("Image service running at :%s", port)
	r.Run(":" + port)
}
