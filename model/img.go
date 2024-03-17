package model

type Image struct {
	Name         string
	DateTime     string
	Make         string
	Model        string
	ExposureTime string
	Aperture     string
	FNumber      string
	ShutterSpeed string
	// Flash string
	FocalLength    string
	ISO            string // exif stores this value in int64
	LensModel      string
	XResolution    string
	YResolution    string
	ResolutionUnit string // exif stores this value in int64
}

var Thumbnails []string
var Fulls []Image