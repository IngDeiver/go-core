package fileDomain

import (
	"io"

	domain "github.com/ingdeiver/go-core/src/files/domain"
)

type BaseFileRepository interface {
	Upload(body io.Reader, folder string, fileName string) (domain.File, error)
	Remove(path string) (domain.File, error) 
	Get(path string) (domain.FileResponse, error)
}