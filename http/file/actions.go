package file

import (
	"net/http"

	"github.com/gin-gonic/gin"
	infra_file "github.com/katin-dev/gallery/infra/file"
)

func ListFiles(c *gin.Context) {
	service := FileRestService{
		infra_file.NewFileRepository(),
	}

	list := service.getList()

	c.JSON(http.StatusOK, gin.H{
		"rows":  list.files,
		"total": list.total,
	})
}

func UploadFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
}
