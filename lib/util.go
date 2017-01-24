package lib

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

//FatalExit is output Fatal error and exit program
func FatalExit(msg interface{}) {
	log.Fatal(msg)
	os.Exit(2)
}

//CheckFilePathExists is exits path
func CheckFilePathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

//CreateDir make dir
func CreateDir(path string) {
	if err := os.MkdirAll(path, 0777); err != nil {
		FatalExit(err)
	}
}

//ReadDir return os.FileInfo
func ReadDir(path string) []os.FileInfo {
	list, err := ioutil.ReadDir(path)
	if err != nil {
		FatalExit(err)
	}
	return list
}

//MoveProcessed
func MoveProcessedImage(storeID int, imageType string, fileName string, dbID int) {
	FileMove(filepath.Join(Config.GetImageSrcPath(storeID, imageType), fileName),
		filepath.Join(Config.GetImageOriginNoLogoPath(storeID, imageType), strconv.Itoa(dbID)+"."+SaveImageExt))
}

//FileMove is move ファイル
func FileMove(source string, target string) {
	if err := os.Rename(source, target); err != nil {
		FatalExit(err)
	}
}

//checkAndMakeDir 構成ディレクトリパスのチェックと作成
func CheckAndMakeDir(storeID int) {
	//cooking
	if !CheckFilePathExists(Config.GetImageCookingOriginPath(storeID)) {
		CreateDir(Config.GetImageCookingOriginPath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageCookingOriginLogoPath(storeID)) {
		CreateDir(Config.GetImageCookingOriginLogoPath(storeID))
		if !CheckFilePathExists(Config.GetImageCookingLargePath(storeID)) {
			CreateDir(Config.GetImageCookingLargePath(storeID))
		}
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
