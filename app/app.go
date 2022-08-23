package app

import (
	"github.com/gin-gonic/gin"
	d "github.com/katin-dev/gallery/app/domain/file"
	"github.com/katin-dev/gallery/app/http/auth"
	"github.com/katin-dev/gallery/app/http/file"
	"github.com/minio/minio-go/v7"
)

type Conf struct {
	Db   ConfDb
	S3   S3Conf
	Port string `env:"APP_LISTEN_PORT" envDefault:"8080"`
	Auth AuthConf
}

type ConfDb struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
}

type S3Conf struct {
	Endpoint string `env:"AWS_ENDPOINT"`
	Key      string `env:"AWS_API_KEY"`
	Secret   string `env:"AWS_API_SECRET"`
	UseSSL   bool   `env:"AWS_SSL"`
	Bucket   string `env:"AWS_BUCKET"`
}

type AuthConf struct {
	Host string `env:"AUTH_HOST"`
}

type App struct {
	Conf     Conf
	FileRepo d.FileRepository
	s3client *minio.Client
}

func NewApp(c Conf, fileRepository d.FileRepository, s3client *minio.Client) *App {
	return &App{
		Conf:     c,
		FileRepo: fileRepository,
		s3client: s3client,
	}
}

func (a *App) Run() {
	r := gin.Default()

	controllerFile := file.NewFilesHttpController(
		file.NewFileRestService(a.FileRepo, a.s3client, a.Conf.S3.Bucket),
	)

	auth := auth.NewRtbAuth(a.Conf.Auth.Host)
	r.Use(auth.Auth)

	private := r.Group("/api/v1", auth.Auth)
	private.POST("/files", controllerFile.Upload)
	private.GET("/files", controllerFile.List)

	r.Run(":" + a.Conf.Port)
}
