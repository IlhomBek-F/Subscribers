package database

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"subscribers/api/route"
	"subscribers/model"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	database   string
	password   string
	username   string
	port       string
	host       string
	schema     string
	dbInstance *gorm.DB
)

func connectDB() *gorm.DB {
	loadEnv()

	if dbInstance != nil {
		return dbInstance
	}

	te := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s", host, username, password, database, port, schema)

	db, err := gorm.Open(postgres.Open(te), &gorm.Config{})

	if err != nil {
		log.Fatal("Error while connecting to db")
	}

	return db
}

func InitServer() *http.Server {
	db := connectDB()

	db.AutoMigrate(&model.User{})

	portToInt, _ := strconv.Atoi(os.Getenv("PORT"))

	server := route.Server{
		Port: portToInt,
		DB:   db,
	}

	serverConfig := http.Server{
		Addr:         fmt.Sprintf(":%d", server.Port),
		Handler:      server.RegisterRoute(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return &serverConfig
}

func loadEnv() {
	err := godotenv.Load("../.env")

	if err != nil {
		fmt.Println("Error while loading env file")
	}

	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port = os.Getenv("DB_PORT")
	host = os.Getenv("DB_HOST")
	schema = os.Getenv("DB_SCHEMA")
}
