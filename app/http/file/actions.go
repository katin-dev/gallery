package file

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FilesHttpController struct {
	service *FileRestService
}

func NewFilesHttpController(fileRestService *FileRestService) *FilesHttpController {
	return &FilesHttpController{fileRestService}
}

func (c *FilesHttpController) List(ginCtx *gin.Context) {
	list, err := c.service.getList()
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginCtx.Header("X-Total-Count", strconv.Itoa(int(list.total)))
	ginCtx.JSON(http.StatusOK, list.files)
}

func (c *FilesHttpController) Upload(ginCtx *gin.Context) {
	file, err := ginCtx.FormFile("file")
	if err != nil {
		log.Printf("Failed to upload file: %s\n", err)
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	uploadedFile, err := file.Open()
	if err != nil {
		log.Printf("Failed to upload file: %s\n", err)
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tmpFile, err := os.CreateTemp("/tmp", "gal_")
	if err != nil {
		log.Printf("Failed to upload file: %s\n", err)
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = io.Copy(tmpFile, uploadedFile)
	if err != nil {
		log.Printf("Failed to upload file: %s\n", err)
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fileDto, err := c.service.UploadFile(tmpFile.Name(), file.Filename, file.Size)
	if err != nil {
		log.Printf("Failed to upload file: %s\n", err)
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginCtx.JSON(http.StatusOK, fileDto)
}
