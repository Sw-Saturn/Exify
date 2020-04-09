package v1

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
	"net/http"
	"os"
)

func getPicture(ctx *gin.Context) string{
	img, err := imageupload.Process(ctx.Request, "file")
	if err != nil {
		panic(err)
	}
	h := generateFileHash(img.Data)
	filePath := fmt.Sprintf("./upload/%s.jpg", h)
	err = img.Save(filePath)
	if err != nil {
		panic(err)
	}
	return filePath
}

func GetExifHandler(ctx *gin.Context) {
	img := getPicture(ctx)
	exifStr := GetExif(img)
	ctx.JSON(http.StatusOK, gin.H{
		"exif": exifStr,
	})
}

func GetExif(img string) string{
	f, err := os.Open(img)
	if err != nil {
		panic(err)
	}
	x, err := ExifInit(f)
	if err != nil {
		panic(err)
	}

	model := GetModel(x)
	println(model)

	focal := GetFocalLength(x)
	println(focal)

	focalin35mm := GetFocalLengthIn35mm(x)
	println(focalin35mm)

	iso := GetISOSpeedRatings(x)
	println(iso)

	fnum := GetFNumber(x)
	println(fnum)

	ss := GetShutterSpeed(x)
	println(ss)

	dt := GetDateTime(x)
	println(dt)

	return fmt.Sprintf("%s %s %s %s %s", model, iso, fnum, focal, ss)
}

func generateFileHash(file []byte) string {
	h := sha256.New()
	h.Write(file)
	bs := h.Sum(nil)

	return hex.EncodeToString(bs)
}