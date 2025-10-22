package api

import (
	"image-service/internal/model"
	"net/http"

	"image-service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	imgSvc service.IImageService
}

func NewHandler(svc service.IImageService) *Handler {
	return &Handler{imgSvc: svc}
}

func (h *Handler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing file"})
		return
	}

	meta, err := h.imgSvc.Upload(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, meta)
}

func (h *Handler) Resize(c *gin.Context) {
	var req model.TransformRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	url, err := h.imgSvc.Resize(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func (h *Handler) Convert(c *gin.Context) {
	var req model.TransformRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	url, err := h.imgSvc.Convert(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func (h *Handler) Filter(c *gin.Context) {
	var req model.TransformRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	url, err := h.imgSvc.Filter(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func (h *Handler) GetImage(c *gin.Context) {
	id := c.Param("id")
	meta, ok := h.imgSvc.GetMetadata(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
		return
	}
	c.JSON(http.StatusOK, meta)
}

func (h *Handler) Crop(c *gin.Context) {
	var req model.TransformRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	url, err := h.imgSvc.Crop(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func (h *Handler) Rotate(c *gin.Context) {
	var req model.TransformRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	url, err := h.imgSvc.Rotate(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func (h *Handler) Flip(c *gin.Context) {
	var req model.TransformRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	url, err := h.imgSvc.Flip(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func (h *Handler) Watermark(c *gin.Context) {
	var req model.TransformRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	url, err := h.imgSvc.Watermark(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func (h *Handler) Compress(c *gin.Context) {
	var req model.TransformRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	url, err := h.imgSvc.Compress(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": url})
}
