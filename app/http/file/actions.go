package file

import (
	"log"
	"net/http"

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

	ginCtx.JSON(http.StatusOK, gin.H{
		"rows":  list.files,
		"total": list.total,
	})
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

	osFile, err := file.Open()
	if err != nil {
		log.Printf("Failed to upload file: %s\n", err)
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fileDto := c.service.UploadFile(osFile, file.Filename)

	ginCtx.JSON(http.StatusOK, fileDto)
}
