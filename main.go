package main

import (
	"flag"
	"log"
	"os"

	"github.com/funayoseyoshito/yakiniku-image/lib"
	"github.com/funayoseyoshito/yakiniku-image/lib/db"
)

var optionID int

func main() {

	var f *flag.FlagSet

	dbSet := db.NewDatabaseSet(
		lib.Config.Database.User,
		lib.Config.Database.Password,
		lib.Config.Database.Host,
		lib.Config.Database.Port,
		lib.Config.Database.Name)

	if len(os.Args) < 1 {
		lib.FatalExit("コマンドを確認してください")
	}

	switch os.Args[1] {
	case lib.CmdInsert:
		log.Println("更新開始")
		f = flag.NewFlagSet(lib.CmdInsert, flag.ExitOnError)
		f.IntVar(&optionID, "store", 0, lib.CmdOptionStore+" store id")
		f.Parse(os.Args[2:])
		checkOptionEmpty(lib.CmdOptionStore)
		lib.InsertExecute(optionID, dbSet)

	case lib.CmdUpdate:
		log.Println("更新開始")
		f = flag.NewFlagSet(lib.CmdUpdate, flag.ExitOnError)
		f.IntVar(&optionID, "store", 0, "update store ID")
		f.Parse(os.Args[2:])
		checkOptionEmpty(lib.CmdOptionImage)

	case lib.CmdDelete:
		log.Println("削除開始")
		f = flag.NewFlagSet(lib.CmdDelete, flag.ExitOnError)
		f.IntVar(&optionID, "image", 0, "delete origin ID")
		f.Parse(os.Args[2:])
		checkOptionEmpty(lib.CmdOptionStore)

	default:
		lib.FatalExit("コマンドを確認してください")
	}

	log.Println("ツール終了")
}

//checkOptionEmpty オプションが正しく渡されているかチェックする
func checkOptionEmpty(o string) {
	if optionID == 0 {
		lib.FatalExit(o + "オプションを正しく指定してください。")
	}
}
