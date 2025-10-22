package api

import (
	"image-service/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, imgSvc *service.ImageService) {
	h := NewHandler(imgSvc)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/upload", h.Upload)
		v1.POST("/resize", h.Resize)
		v1.POST("/convert", h.Convert)
		v1.POST("/filter", h.Filter)
		v1.GET("/images/:id", h.GetImage)
		v1.POST("/crop", h.Crop)
		v1.POST("/rotate", h.Rotate)
		v1.POST("/flip", h.Flip)
		v1.POST("/watermark", h.Watermark)
		v1.POST("/compress", h.Compress)
	}
}
