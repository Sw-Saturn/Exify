package main

import (
	"log"
	"os"
	"picture-exif-api/controllers"
)

func main() {
	if len(os.Args) == 2 {
		frame := os.Args[1]

		f, err := os.Open(frame)
		if err != nil {
			log.Fatal(err)
		}
		x, err := controllers.ExifInit(f)
		if err != nil {
			log.Fatal(err)
		}

		model := controllers.GetModel(x)
		println(model)

		focal := controllers.GetFocalLength(x)
		println(focal)

		focalin35mm := controllers.GetFocalLengthIn35mm(x)
		println(focalin35mm)

		iso := controllers.GetISOSpeedRatings(x)
		println(iso)

		fnum := controllers.GetFNumber(x)
		println(fnum)

		ss := controllers.GetShutterSpeed(x)
		println(ss)

		dt := controllers.GetDateTime(x)
		println(dt)
	}
}
