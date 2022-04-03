package repositories

import (
	"backend/internal/domain"
	"image"
	"os"
)

type FileRepository struct {
}

func NewFileRepository() domain.FileRepository {
	return &FileRepository{}
}

func (fr *FileRepository) GetImageFileFromPath(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}
