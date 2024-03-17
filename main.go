package main

import (
	"log"
	"os"

	"github.com/captv89/photoGallery/cmd"
	"github.com/captv89/photoGallery/model"
	"github.com/captv89/photoGallery/web"

	"github.com/joho/godotenv"
)

func init() {
	// Log the line number and file name of the error
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Run the action
	action()

	// Start the server
	web.StartServer()
}

func action() {
	// Image Thumbnails List
	thumbsPath := "web/static/photos/thumbs"
	fullsPath := "web/static/photos/fulls"

	///////////////////////////////////////////////// FULLS //////////////////////////////////////////////////
	// Check if the fulls folder exists
	if _, err := os.Stat(fullsPath); os.IsNotExist(err) {
		// Exit if the fulls folder does not exist
		log.Fatal("Fulls folder does not exist")
	}

	var fCount int
	if os.Getenv("SKIP_LOADING") == "false" {
		// Count the number of image files in the fulls folder
		fCount = cmd.CountImagesInFolder(fullsPath)
		log.Printf("Image file count in directory: %d\n", fCount)

		// Check if the image file count is zero
		if fCount == 0 {
			// Quit if there are no image files
			log.Fatal("No image files found")
		}

		// Check if the imageNames.csv file exists
		if _, err := os.Stat("imageMetadata.csv"); os.IsNotExist(err) {
			log.Println("imageMetadata.csv does not exist, creating it...")

			// Get all Full image metadata
			log.Println("Getting full image metadata...")
			imagesMetaData := cmd.WalkImageFiles(fullsPath)

			// Save the image file names to the imageMetadata.csv file
			log.Println("Saving full image metadata to imageMetadata.csv...")
			cmd.SaveImageMetadata(imagesMetaData, "imageMetadata.csv")
		}
	}

	// Load the image file names from the imageNames.txt file
	fImageFiles := cmd.LoadImageMetadata("imageMetadata.csv")
	log.Printf("Image file count in imageMetadata.csv: %d\n", len(fImageFiles))

	if os.Getenv("SKIP_LOADING") == "false" {
		// Check if the image file count is different from the number of image file names
		if fCount != len(fImageFiles) {
			log.Println("Image file count is different from the number of image file names, updating imageMetadata.csv...")

			// Get all Full image metadata
			log.Println("Getting full image metadata...")
			imagesMetaData := cmd.WalkImageFiles(fullsPath)

			// Save the image file names to the imageMetadata.csv file
			log.Println("Saving full image metadata to imageMetadata.csv...")
			cmd.SaveImageMetadata(imagesMetaData, "imageMetadata.csv")
		}
	}
	// Set the image files in the thumbnails template
	model.Fulls = fImageFiles

	///////////////////////////////////////////////// THUMBS //////////////////////////////////////////////////
	// Check if the thumbs folder exists
	if _, err := os.Stat(thumbsPath); os.IsNotExist(err) {
		// Exit if the thumbs folder does not exist
		log.Fatal("Thumbs folder does not exist")
	}

	var tCount int
	if os.Getenv("SKIP_LOADING") == "false" {
		// Count the number of image files in the thumbs folder
		tCount = cmd.CountImagesInFolder(thumbsPath)
		log.Printf("Image file count in directory: %d\n", tCount)

		// Check if the image file count is zero
		if tCount == 0 {
			// Quit if there are no image files
			log.Fatal("No image files found")
		}

		// Check if the imageNames.csv file exists
		if _, err := os.Stat("imageNames.csv"); os.IsNotExist(err) {
			log.Println("imageNames.csv does not exist, creating it...")
			// Get the image files in the thumbs folder
			imageFiles := cmd.GetImageFilesInFolder(thumbsPath)

			// Save the image file names to the imageNames.txt file
			cmd.SaveImageNames(imageFiles)
		}
	}

	// Load the image file names from the imageNames.txt file
	tImageFiles := cmd.LoadImageNames("imageNames.csv")
	log.Printf("Image file count in imageNames.csv: %d\n", len(tImageFiles))

	if os.Getenv("SKIP_LOADING") == "false" {
		// Check if the image file count is different from the number of image file names
		if tCount != len(tImageFiles) {
			log.Println("Image file count is different from the number of image file names, updating imageNames.csv...")
			// Get the image files in the thumbs folder
			tImageFiles = cmd.GetImageFilesInFolder(thumbsPath)

			// Save the image file names to the imageNames.txt file
			cmd.SaveImageNames(tImageFiles)
		}
	}

	// Set the image files in the thumbnails template
	model.Thumbnails = tImageFiles

	///////////////////////////////////////////////// COMPARE //////////////////////////////////////////////////
	if os.Getenv("SKIP_LOADING") == "false" {
		// Check if the fCount and tCount are the same
		if fCount != tCount {
			log.Fatal("Fulls and Thumbs folder counts are different")
		} else {
			log.Println("Fulls and Thumbs folder counts are the same. Continue...")
		}
	}
}
