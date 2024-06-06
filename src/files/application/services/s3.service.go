package fileService

import (
	"io"

	domain "github.com/ingdeiver/go-core/src/files/domain"
	interfaces "github.com/ingdeiver/go-core/src/files/domain/interfaces"
	repositories "github.com/ingdeiver/go-core/src/files/infrastructure/s3/repositories"
)

// implements BaseFileService
type S3Service struct {
	FileRepository  interfaces.BaseFileRepository
}

func New()  *S3Service {
	fileRepo := repositories.New()
	return &S3Service{FileRepository: fileRepo}
}

func (*S3Service) Upload(body io.Reader, folder string, fileName string) (domain.File, error){
	return domain.New(), nil
}

func (*S3Service) Remove(path string) (domain.File, error) {
	return domain.New(), nil
}

func (*S3Service)  Get(path string) (domain.FileResponse, error) {
	return domain.FileResponse{}, nil
}