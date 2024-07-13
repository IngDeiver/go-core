package fileDomain

import (
	"io"
)

type FileResponse struct {
	Body io.Reader
	ContentType string
	ContentLength *int64
}
