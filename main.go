package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/katin-dev/gallery/app"
	"github.com/katin-dev/gallery/app/infra/file"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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
		log.Fatalf("Failed to read config: %+v\n", err)
	}

	dbc := cfg.Db
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow", dbc.Host, dbc.User, dbc.Password, dbc.Name, dbc.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection failed: %+v\n", err)
	}

	fileRepo := file.NewPostgresFileRepository(db)

	// S3 Client
	s3client, err := minio.New(cfg.S3.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(
			cfg.S3.Key,
			cfg.S3.Secret,
			"",
		),
		Secure: cfg.S3.UseSSL,
	})
	if err != nil {
		log.Fatalf("Failed to create S3 client: %s\n", err)
	}

	fmt.Println("[s3 bucket] Try to create buckets...")
	bucketName := cfg.S3.Bucket
	bucketOptions := minio.MakeBucketOptions{Region: "ru-msk-zel"}

	bucketTry := 3
	for err := s3client.MakeBucket(context.Background(), bucketName, bucketOptions); err != nil && bucketTry != 0; {
		bucketTry--
		time.Sleep(time.Second)
		fmt.Printf("[s3 bucket] Failed: %s\n", err)
	}
	fmt.Println("[s3 bucket] Done!")

	app := app.NewApp(cfg, fileRepo, s3client)
	app.Run()
}
