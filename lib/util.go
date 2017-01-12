package lib

import (
	"log"
	"os"
)

func FatalExit(msg interface{}) {
	log.Fatal(msg)
	os.Exit(2)
}
