package main

import (
	"flag"
	"log"
	"os"

	"github.com/funayoseyoshito/yakiniku-image/lib"
	"github.com/funayoseyoshito/yakiniku-image/lib/db"
)

func main() {

	var f *flag.FlagSet
	var optionID int

	dbSet := db.GetDatabaseSet(
		lib.Config.Database.User,
		lib.Config.Database.Password,
		lib.Config.Database.Host,
		lib.Config.Database.Port,
		lib.Config.Database.Name)

	switch os.Args[1] {
	case lib.CmdInsert:
		log.Println("更新開始")
		f = flag.NewFlagSet(lib.CmdInsert, flag.ExitOnError)
		f.IntVar(&optionID, "store", 0, "insert store ID")
		f.Parse(os.Args[2:])
		//lib.InsertExecute(optionID, db)
		lib.InsertExecute(optionID, dbSet)
	case lib.CmdUpdate:
		log.Println("更新開始")
		f = flag.NewFlagSet(lib.CmdUpdate, flag.ExitOnError)
		f.IntVar(&optionID, "store", 0, "update store ID")
		f.Parse(os.Args[2:])

	case lib.CmdDelete:
		log.Println("削除開始")
		f = flag.NewFlagSet(lib.CmdDelete, flag.ExitOnError)
		f.IntVar(&optionID, "image", 0, "delete origin ID")
		f.Parse(os.Args[2:])

	default:
		lib.FatalExit("コマンドを確認してください")
	}

	log.Println("ツール終了")
}
