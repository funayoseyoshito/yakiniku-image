package lib

import (
	"io/ioutil"
	"log"
	"os"
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
