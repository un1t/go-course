package main

import (
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/pressly/goose/v3"
)

type User struct {
	Id     int
	Name   string
	Email  string
	Photos []Photo
}

func (User) TableName() string {
	return "users"
}

type Photo struct {
	UserId    int
	Filename  string
	Width     int
	Height    int
	CreatedAt time.Time
}

func (Photo) TableName() string {
	return "photos"
}

func main() {
	os.Setenv("DATABASE_URL", "postgres://postgres:123@localhost:5432/go_dev")
	databaseUrl := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(sqlDB, "migrations"); err != nil {
		panic(err)
	}

}
