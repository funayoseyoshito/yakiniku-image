package lib

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
)

//LogoImages LogoImage is logo image pointer
type LogoImages struct {
	OriginLogo *image.Image
	LargeLogo  *image.Image
	MediumLogo *image.Image
}

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

//SrcImage is source image struct
type SrcImage struct {
	kind int
}

//NewLogos return LogoImages struct pointer
func NewLogos() *LogoImages {
	return &LogoImages{
		OriginLogo: GetImageByPath(Config.GetOriginLogoPath()),
		LargeLogo:  GetImageByPath(Config.GetLargeLogoPath()),
		MediumLogo: GetImageByPath(Config.GetMediumLogoPath())}
}

//GetLgoImageByKind return logo pointer
func (logo *LogoImages) GetLogoImageByKind(kind int) *image.Image {
	var l *image.Image
	switch kind {
	case Config.Cooking.OriginID, Config.Other.OriginID:
		l = logo.OriginLogo
	case Config.Cooking.LargeID, Config.Other.LargeID:
		l = logo.LargeLogo
	case Config.Cooking.MediumID, Config.Other.MediumID:
		l = logo.MediumLogo
	default:
		FatalExit("logo kind が一致しませんでした。")
	}
	return l
}

func (logo *LogoImages) LogoMixImageRGBA(kind int, originImg image.Image) image.Image {
	o := originImg
	l := *logo.GetLogoImageByKind(kind)
	startPointLogo := image.Point{o.Bounds().Dx() - l.Bounds().Dx() - 50, o.Bounds().Dy() - l.Bounds().Dy() - 50}
	logoRectangle := image.Rectangle{startPointLogo, startPointLogo.Add(l.Bounds().Size())}
	originRectangle := image.Rectangle{image.Point{0, 0}, o.Bounds().Size()}

	rgba := image.NewRGBA(originRectangle)
	draw.Draw(rgba, originRectangle, o, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, logoRectangle, l, image.Point{0, 0}, draw.Over)

	return rgba
}

////GetLogoMixImageRGBA 合成画像を生成する
//func GetLogoMixImageRGBA(originImg image.Image, logoImg image.Image) *image.Image {
//	startPointLogo := image.Point{
//		originImg.Bounds().Dx() - logoImg.Bounds().Dx() - 10, originImg.Bounds().Dy() - logoImg.Bounds().Dy() - 10}
//
//	logoRectangle := image.Rectangle{startPointLogo, startPointLogo.Add(logoImg.Bounds().Size())}
//	originRectangle := image.Rectangle{image.Point{0, 0}, originImg.Bounds().Size()}
//
//	rgba := image.NewRGBA(originRectangle)
//	draw.Draw(rgba, originRectangle, originImg, image.Point{0, 0}, draw.Src)
//	draw.Draw(rgba, logoRectangle, logoImg, image.Point{0, 0}, draw.Over)
//
//	return &rgba
//}

//GetImageByPath ファイルパスからロゴimageを生成
func GetImageByPath(path string) *image.Image {
	logoFile, err := os.Open(path)

	if err != nil {
		FatalExit(err)
	}
	logoImg, _, err := image.Decode(logoFile)
	if err != nil {
		FatalExit(err)
	}

	return &logoImg
}

//SaveImageToFile is save image except origi no logo image
func SaveImageToFile(img image.Image, imageKind int, insertID int, storeID int) {

	out, err := os.Create(filepath.Join(
		Config.GetStoreImagePath(storeID), Config.GetImageTypeByKind(imageKind),
		Config.GetImageKindNameByKind(imageKind), strconv.Itoa(insertID)+"."+SaveImageExt))
	if err != nil {
		fmt.Println(err)
	}

	var opt jpeg.Options
	opt.Quality = SaveImageQuality50
	jpeg.Encode(out, img, &opt)
}
