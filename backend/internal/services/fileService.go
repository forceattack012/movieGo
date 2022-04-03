package services

import (
	"backend/internal/domain"
	"image"
)

type FileService struct {
	repo *domain.FileRepository
}

func NewFileService(re *domain.FileRepository) domain.FileService {
	fileService := &FileService{
		repo: re,
	}
	return fileService
}

// GetImageFileFromPath implements domain.FileService
func (fs *FileService) GetImageFileFromPath(path string) (image.Image, error) {
	panic("unimplemented")
}
