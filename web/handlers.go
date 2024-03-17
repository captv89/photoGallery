package web

import (
	"time"

	"github.com/captv89/photoGallery/cmd"
	tf "github.com/captv89/photoGallery/web/tfs"

	"github.com/labstack/echo/v4"
)

func home(e echo.Context) error {
	home := tf.Home()
	return Render(e, 200, home)
}

func imgModal(e echo.Context) error {
	imgName := e.Param("id")
	metaData := cmd.GetImageMetaData(imgName)
	previous, next := cmd.GetPreviousAndNextImage(imgName)
	// log.Printf("Image Metadata: %v\n", metaData)
	imgModal := tf.ModalWrapperWithImage(imgName, previous, next, metaData)
	// sleep for 2 seconds
	time.Sleep(1 * time.Second)
	return Render(e, 200, imgModal)
}
