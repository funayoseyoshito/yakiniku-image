package lib

import (
	"fmt"
	"path/filepath"

	"regexp"

	"github.com/funayoseyoshito/yakiniku-image/lib/db"
)

//Insert Execute 処理メイン
func InsertExecute(storeID int, db *db.DatabaseSet) {
	fmt.Println(storeID)
	checkAndMakeDir(storeID)

	types := []string{TypeCookingName, TypeOtherName}

	for _, imageType := range types {

		fmt.Println(imageType)
		fmt.Println(Config.GetImageSrcPath(storeID, imageType))

		list := ReadDir(Config.GetImageSrcPath(storeID, imageType))
		for _, finfo := range list {

			r, _ := regexp.Compile(`^([0-9]+)\.(JPG|jpg|jpeg|JPEG)$`)
			fileInfoArray := r.FindAllStringSubmatch(finfo.Name(), -1)
			if len(fileInfoArray) != 1 || len(fileInfoArray[0]) != 3 {
				//fmt.Println(fileInfoArray)
				continue
			}
			order := fileInfoArray[0][1]
			fmt.Println(finfo.Name(), order)
			originImg := GetOriginImage(filepath.Join(
				Config.GetImageSrcPath(storeID, imageType), finfo.Name()))
			//origin
			//origin logo
			//large
			//medium
			//small
			//micro
			fmt.Println(originImg)
		}

		//
		//list := ReadDir(Config.GetImageCookingSrcPath(storeID))
	}

	//fmt.Println(db.Connection())
}

//checkAndMakeDir 構成ディレクトリパスのチェックと作成
func checkAndMakeDir(storeID int) {
	//cooking
	if !CheckFilePathExists(Config.GetImageCookingOriginPath(storeID)) {
		CreateDir(Config.GetImageCookingOriginPath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageCookingOriginLogoPath(storeID)) {
		CreateDir(Config.GetImageCookingOriginLogoPath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageCookingLargePath(storeID)) {
		CreateDir(Config.GetImageCookingLargePath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageCookingMediumPath(storeID)) {
		CreateDir(Config.GetImageCookingMediumPath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageCookingSmallPath(storeID)) {
		CreateDir(Config.GetImageCookingSmallPath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageCookingMicroPath(storeID)) {
		CreateDir(Config.GetImageCookingMicroPath(storeID))
	}

	//other
	if !CheckFilePathExists(Config.GetImageOtherOriginPath(storeID)) {
		CreateDir(Config.GetImageOtherOriginPath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageOtherOriginLogoPath(storeID)) {
		CreateDir(Config.GetImageOtherOriginLogoPath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageOtherLargePath(storeID)) {
		CreateDir(Config.GetImageOtherLargePath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageOtherMediumPath(storeID)) {
		CreateDir(Config.GetImageOtherMediumPath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageOtherSmallPath(storeID)) {
		CreateDir(Config.GetImageOtherSmallPath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageOtherMicroPath(storeID)) {
		CreateDir(Config.GetImageOtherMicroPath(storeID))
	}
}
