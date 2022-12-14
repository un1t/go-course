package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

var MyLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Info, // Log level
		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
		Colorful:                  true,        // Disable color
	},
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		panic("DATABASE_URL is empty")
	}

	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{Logger: MyLogger})
	if err != nil {
		panic("failed to connect database")
	}

	// user, err := GetUser(db, 222)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("user: %+v\n", user)

	users, err := GetUsers(db)
	if err != nil {
		panic(err)
	}
	PrintJson(users)

	// userId, err := InsertUser(db, User{Name: "AAA", Email: "aaa@bbb.cc"})
	// if err != nil {
	// 	panic(err)Photo
	// }
	// fmt.Printf("new user id: %d\n", userId)

	// rowsAffected, err := DeleteUser(db, 2)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("rows deleted: ", rowsAffected)

}

func PrintJson(v any) {
	bytes, _ := json.MarshalIndent(v, "", "    ")
	fmt.Println(string(bytes))
}

func GetUser(db *gorm.DB, userId int) (User, error) {
	var user User
	err := db.Take(&user, userId).Error
	return user, err
}

func GetUserByName(db *gorm.DB, name string) (User, error) {
	var user User
	err := db.Where("name = ?", name).Take(&user).Error
	return user, err
}

func GetUsers(db *gorm.DB) ([]User, error) {
	users := make([]User, 0)
	err := db.Preload("Photos").Find(&users).Error
	return users, err
}

func InsertUser(db *gorm.DB, user User) (int, error) {
	err := db.Create(&user).Error
	return user.Id, err
}

func DeleteUser(db *gorm.DB, userId int) (int, error) {
	tx := db.Delete(&User{}, userId)
	return int(tx.RowsAffected), tx.Error
}
