package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

const maxMemory = int64(1 * 1024 * 1024)

func ImageChecker(c *gin.Context, file *multipart.FileHeader) {
	allowedExt := []string{".jpg", ".jpeg", ".png"}

	isAllowed := false
	for _, ext := range allowedExt {
		if filepath.Ext(file.Filename) == ext {
			isAllowed = true
		}
	}

	if !isAllowed {
		c.JSON(http.StatusBadRequest, ErrorResponse(errors.New("file extension not allowed")))
		return
	}
	if file.Size > maxMemory {
		c.JSON(http.StatusBadRequest, ErrorResponse(errors.New("file size too big")))
		return
	}
	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(errors.New("error saving file")))
		return
	}
}
