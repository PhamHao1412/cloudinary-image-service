package service

import (
	"fmt"
	"image-service/internal/model"
	"mime/multipart"
	"time"

	"image-service/internal/entity"
	"image-service/internal/storage"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IImageService interface {
	Upload(file *multipart.FileHeader) (*entity.Image, error)
	GetMetadata(id string) (*entity.Image, bool)
	Resize(req model.TransformRequest) (string, error)
	Convert(req model.TransformRequest) (string, error)
	Filter(req model.TransformRequest) (string, error)
	Crop(req model.TransformRequest) (string, error)
	Rotate(req model.TransformRequest) (string, error)
	Flip(req model.TransformRequest) (string, error)
	Watermark(req model.TransformRequest) (string, error)
	Compress(req model.TransformRequest) (string, error)
}

type ImageService struct {
	store storage.CloudStorage
	db    *gorm.DB
}

func NewImageService(store storage.CloudStorage, db *gorm.DB) *ImageService {
	return &ImageService{store: store, db: db}
}

func (s *ImageService) Upload(file *multipart.FileHeader) (*entity.Image, error) {
	id := uuid.NewString()

	url, format, err := s.store.UploadFile(file, id)
	if err != nil {
		return nil, err
	}

	img := &entity.Image{
		ID:         id,
		URL:        url,
		Format:     format,
		UploadedAt: time.Now(),
	}

	if err := s.db.Create(img).Error; err != nil {
		return nil, err
	}
	return img, nil
}

func (s *ImageService) GetMetadata(id string) (*entity.Image, bool) {
	var img entity.Image
	if err := s.db.First(&img, "id = ?", id).Error; err != nil {
		return nil, false
	}
	return &img, true
}

func (s *ImageService) Resize(req model.TransformRequest) (string, error) {
	var img entity.Image
	if err := s.db.First(&img, "id = ?", req.ID).Error; err != nil {
		return "", fmt.Errorf("image not found")
	}

	url, err := s.store.Resize(img.URL, req.Width, req.Height)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *ImageService) Convert(req model.TransformRequest) (string, error) {
	var img entity.Image
	if err := s.db.First(&img, "id = ?", req.ID).Error; err != nil {
		return "", fmt.Errorf("image not found")
	}

	url, err := s.store.Convert(img.URL, req.Format)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *ImageService) Filter(req model.TransformRequest) (string, error) {
	var img entity.Image
	if err := s.db.First(&img, "id = ?", req.ID).Error; err != nil {
		return "", fmt.Errorf("image not found")
	}

	url, err := s.store.ApplyFilter(img.URL, req.Filter)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *ImageService) Crop(req model.TransformRequest) (string, error) {
	var img entity.Image
	if err := s.db.First(&img, "id = ?", req.ID).Error; err != nil {
		return "", fmt.Errorf("image not found")
	}
	url, err := s.store.Crop(img.URL, req.Width, req.Height, req.X, req.Y)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (s *ImageService) Rotate(req model.TransformRequest) (string, error) {
	var img entity.Image
	if err := s.db.First(&img, "id = ?", req.ID).Error; err != nil {
		return "", fmt.Errorf("image not found")
	}
	url, err := s.store.Rotate(img.URL, req.Angle)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (s *ImageService) Flip(req model.TransformRequest) (string, error) {
	var img entity.Image
	if err := s.db.First(&img, "id = ?", req.ID).Error; err != nil {
		return "", fmt.Errorf("image not found")
	}
	url, err := s.store.Flip(img.URL, req.FlipAxis)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (s *ImageService) Watermark(req model.TransformRequest) (string, error) {
	var img entity.Image
	if err := s.db.First(&img, "id = ?", req.ID).Error; err != nil {
		return "", fmt.Errorf("image not found")
	}
	url, err := s.store.Watermark(img.URL, req.Watermark)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (s *ImageService) Compress(req model.TransformRequest) (string, error) {
	var img entity.Image
	if err := s.db.First(&img, "id = ?", req.ID).Error; err != nil {
		return "", fmt.Errorf("image not found")
	}
	url, err := s.store.Compress(img.URL, req.Quality)
	if err != nil {
		return "", err
	}
	return url, nil
}
