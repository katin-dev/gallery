package app

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	d "github.com/katin-dev/gallery/app/domain/file"
	"github.com/katin-dev/gallery/app/http/file"
)

type Conf struct {
	Db   ConfDb
	Port string `env:"APP_LISTEN_PORT" envDefault:"8080"`
}

type ConfDb struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
}

type App struct {
	Conf     Conf
	FileRepo d.FileRepository
	s3client *s3.Client
}

func NewApp(c Conf, fileRepository d.FileRepository, s3client *s3.Client) *App {
	return &App{
		Conf:     c,
		FileRepo: fileRepository,
		s3client: s3client,
	}
}

func (a *App) Run() {
	r := gin.Default()

	controllerFile := file.NewFilesHttpController(
		file.NewFileRestService(a.FileRepo, a.s3client),
	)

	r.POST("/api/v1/files", controllerFile.Upload)
	r.GET("/api/v1/files", controllerFile.List)

	r.Run(":" + a.Conf.Port)
}
