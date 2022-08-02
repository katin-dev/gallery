package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/katin-dev/gallery/app"
	"github.com/katin-dev/gallery/app/infra/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	cfg := app.Conf{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("Failed to read config: %+v\n", err)
		os.Exit(1)
	}

	dbc := cfg.Db
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow", dbc.Host, dbc.User, dbc.Password, dbc.Name, dbc.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Database connection failed: %+v\n", err)
		os.Exit(1)
	}

	fileRepo := file.NewPostgresFileRepository(db)

	s3config, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		fmt.Printf("AWS S3 configuration failes: %+v\n", err)
		os.Exit(1)
	}

	s3client := s3.NewFromConfig(s3config)

	app := app.NewApp(cfg, fileRepo, s3client)
	app.Run()
}
