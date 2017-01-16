package lib

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

//GetOriginImage オリジナル画像のimage.Imageを取得する
func GetOriginImage(path string) *image.Image {

	// オリジナル画像open
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		FatalExit(err)
	}
	return &img
}
