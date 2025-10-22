package storage

import (
	"fmt"
	"image"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

type Storage interface {
	SaveFile(file *multipart.FileHeader, id string) (string, error)
	SaveImage(img image.Image, relPath string) error
	FullPath(rel string) string
}

type LocalStorage struct {
	root string
}

func NewLocalStorage(root string) *LocalStorage {
	return &LocalStorage{root: root}
}

func (l *LocalStorage) SaveFile(file *multipart.FileHeader, id string) (string, error) {
	ext := filepath.Ext(file.Filename)
	path := fmt.Sprintf("originals/%s%s", id, ext)
	full := filepath.Join(l.root, path)

	if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(full)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return path, nil
}

func (l *LocalStorage) SaveImage(img image.Image, relPath string) error {
	full := filepath.Join(l.root, relPath)
	if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
		return err
	}
	return imaging.Save(img, full)
}

func (l *LocalStorage) FullPath(rel string) string {
	return filepath.Join(l.root, rel)
}
