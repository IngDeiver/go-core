package fileDomain

import (
	"io"

	domain "github.com/ingdeiver/go-core/src/files/domain"
)

type BaseFileService interface {
	Upload(body io.Reader, folder string, fileName string) (domain.File, error)
	Remove(folder string, key string) (bool, error) 
	Get(folder string, key string) (domain.FileResponse, error)
}
