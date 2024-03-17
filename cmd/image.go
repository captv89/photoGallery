package cmd

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/captv89/photoGallery/model"
)

// isImageFile checks if the file has a recognized image extension
func isImageFile(path string) bool {
	switch filepath.Ext(path) {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	}
	return false
}

// SaveImageNames saves the image file names to a CSV file
func SaveImageNames(imageFiles []string) {
	// Create a file
	file, err := os.Create("imageNames.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the image file names to the CSV file
	for _, name := range imageFiles {
		err := writer.Write([]string{name})
		if err != nil {
			log.Fatal(err)
		}
	}
}

// SaveImageMetadata saves the image file metadata to a CSV file
func SaveImageMetadata(images []model.Image, fileName string) {
	// Create a file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the image file metadata to the CSV file
	for _, image := range images {
		err := writer.Write([]string{image.Name, image.DateTime, image.Make, image.Model, image.ExposureTime, image.Aperture, image.FNumber, image.ShutterSpeed, image.FocalLength, image.ISO, image.LensModel, image.XResolution, image.YResolution, image.ResolutionUnit})
		if err != nil {
			log.Fatal(err)
		}
	}
}

// LoadImageNames loads the image file names from a file
func LoadImageNames(fName string) []string {
	// Open the file
	file, err := os.Open(fName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Define a variable to store the image file names
	var imageFiles []string

	// Read the image file names from the CSV file
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		imageFiles = append(imageFiles, record[0])
	}

	return imageFiles
}

// LoadImageNames loads the image file names from a file
func LoadImageMetadata(fName string) []model.Image {
	// Open the file
	file, err := os.Open(fName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Define a variable to store the image file names
	var imageFiles []model.Image

	// Read the image file names from the CSV file
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		imageFiles = append(imageFiles, model.Image{Name: record[0], DateTime: record[1], Make: record[2], Model: record[3], ExposureTime: record[4], Aperture: record[5], FNumber: record[6], ShutterSpeed: record[7], FocalLength: record[8], ISO: record[9], LensModel: record[10], XResolution: record[11], YResolution: record[12], ResolutionUnit: record[13]})
	}

	return imageFiles
}

// CountImagesInFolder counts the number of image files in the folder
func CountImagesInFolder(folder string) int {
	// Define a variable to store the count
	count := 0

	// Open the folder
	dir, err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	// Read the files in the folder
	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	// Count the image files in the folder
	for _, file := range files {
		if !file.IsDir() && isImageFile(file.Name()) {
			count++
		}
	}

	return count
}

// GetImageFilesInFolder gets the image files in the folder
func GetImageFilesInFolder(folder string) []string {
	// Define a variable to store the image files
	var imageFiles []string

	// Walk through the folder and add the image files to the list
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		// Check if the file is a regular file and has a recognized image extension
		if err == nil && !info.IsDir() && isImageFile(path) {
			// Add the image file to the list
			fName := filepath.Base(path)
			imageFiles = append(imageFiles, fName)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return imageFiles
}

// GetPreviousAndNextImage gets the previous and next image file names
func GetPreviousAndNextImage(imgName string) (string, string) {
	// Load the image file names from the CSV file
	imageFileNames := model.Thumbnails

	// Define a variable to store the previous and next image file names
	var previous, next string

	// Find the index of the current image file name
	for i, name := range imageFileNames {
		if name == imgName {
			// Get the previous and next image file names
			if i > 0 {
				previous = imageFileNames[i-1]
			}
			if i < len(imageFileNames)-1 {
				next = imageFileNames[i+1]
			}
		}
	}

	return previous, next
}
