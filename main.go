package main

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	httpfile "github.com/katin-dev/gallery/app/http/file"
)

type Conf struct {
	Db ConfDb
}

type ConfDb struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	cfg := Conf{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("Failed to read config: %+v\n", err)
		os.Exit(1)
	}

	fmt.Println(cfg)

	r := gin.Default()

	r.POST("/api/v1/files", httpfile.UploadFile)
	r.GET("/api/v1/files", httpfile.ListFiles)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
