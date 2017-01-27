package lib

import (
	"fmt"

	"log"

	"github.com/funayoseyoshito/yakiniku-image/lib/db"
)

func DeleteExecute(imageID int, dbSet *db.DatabaseSet, awsSet *AwsIni) {
	defer dbSet.Connection().Close()

	var deleteIDs []int
	var images []*db.Images

	dbSet.Connection().Where("id = ? or origin_id = ?", imageID, imageID).Find(&images)
	for _, v := range images {
		deleteIDs = append(deleteIDs, v.ID)
	}

	fmt.Println("対象key:", deleteIDs)

	log.Println("DB==============================")
	if len(deleteIDs) > 0 {
		dbSet.Connection().Where("id in (?)", deleteIDs).Delete(db.Images{})
	}
	log.Print("DB 削除完了")

	log.Println("S3==============================")
	awsSet.DeleteS3Images(deleteIDs)
	log.Print("DB 削除完了")
}
