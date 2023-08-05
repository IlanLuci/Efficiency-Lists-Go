package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	godotenv.Load()

	r := gin.Default()

	// serve static files
	r.Static("/ui", "./static")

	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/ui")
	})

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// connect to db
	println("Connecting to Databse")
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + name + " port=" + port + " sslmode=disable TimeZone=America/New_York"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	println("Syncing Database...")
	db.AutoMigrate(&User{})

	api := r.Group("/api/v1")
	api.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to the Efficiency Lists API")
	})

	api.POST("/login", login)
	api.GET("/data", getData)
	api.POST("/data", setData)

	if err != nil {
		println(err)
		return
	}

	r.Run(":80")
}
