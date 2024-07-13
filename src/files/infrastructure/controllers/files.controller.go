package infrastrucure

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	domain "github.com/ingdeiver/go-core/src/files/domain/interfaces"
)

type FilesController struct {
	fileService domain.BaseFileService
}

func New (fileService domain.BaseFileService) *FilesController{
	return &FilesController{fileService}
}

func (s *FilesController) Upload(c *gin.Context) {
	// Get the file from the request
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.Error(err)
		return
	}
	defer file.Close()

	
	// Upload the file using the file service
	response, err := s.fileService.Upload(file, "files", header.Filename)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (s *FilesController) Get(c *gin.Context) {
	key := c.Param("key")
	response, err := s.fileService.Get("files",key)
	if err != nil {
		c.Error(err)
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", key))
	c.Header("Content-Type", response.ContentType)
	c.Header("Content-Length", strconv.FormatInt(*response.ContentLength, 10))
	c.Stream(func(w io.Writer) bool {
		_, err := io.Copy(w, response.Body)
		return err == nil
	})
}

func (s *FilesController) Remove(c *gin.Context) {
	key := c.Param("key")

	_, err := s.fileService.Remove("files", key)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File successfully deleted"})
}