package fileRepository

import (
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	logger "github.com/ingdeiver/go-core/src/commons/infrastructure/logs"
	domain "github.com/ingdeiver/go-core/src/files/domain"
)

var l = logger.Get()

// implements BaseFileRepository
type S3Repository struct {
	bucket string
	client  *s3.Client
}

func New () *S3Repository {
	bucket := os.Getenv("S3_MAIN_BUCKET")

	// Load the from ENV AWS config
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
        l.Fatal().Err(err)
    }

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)
	return &S3Repository{bucket, client}
}



func (*S3Repository) Upload(body io.Reader, folder string, fileName string) (domain.File, error){
	return domain.New(), nil
}

func (*S3Repository) Remove(path string) (domain.File, error) {
	return domain.New(), nil
}

func (*S3Repository)  Get(path string) (domain.FileResponse, error) {
	return domain.FileResponse{}, nil
}