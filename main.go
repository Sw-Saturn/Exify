package main

import (
	"exify/configs"
	"exify/controllers/api/v1"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	configs.LoadConfig()
	r := gin.Default()
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")
	r.StaticFile("/android-chrome-512x512.png", "./assets/android-chrome-512x512.png")
	r.StaticFile("/apple-touch-icon.png", "./assets/apple-touch-icon.png")
	r.MaxMultipartMemory = 8 << 20
	r.LoadHTMLFiles("./template/index.html")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.POST("/api/v1/getexif", v1.GetExifHandler)
	r.Run(GetAddress())
}

func GetAddress() string {
	serverAddress := fmt.Sprintf("%s:%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)
	return serverAddress
}