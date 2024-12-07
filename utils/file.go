package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

const maxMemory = int64(1 * 1024 * 1024)

func ParseFileRequest(c *gin.Context, key string, id string, fileType string) string {

	// Multipart Form
	form, err := c.MultipartForm()
	if err != nil {
		ErrorResponse(err)
	}

	// take first
	files := form.File[key]
	fmt.Println(files)

	if len(files) == 0 {
		ErrorResponse(errors.New(fmt.Sprintf("file %s not exist", key)))
		return ""
	}

	// path and naming
	filename := ""
	for _, file := range files {
		filename = fmt.Sprintf("%s%s/%s%s.%s", "files/", key, id, ".", filepath.Ext(file.Filename))

		if fileType == "image" {
			err = ImageChecker(file)
			if err != nil {
				ErrorResponse(err)
				return ""
			}
		}

		fmt.Println(filename)
		if err = c.SaveUploadedFile(file, filename); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return ""
		}
	}
	return filename
}

func ImageChecker(file *multipart.FileHeader) error {
	allowedExt := []string{".jpg", ".jpeg", ".png"}

	isAllowed := false
	for _, ext := range allowedExt {
		if filepath.Ext(file.Filename) == ext {
			isAllowed = true
		}
	}

	if !isAllowed {
		//c.JSON(http.StatusBadRequest, ErrorResponse(errors.New("file extension not allowed")))
		return errors.New("file extension not allowed")
	}
	if file.Size > maxMemory {
		//c.JSON(http.StatusBadRequest, ErrorResponse(errors.New("file size too big")))
		return errors.New("file size too big")
	}
	return nil
}
