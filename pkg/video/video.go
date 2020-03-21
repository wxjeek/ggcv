package video

import (
	"gocv.io/x/gocv"
)

func Video()  {
	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()
	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
