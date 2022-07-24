package main

import (
	"fmt"
	"log"
	"os"

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
		panic("failed to connect database")
	}

	fileRepo := file.NewPostgresFileRepository(db)

	app := app.NewApp(cfg, fileRepo)
	app.Run()
}
