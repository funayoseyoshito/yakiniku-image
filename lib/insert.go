package lib

import (
	"fmt"

	"github.com/funayoseyoshito/yakiniku-image/lib/db"
)

func InsertExecute(storeID int, db *db.DatabaseSet) {
	fmt.Println(storeID)
	fmt.Println(db.Connection())
}
