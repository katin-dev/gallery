package file

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	infra_file "github.com/katin-dev/gallery/app/infra/file"
)

var service = NewFileRestService()

func ListFiles(c *gin.Context) {
	list := service.getList()

	c.JSON(http.StatusOK, gin.H{
		"rows":  list.files,
		"total": list.total,
	})
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("Failed to upload file: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	osFile, err := file.Open()
	if err != nil {
		log.Printf("Failed to upload file: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fileDto := service.UploadFile(osFile, file.Filename)

	c.JSON(http.StatusOK, fileDto)
}

func NewFileRestService() *FileRestService {
	return &FileRestService{
		infra_file.NewFileRepository(),
	}
}
