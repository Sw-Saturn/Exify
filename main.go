package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"picture-exif-api/controllers/api/v1"
)

func main() {
	if len(os.Args) == 2 {
		cli()
	} else {
		r := gin.Default()
		r.MaxMultipartMemory = 8 << 20
		r.GET("/", func(context *gin.Context) {
			context.File("./template/index.html")
		})
		r.POST("/api/v1/getexif", v1.GetExifHandler)
		r.Run(":8080")
	}
}

func cli() {
	frame := os.Args[1]
	println(v1.GetExif(frame))
}
