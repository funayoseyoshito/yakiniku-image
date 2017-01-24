package lib

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"

	"regexp"

	"github.com/funayoseyoshito/yakiniku-image/lib/db"
	"github.com/nfnt/resize"
)

//InsertExecute is main
func UpdateExecute(storeID int, dbSet *db.DatabaseSet, awsSet *AwsIni) {
	fmt.Printf("[店舗ID] %d\n", storeID)
	defer dbSet.Connection().Close()
	logo := GetLogo()

	CheckAndMakeDir(storeID)

	//SaveToS3()
	types := []string{TypeCookingName, TypeOtherName}
	for _, imageType := range types {

		list := ReadDir(Config.GetImageSrcPath(storeID, imageType))

		log.Println("==============================")
		fmt.Printf("[%s] %s\n", imageType, Config.GetImageSrcPath(storeID, imageType))

		for _, finfo := range list {
			r, _ := regexp.Compile(`^([0-9]+)\.(JPG|jpg|jpeg|JPEG)$`)
			fileInfoArray := r.FindAllStringSubmatch(finfo.Name(), -1)
			if len(fileInfoArray) != 1 || len(fileInfoArray[0]) != 3 {
				continue
			}

			log.Println("------------")
			fmt.Println(finfo.Name())

			var originID int
			originID, _ = strconv.Atoi(fileInfoArray[0][1])

			rows, _ := dbSet.Connection().Model(&db.Images{}).Where("id = ? or origin_id = ?", originID, originID).Rows()
			var imageRowMap map[int]db.Images = make(map[int]db.Images)
			var imageRow db.Images
			for rows.Next() {
				dbSet.Connection().ScanRows(rows, &imageRow)

				imageRowMap[imageRow.Kind] = imageRow
			}

			if len(imageRowMap) != 5 {
				FatalExit("imagesテーブルに種類ごとのレコードが存在するか、確認してください。")
			}

			originNoLogoImage := *GetImageByPath(filepath.Join(Config.GetImageSrcPath(storeID, imageType), finfo.Name()))

			// -------------------
			//origin
			// -------------------
			kind := Config.GetKindByKindNameAndTypeName(ImageOriginName, imageType)
			dbRow := imageRowMap[kind]
			mixedOriginImg := LogoMixImageRGBA(kind, originNoLogoImage, logo)
			awsSet.SaveImageToS3(dbRow.ID, mixedOriginImg, SaveImageQuality50, S3ACLPublic)
			SaveImageToFile(mixedOriginImg, kind, dbRow.ID, storeID)
			dbSet.Connection().Save(&dbRow)
			log.Println(ImageOriginName, dbRow.ID)

			// -------------------
			//origin_nologo
			// -------------------
			MoveProcessedImage(storeID, imageType, finfo.Name(), originID)
			log.Println(ImageOriginNoLogoName)

			// -------------------
			//Large
			// -------------------
			kind = Config.GetKindByKindNameAndTypeName(ImageLargeName, imageType)
			dbRow = imageRowMap[kind]
			mixedImage := resize.Thumbnail(uint(Config.ImageSize.LargeWidth), uint(Config.ImageSize.LargeHeight), mixedOriginImg, resize.Lanczos3)
			awsSet.SaveImageToS3(dbRow.ID, mixedImage, SaveImageQuality100, S3ACLPublic)
			SaveImageToFile(mixedImage, kind, dbRow.ID, storeID)
			dbSet.Connection().Save(&dbRow)
			log.Println(ImageLargeName, dbRow.ID)

			// -------------------
			//Medium
			// -------------------
			kind = Config.GetKindByKindNameAndTypeName(ImageMediumName, imageType)
			dbRow = imageRowMap[kind]
			mixedImage = resize.Thumbnail(uint(Config.ImageSize.MediumWidth), uint(Config.ImageSize.MediumHeight), mixedOriginImg, resize.Lanczos3)
			awsSet.SaveImageToS3(dbRow.ID, mixedImage, SaveImageQuality100, S3ACLPublic)
			SaveImageToFile(mixedImage, kind, dbRow.ID, storeID)
			dbSet.Connection().Save(&dbRow)
			log.Println(ImageMediumName, dbRow.ID)

			// -------------------
			//Small
			// -------------------
			kind = Config.GetKindByKindNameAndTypeName(ImageSmallName, imageType)
			dbRow = imageRowMap[kind]
			resizeImage := resize.Thumbnail(uint(Config.ImageSize.SmallWidth), uint(Config.ImageSize.SmallHeight), originNoLogoImage, resize.Lanczos3)
			awsSet.SaveImageToS3(dbRow.ID, resizeImage, SaveImageQuality100, S3ACLPublic)
			SaveImageToFile(resizeImage, kind, dbRow.ID, storeID)
			dbSet.Connection().Save(&dbRow)
			log.Println(ImageSmallName, dbRow.ID)

			// -------------------
			//micro
			// -------------------
			kind = Config.GetKindByKindNameAndTypeName(ImageMicroName, imageType)
			dbRow = imageRowMap[kind]
			resizeImage = resize.Thumbnail(uint(Config.ImageSize.MicroWidth), uint(Config.ImageSize.MicroHeight), originNoLogoImage, resize.Lanczos3)
			awsSet.SaveImageToS3(dbRow.ID, resizeImage, SaveImageQuality100, S3ACLPublic)
			SaveImageToFile(resizeImage, kind, dbRow.ID, storeID)
			dbSet.Connection().Save(&dbRow)
			log.Println(ImageMicroName, dbRow.ID)
		}
	}
}
