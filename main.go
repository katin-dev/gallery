package main

import (
	"github.com/gin-gonic/gin"
	httpfile "github.com/katin-dev/gallery/app/http/file"
)

func main() {
	r := gin.Default()

	r.POST("/api/v1/files", httpfile.UploadFile)
	r.GET("/api/v1/files", httpfile.ListFiles)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
