package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picture-exif-api/controllers/api/v1"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Static("/resources", "./resources")
	r.MaxMultipartMemory = 8 << 20
	r.LoadHTMLGlob("template/*.html")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.POST("/api/v1/getexif", v1.GetExifHandler)
	r.Run(":8080")
}