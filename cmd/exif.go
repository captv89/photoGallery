package cmd

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"unicode"

	"github.com/captv89/photoGallery/model"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

// WalkImageFiles walks the thumbs folder and returns the image file names
func WalkImageFiles(folder string) []model.Image {
	// Define a variable to store the image file names
	var imageFiles []model.Image

	// Walk the thumbs folder
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		// Check if the file is a directory
		if info.IsDir() {
			return nil
		}

		// Check if the file is an image
		if isImageFile(path) {
			// Read the image metadata
			image, err := ReadImageMetadata(path)
			if err != nil {
				return err
			}

			// Add the image to the list
			imageFiles = append(imageFiles, image)
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return imageFiles
}

// ReadImageMetadata reads the metadata from an image file
func ReadImageMetadata(path string) (model.Image, error) {
	log.Printf("Reading metadata from %s\n", path)

	var image model.Image

	// Open the image file
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error opening %s: %s\n", path, err)
		return image, err
	}
	defer file.Close()

	// Get Image Name
	image.Name = filepath.Base(path)

	// Decode the image file
	exif.RegisterParsers(mknote.All...)
	x, err := exif.Decode(file)
	if err != nil {

		if err == io.EOF {
			log.Printf("No EXIF data found in %s\n", path)
			return image, nil
		}

		log.Printf("Error decoding %s: %s\n", path, err)
		return image, err
	}

	// Print the values of all tags
	log.Println("Image Metadata:", x.String())

	// Get the image metadata

	// DateTime
	dateTime, _ := x.Get(exif.DateTime)
	// Check if dateTime is not nil
	if dateTime != nil {
		image.DateTime, _ = dateTime.StringVal()
	}

	// Make
	make, _ := x.Get(exif.Make)
	// Check if make is not nil
	if make != nil {
		image.Make, _ = make.StringVal()
	}

	// Model
	model, _ := x.Get(exif.Model)
	// Check nil
	if model != nil {
		image.Model, _ = model.StringVal()
	}

	// ExposureTime
	exposureTime, _ := x.Get(exif.ExposureTime)
	// Check nil
	if exposureTime != nil {
		image.ExposureTime = cleanFractionString(exposureTime.String())
	}

	// Aperture
	aperture, _ := x.Get(exif.ApertureValue)
	// Check nil
	if aperture != nil {
		image.Aperture = cleanFractionString(aperture.String())
	}

	// FNumber
	fNumber, _ := x.Get(exif.FNumber)
	// Check nil
	if fNumber != nil {
		image.FNumber = cleanFractionString(fNumber.String())
	}

	// ShutterSpeed
	shutterSpeed, _ := x.Get(exif.ShutterSpeedValue)
	// Check nil
	if shutterSpeed != nil {
		image.ShutterSpeed = cleanFractionString(shutterSpeed.String())
	}

	// FocalLength
	focalLength, _ := x.Get(exif.FocalLength)
	// Check nil
	if focalLength != nil {
		image.FocalLength = cleanFractionString(focalLength.String())
	}

	// ISO
	iso, _ := x.Get(exif.ISOSpeedRatings)
	// Check nil
	if iso != nil {
		image.ISO = iso.String()
	}

	// LensModel
	lensModel, _ := x.Get(exif.LensModel)
	// Check nil
	if lensModel != nil {
		image.LensModel, _ = lensModel.StringVal()
	}

	// XResolution
	xResolution, _ := x.Get(exif.XResolution)
	// Check nil
	if xResolution != nil {
		image.XResolution = cleanFractionString(xResolution.String())
	}

	// YResolution
	yResolution, _ := x.Get(exif.YResolution)
	// Check nil
	if yResolution != nil {
		image.YResolution = cleanFractionString(yResolution.String())
	}

	// ResolutionUnit
	resolutionUnit, _ := x.Get(exif.ResolutionUnit)
	// Check nil
	if resolutionUnit != nil {
		image.ResolutionUnit = resolutionUnit.String()
	}

	return image, nil
}

// GetImageMetaData(imgName) returns the metadata for the image file
func GetImageMetaData(imgName string) model.Image {
	var image model.Image
	for _, img := range model.Fulls {
		if img.Name == imgName {
			image = img
			break
		}
	}
	return image
}

// cleanFractionString cleans up a fraction string by removing any extraneous characters
func cleanFractionString(fraction string) string {
	// Remove any non-numeric characters except '/'
	cleaned := ""
	for _, char := range fraction {
		if unicode.IsDigit(char) || char == '/' {
			cleaned += string(char)
		}
	}
	return cleaned
}
