package fileService

import (
	"io"

	domain "github.com/ingdeiver/go-core/src/files/domain"
	interfaces "github.com/ingdeiver/go-core/src/files/domain/interfaces"
)

// implements BaseFileService
type S3Service struct {
	fileRepository  interfaces.BaseFileRepository
}

func New(fileRepository interfaces.BaseFileRepository)  *S3Service {
	return &S3Service{fileRepository}
}

func (service *S3Service) Upload(body io.Reader, folder string, fileName string) (domain.File, error){
	return service.fileRepository.Upload(body,folder, fileName)
}

func (service *S3Service) Remove(folder string, key string) (bool, error) {
	return service.fileRepository.Remove(folder, key)
}

func (service *S3Service)  Get(folder string,key string) (domain.FileResponse, error) {
	return service.fileRepository.Get(folder, key)
}