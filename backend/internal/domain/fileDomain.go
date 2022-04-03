package domain

import "image"

type FileService interface {
	GetImageFileFromPath(path string) (image.Image, error)
}

type FileRepository interface {
	GetImageFileFromPath(path string) (image.Image, error)
}
