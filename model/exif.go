package model

//Exif is the model of exif information.
type Exif struct {
	ModelName    string `json:"model_name"`
	LensModel    string `json:"lens_model"`
	Iso          string `json:"iso"`
	FNumber      string `json:"f_number"`
	FocalLength  string `json:"focal_length"`
	ShutterSpeed string `json:"shutter_speed"`
}
