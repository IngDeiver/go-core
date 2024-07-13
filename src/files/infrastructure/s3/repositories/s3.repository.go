package fileRepository

import (
	"context"
	"errors"
	"io"
	"mime"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	errorsDomain "github.com/ingdeiver/go-core/src/commons/domain/errors"
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



func (r *S3Repository) Upload(body io.Reader, folder string, fileName string) (domain.File, error) {
	key := folder + "/" + fileName

	// Detect the content type
	ext := filepath.Ext(fileName)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		// Default to binary/octet-stream if the content type cannot be determined
		contentType = "application/octet-stream"
	}
	_, err := r.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(key),
		Body:   body,
		ContentType: aws.String(contentType),
		//ACL:    types.ObjectCannedACLPublicRead,
	})

	if err != nil {
		l.Error().Err(err).Msg("Failed to upload file to S3")
		return domain.File{}, err
	}

	return domain.File{
		Name: key,
		Path:  "https://" + r.bucket + ".s3.amazonaws.com/" + key,
		Metadata: nil,
	}, nil
}

func (r *S3Repository) Remove(folder string, key string) (bool, error) {
	_, err := r.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(folder + "/" + key),
	})

	if err != nil {
		var notFoundErr *types.NotFound
		if errors.As(err, &notFoundErr) {
			l.Error().Err(errorsDomain.ErrNotFoundError)
			return false, errorsDomain.ErrNotFoundError
		}
		l.Error().Err(err).Msg("Failed to delete file from S3")
		return false, err
	}

	return true, nil
}

func (r *S3Repository) Get(folder string, key string) (domain.FileResponse, error) {
	resp, err := r.client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(folder + "/" + key),
	})

	
	if err != nil {
		var notFoundErr *types.NoSuchKey
		if errors.As(err, &notFoundErr) {
			l.Error().Err(errorsDomain.ErrNotFoundError)
			return domain.FileResponse{}, errorsDomain.ErrNotFoundError
		}
		l.Error().Err(err).Msg("Failed to get file from S3")
		return domain.FileResponse{}, err
	}

	return domain.FileResponse{
		Body: resp.Body,
		ContentType: aws.ToString(resp.ContentType),
		ContentLength: resp.ContentLength,
	}, nil
}