package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/funayoseyoshito/yakiniku-image/lib"
)

var (
	//StoreID 処理対象店舗ID
	StoreID int
	//ImgID 処理対象画像ID
	ImgID int
	//CMD 処理中のコマンド
	CMD string
)

func init() {

	var f *flag.FlagSet

	switch CMD = os.Args[1]; CMD {
	case lib.CmdInsert:
		log.Println("更新開始")
		f = flag.NewFlagSet(lib.CmdInsert, flag.ExitOnError)
		f.IntVar(&StoreID, "store", 0, "insert store ID")

	case lib.CmdUpdate:
		log.Println("更新開始")
		f = flag.NewFlagSet(lib.CmdUpdate, flag.ExitOnError)
		f.IntVar(&StoreID, "store", 0, "update store ID")

	case lib.CmdDelete:
		log.Println("削除開始")
		f = flag.NewFlagSet(lib.CmdDelete, flag.ExitOnError)
		f.IntVar(&ImgID, "image", 0, "delete origin ID")

	default:
		lib.FatalExit("コマンドを確認してください")
	}

	f.Parse(os.Args[2:])

	fmt.Println("opt1:", StoreID)
	fmt.Println("opt1:", ImgID)
}

func main() {
}
