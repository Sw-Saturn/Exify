package v1

import (
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"log"
	"math/big"
	"os"
)

//GetModel is returning a Camera Model.
func GetModel(x *exif.Exif) string {
	camModel, err := x.Get(exif.Model)
	result, err := camModel.StringVal()
	if err != nil {
		return ""
	}
	return result
}

//GetFocalLength is returning a FocalLength.
func GetFocalLength(x *exif.Exif) string {
	focal, err := x.Get(exif.FocalLength)
	if err != nil {
		return ""
	}
	numer, denom, _ := focal.Rat2(0)
	return fmt.Sprintf("%.1fmm", float64(numer)/float64(denom))
}

//GetLensModel is returning a LensModel.
func GetLensModel(x *exif.Exif) string {
	lensModel, err := x.Get(exif.LensModel)
	result, err := lensModel.StringVal()
	if err != nil {
		return ""
	}
	return result
}

//GetFocalLengthIn35mm is returning a FocalLength equivalent to 35mm focal length film.
func GetFocalLengthIn35mm(x *exif.Exif) string {
	focal35, err := x.Get(exif.FocalLengthIn35mmFilm)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%vmm", focal35)
}

//GetISOSpeedRatings is returning an ISOSpeedRatings.
func GetISOSpeedRatings(x *exif.Exif) string {
	iso, err := x.Get(exif.ISOSpeedRatings)
	if err != nil {
		return ""
	}
	return iso.String()
}

//GetFNumber is returning a FNumber.
func GetFNumber(x *exif.Exif) string {
	fnumber, err := x.Get(exif.FNumber)
	if err != nil {
		return ""
	}
	numer, denom, _ := fnumber.Rat2(0)
	return fmt.Sprintf("%.1f", float64(numer)/float64(denom))
}

//GetShutterSpeed is returning a shutter speed.
func GetShutterSpeed(x *exif.Exif) string {
	shutterspeed, err := x.Get(exif.ExposureTime)
	if err != nil {
		return ""
	}
	numer, denom, _ := shutterspeed.Rat2(0)
	ss := big.NewRat(numer, denom)
	return ss.String()
}

//GetDateTime is returning a datetime of shooting.
func GetDateTime(x *exif.Exif) string {
	tm, err := x.DateTime()
	if err != nil {
		return ""
	}
	return tm.String()
}

//ExifInit is initialize operation to use an exif module.
func ExifInit(f *os.File) (*exif.Exif, error) {
	exif.RegisterParsers(mknote.All...)
	x, err := exif.Decode(f)
	if err != nil {
		if exif.IsTagNotPresentError(err) {
			log.Println(err)
		} else if exif.IsExifError(err) {
			log.Println(err)
		} else {
			return nil, err
		}
	}
	return x, nil
}
