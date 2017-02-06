package lib

import (
	"fmt"

	"regexp"

	"path/filepath"

	"strconv"

	"log"

	"github.com/funayoseyoshito/yakiniku-image/lib/db"
	"github.com/nfnt/resize"
)

//InsertExecute is main
func InsertExecute(storeID int, dbSet *db.DatabaseSet, awsSet *AwsIni) {
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

			var originId int
			var order int

			order, _ = strconv.Atoi(fileInfoArray[0][1])
			originNoLogoImage := *GetImageByPath(filepath.Join(Config.GetImageSrcPath(storeID, imageType), finfo.Name()))

			// -------------------
			//origin
			// -------------------
			kind := Config.GetKindByKindNameAndTypeName(ImageOriginName, imageType)
			mixedOriginImg := LogoMixImageRGBA(kind, originNoLogoImage, logo)
			imageTable := &db.Images{
				StoreID: storeID,
				Kind:    kind,
				Order:   order,
			}
			//insert
			imageTable.Create(dbSet.Connection())
			originId = imageTable.ID
			//update origin_id
			imageTable.OriginID = originId
			dbSet.Connection().Save(&imageTable)
			dbSet.Connection().Create(&db.ImageDescriptions{OriginID: originId})

			//save to s3
			awsSet.SaveImageToS3(originId, mixedOriginImg, SaveImageQuality50, S3ACLPublic)
			SaveImageToFile(mixedOriginImg, kind, imageTable.ID, storeID)
			log.Println(ImageOriginName, originId)

			// -------------------
			//origin_nologo
			// -------------------
			MoveProcessedImage(storeID, imageType, finfo.Name(), originId)
			log.Println(ImageOriginNoLogoName)

			// -------------------
			//Large
			// -------------------
			kind = Config.GetKindByKindNameAndTypeName(ImageLargeName, imageType)
			imageTable = &db.Images{
				StoreID:  storeID,
				Kind:     kind,
				OriginID: originId,
				Order:    order,
			}
			imageTable.Create(dbSet.Connection())
			mixedImage := resize.Thumbnail(uint(Config.ImageSize.LargeWidth), uint(Config.ImageSize.LargeHeight), mixedOriginImg, resize.Lanczos3)
			awsSet.SaveImageToS3(imageTable.ID, mixedImage, SaveImageQuality100, S3ACLPublic)
			SaveImageToFile(mixedImage, kind, imageTable.ID, storeID)
			log.Println(ImageLargeName, imageTable.ID)

			// -------------------
			//Medium
			// -------------------
			kind = Config.GetKindByKindNameAndTypeName(ImageMediumName, imageType)
			imageTable = &db.Images{
				StoreID:  storeID,
				Kind:     kind,
				OriginID: originId,
				Order:    order,
			}
			imageTable.Create(dbSet.Connection())
			mixedImage = resize.Thumbnail(uint(Config.ImageSize.MediumWidth), uint(Config.ImageSize.MediumHeight), mixedOriginImg, resize.Lanczos3)
			awsSet.SaveImageToS3(imageTable.ID, mixedImage, SaveImageQuality100, S3ACLPublic)
			SaveImageToFile(mixedImage, kind, imageTable.ID, storeID)
			log.Println(ImageMediumName, imageTable.ID)

			// -------------------
			//Small
			// -------------------
			kind = Config.GetKindByKindNameAndTypeName(ImageSmallName, imageType)
			imageTable = &db.Images{
				StoreID:  storeID,
				Kind:     kind,
				OriginID: originId,
				Order:    order,
			}
			imageTable.Create(dbSet.Connection())
			mixedImage = resize.Thumbnail(uint(Config.ImageSize.SmallWidth), uint(Config.ImageSize.SmallHeight), originNoLogoImage, resize.Lanczos3)
			awsSet.SaveImageToS3(imageTable.ID, mixedImage, SaveImageQuality100, S3ACLPublic)
			SaveImageToFile(mixedImage, kind, imageTable.ID, storeID)
			log.Println(ImageSmallName, imageTable.ID)

			// -------------------
			//micro
			// -------------------
			kind = Config.GetKindByKindNameAndTypeName(ImageMicroName, imageType)
			imageTable = &db.Images{
				StoreID:  storeID,
				Kind:     kind,
				OriginID: originId,
				Order:    order,
			}
			imageTable.Create(dbSet.Connection())
			mixedImage = resize.Thumbnail(uint(Config.ImageSize.MicroWidth), uint(Config.ImageSize.MicroHeight), originNoLogoImage, resize.Lanczos3)
			awsSet.SaveImageToS3(imageTable.ID, mixedImage, SaveImageQuality100, S3ACLPublic)
			SaveImageToFile(mixedImage, kind, imageTable.ID, storeID)
			log.Println(ImageMicroName, imageTable.ID)
		}
	}
}

////checkAndMakeDir 構成ディレクトリパスのチェックと作成
//func CheckAndMakeDir(storeID int) {
//	//cooking
//	if !CheckFilePathExists(Config.GetImageCookingOriginPath(storeID)) {
//		CreateDir(Config.GetImageCookingOriginPath(storeID))
//	}
//	if !CheckFilePathExists(Config.GetImageCookingOriginLogoPath(storeID)) {
//		CreateDir(Config.GetImageCookingOriginLogoPath(storeID))
//	}
//	if !CheckFilePathExists(Config.GetImageCookingLargePath(storeID)) {
//		CreateDir(Config.GetImageCookingLargePath(storeID))
//	}
//	if !CheckFilePathExists(Config.GetImageCookingMediumPath(storeID)) {
//		CreateDir(Config.GetImageCookingMediumPath(storeID))
//	}
//	if !CheckFilePathExists(Config.GetImageCookingSmallPath(storeID)) {
//		CreateDir(Config.GetImageCookingSmallPath(storeID))
//	}
//	if !CheckFilePathExists(Config.GetImageCookingMicroPath(storeID)) {
//		CreateDir(Config.GetImageCookingMicroPath(storeID))
//	}
//
//	//other
//	if !CheckFilePathExists(Config.GetImageOtherOriginPath(storeID)) {
//		CreateDir(Config.GetImageOtherOriginPath(storeID))
//	}
//	if !CheckFilePathExists(Config.GetImageOtherOriginLogoPath(storeID)) {
//		CreateDir(Config.GetImageOtherOriginLogoPath(storeID))
//	}
//	if !CheckFilePathExists(Config.GetImageOtherLargePath(storeID)) {
//		CreateDir(Config.GetImageOtherLargePath(storeID))
//	}
//	if !CheckFilePathExists(Config.GetImageOtherMediumPath(storeID)) {
//		CreateDir(Config.GetImageOtherMediumPath(storeID))
//	}
//	if !CheckFilePathExists(Config.GetImageOtherSmallPath(storeID)) {
//		CreateDir(Config.GetImageOtherSmallPath(storeID))
//	}
//	if !CheckFilePathExists(Config.GetImageOtherMicroPath(storeID)) {
//		CreateDir(Config.GetImageOtherMicroPath(storeID))
//	}
//}
