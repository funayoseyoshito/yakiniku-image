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
