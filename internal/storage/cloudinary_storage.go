package storage

import (
	"context"
	"fmt"
	"image-service/internal/app"
	"mime/multipart"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudStorage interface {
	UploadFile(file interface{}, id string) (string, string, error)
	Resize(url string, width, height int) (string, error)
	Convert(url, format string) (string, error)
	ApplyFilter(url, filter string) (string, error)
	Crop(url string, width, height, x, y int) (string, error)
	Rotate(url string, angle int) (string, error)
	Flip(url, axis string) (string, error)
	Watermark(url, overlay string) (string, error)
	Compress(url, quality string) (string, error)
}

type CloudinaryStorage struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryStorage(cfg app.CloudinaryConfig) (*CloudinaryStorage, error) {
	cld, err := cloudinary.NewFromParams(
		cfg.CloudName,
		cfg.ApiKey,
		cfg.Secret,
	)
	if err != nil {
		return nil, err
	}
	return &CloudinaryStorage{cld: cld}, nil
}

func (s *CloudinaryStorage) UploadFile(file interface{}, id string) (string, string, error) {
	fh, ok := file.(*multipart.FileHeader)
	if !ok {
		return "", "", fmt.Errorf("invalid file type")
	}
	ctx := context.Background()
	f, err := fh.Open()
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	res, err := s.cld.Upload.Upload(ctx, f, uploader.UploadParams{
		PublicID: id,
		Folder:   "image-service/originals",
	})
	if err != nil {
		return "", "", err
	}
	return res.SecureURL, res.Format, nil
}

func (s *CloudinaryStorage) Resize(url string, width, height int) (string, error) {
	transformation := fmt.Sprintf("w_%d,h_%d,c_fill", width, height)
	return s.transformURL(url, transformation), nil
}

func (s *CloudinaryStorage) Convert(url, format string) (string, error) {
	format = strings.ToLower(format)
	return s.transformURL(url, fmt.Sprintf("f_%s", format)), nil
}

func (s *CloudinaryStorage) ApplyFilter(url, filter string) (string, error) {
	switch strings.ToLower(filter) {
	case "grayscale":
		return s.transformURL(url, "e_grayscale"), nil
	case "blur":
		return s.transformURL(url, "e_blur:200"), nil
	case "sharpen":
		return s.transformURL(url, "e_sharpen"), nil
	default:
		return "", fmt.Errorf("unsupported filter: %s", filter)
	}
}

func (s *CloudinaryStorage) transformURL(originalURL, transformation string) string {
	parts := strings.Split(originalURL, "/upload/")
	if len(parts) != 2 {
		return originalURL
	}
	rs := fmt.Sprintf("%s/upload/%s/%s", parts[0], transformation, parts[1])
	return rs
}

func (s *CloudinaryStorage) Crop(url string, width, height, x, y int) (string, error) {
	tf := fmt.Sprintf("c_crop,w_%d,h_%d,x_%d,y_%d", width, height, x, y)
	return s.transformURL(url, tf), nil
}

func (s *CloudinaryStorage) Rotate(url string, angle int) (string, error) {
	tf := fmt.Sprintf("a_%d", angle)
	return s.transformURL(url, tf), nil
}

func (s *CloudinaryStorage) Flip(url, axis string) (string, error) {
	switch axis {
	case "horizontal":
		return s.transformURL(url, "a_hflip"), nil
	case "vertical":
		return s.transformURL(url, "a_vflip"), nil
	default:
		return "", fmt.Errorf("invalid flip axis: %s", axis)
	}
}

func (s *CloudinaryStorage) Watermark(url, overlay string) (string, error) {
	if overlay == "" {
		return "", fmt.Errorf("missing watermark public_id")
	}
	tf := fmt.Sprintf("l_%s,g_south_east,x_10,y_10,w_150", overlay)
	return s.transformURL(url, tf), nil
}

func (s *CloudinaryStorage) Compress(url, quality string) (string, error) {
	if quality == "" {
		quality = "auto"
	}
	tf := fmt.Sprintf("q_%s,f_auto", quality)
	return s.transformURL(url, tf), nil
}
