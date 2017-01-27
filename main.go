package main

import (
	"flag"
	"log"
	"os"

	"fmt"

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

	awsSet := lib.NewAws()

	if len(os.Args) < 1 {
		lib.FatalExit("コマンドを確認してください")
	}

	fmt.Println("#####################################################################")

	switch os.Args[1] {
	case lib.CmdInsert:
		log.Println("INSERT 開始")
		f = flag.NewFlagSet(lib.CmdInsert, flag.ExitOnError)
		f.IntVar(&optionID, "store", 0, lib.CmdOptionStore+" store ID")
		f.Parse(os.Args[2:])
		checkOptionEmpty(lib.CmdOptionStore)
		lib.InsertExecute(optionID, dbSet, awsSet)

	case lib.CmdUpdate:
		log.Println("UPDATE 開始")
		f = flag.NewFlagSet(lib.CmdUpdate, flag.ExitOnError)
		f.IntVar(&optionID, "store", 0, lib.CmdUpdate+" store ID")
		f.Parse(os.Args[2:])
		checkOptionEmpty(lib.CmdOptionImage)
		lib.UpdateExecute(optionID, dbSet, awsSet)

	case lib.CmdDelete:
		log.Println("DELETE 開始")
		f = flag.NewFlagSet(lib.CmdDelete, flag.ExitOnError)
		f.IntVar(&optionID, "image", 0, "delete origin ID")
		f.Parse(os.Args[2:])
		checkOptionEmpty(lib.CmdOptionStore)
		lib.DeleteExecute(optionID, dbSet, awsSet)

	default:
		lib.FatalExit("コマンドを確認してください")
	}

	log.Println("終了")
	fmt.Println("#####################################################################")
}

//checkOptionEmpty オプションが正しく渡されているかチェックする
func checkOptionEmpty(o string) {
	if optionID == 0 {
		lib.FatalExit(o + "オプションを正しく指定してください。")
	}
}
