package lib

import (
	"log"
	"os"
)

func FatalExit(msg string) {
	log.Fatal(msg)
	os.Exit(2)
}
