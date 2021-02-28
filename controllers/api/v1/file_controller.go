package v1

import (
	"encoding/json"
	"exify/model"
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
	"net/http"
	"os"
)

func getPicture(ctx *gin.Context) string {
	img, err := imageupload.Process(ctx.Request, "file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return ""
	}
	filePath := "./upload/tmp"
	err = img.Save(filePath)
	if err != nil {
		panic(err)
	}
	return filePath
}

//GetExifHandler is the handler of ImageUpload.
func GetExifHandler(ctx *gin.Context) {
	img := getPicture(ctx)
	exifStr := GetExif(img)
	ctx.JSON(http.StatusOK, exifStr)
}

//GetExif is returning the EXIF of an Image.
func GetExif(img string) string {
	f, err := os.Open(img)
	if err != nil {
		panic(err)
	}
	x, err := ExifInit(f)
	if err != nil {
		panic(err)
	}
	exif := model.Exif{}
	exif.ModelName = GetModel(x)
	exif.Iso = GetISOSpeedRatings(x)
	exif.FNumber = GetFNumber(x)
	exif.LensModel = GetLensModel(x)
	exif.FocalLength = GetFocalLength(x)
	exif.ShutterSpeed = GetShutterSpeed(x)

	result, err := json.Marshal(exif)
	if err != nil {
		panic(err)
	}
	print(string(result))
	return string(result)
}
