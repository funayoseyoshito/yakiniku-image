package lib

import (
	"fmt"

	"regexp"

	"path/filepath"

	"strconv"

	"log"

	"github.com/funayoseyoshito/yakiniku-image/lib/db"
)

var logo *LogoImages
var yakinikuAws *AwsIni

//var aws *AwsIni

func init() {

	logo = NewLogos()

	yakinikuAws = &AwsIni{
		AccessKeyID:     Config.Aws.GetAwsAccessKeyID(),
		SecretAccessKey: Config.Aws.GetAwsSecretAccessKey(),
		S3BucketName:    Config.Aws.GetAwsBucketName()}
	//panic("")
}

//InsertExecute is main
func InsertExecute(storeID int, dbSet *db.DatabaseSet) {

	fmt.Println(storeID)
	checkAndMakeDir(storeID)

	//SaveToS3()
	types := []string{TypeCookingName, TypeOtherName}
	for _, imageType := range types {

		fmt.Println(imageType)
		fmt.Println(Config.GetImageSrcPath(storeID, imageType))

		list := ReadDir(Config.GetImageSrcPath(storeID, imageType))
		for _, finfo := range list {

			r, _ := regexp.Compile(`^([0-9]+)\.(JPG|jpg|jpeg|JPEG)$`)
			fileInfoArray := r.FindAllStringSubmatch(finfo.Name(), -1)
			if len(fileInfoArray) != 1 || len(fileInfoArray[0]) != 3 {
				continue
			}
			var originId int
			var order int
			order, _ = strconv.Atoi(fileInfoArray[0][1])
			fmt.Println(finfo.Name(), order)

			//origin_nologo
			//TODO: Orderを登録する
			imageTable := &db.Images{
				StoreID: storeID,
				Kind:    Config.GetKindByKindNameAndTypeName(ImageOriginNoLogoName, imageType)}

			imageTable.Create(dbSet.Connection())
			originId = imageTable.ID
			originId = imageTable.ID
			fmt.Println(originId)
			originNoLogoImage := *GetImageByPath(filepath.Join(Config.GetImageSrcPath(storeID, imageType), finfo.Name()))
			yakinikuAws.SaveImageToS3(imageTable.ID, originNoLogoImage, SaveOriginImageQuality, S3ACLPrivate)

			//TODO: 一旦移動ロジックをコメントアウト
			//MoveProcessedImage(storeID, imageType, finfo.Name(), originId)
			log.Println(ImageOriginNoLogoName)
			//origin
			//TODO: Orderを登録する
			kind := Config.GetKindByKindNameAndTypeName(ImageOriginName, imageType)
			imageTable = &db.Images{
				StoreID:  storeID,
				Kind:     kind,
				OriginID: originId,
				//Order: orderId
			}

			imageTable.Create(dbSet.Connection())
			mixedImg := logo.LogoMixImageRGBA(kind, originNoLogoImage)
			yakinikuAws.SaveImageToS3(imageTable.ID, mixedImg, SaveNoOriginImageQuality, S3ACLPublic)
			SaveImageToFile(mixedImg, kind, imageTable.ID, storeID)
			log.Println(ImageOriginName)
			panic("===========================")
			//mixedImg := logo.LogoMixImageRGBA(Config.Cooking.OriginID, originImg)

			//var mixed Image image.Image = logo.LogoMixImageRGBA(Config.Cooking.OriginID, originImg)
			//fmt.Printf("%T")

			//Config.Cooking.MediumID
			//origin
			//origin logo
			//large
			//medium
			//small
			//micro
			//fmt.println(originimg)
		}
	}
}

//checkAndMakeDir 構成ディレクトリパスのチェックと作成
func checkAndMakeDir(storeID int) {
	//cooking
	if !CheckFilePathExists(Config.GetImageCookingOriginPath(storeID)) {
		CreateDir(Config.GetImageCookingOriginPath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageCookingOriginLogoPath(storeID)) {
		CreateDir(Config.GetImageCookingOriginLogoPath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageCookingLargePath(storeID)) {
		CreateDir(Config.GetImageCookingLargePath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageCookingMediumPath(storeID)) {
		CreateDir(Config.GetImageCookingMediumPath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageCookingSmallPath(storeID)) {
		CreateDir(Config.GetImageCookingSmallPath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageCookingMicroPath(storeID)) {
		CreateDir(Config.GetImageCookingMicroPath(storeID))
	}

	//other
	if !CheckFilePathExists(Config.GetImageOtherOriginPath(storeID)) {
		CreateDir(Config.GetImageOtherOriginPath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageOtherOriginLogoPath(storeID)) {
		CreateDir(Config.GetImageOtherOriginLogoPath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageOtherLargePath(storeID)) {
		CreateDir(Config.GetImageOtherLargePath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageOtherMediumPath(storeID)) {
		CreateDir(Config.GetImageOtherMediumPath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageOtherSmallPath(storeID)) {
		CreateDir(Config.GetImageOtherSmallPath(storeID))
	}
	if !CheckFilePathExists(Config.GetImageOtherMicroPath(storeID)) {
		CreateDir(Config.GetImageOtherMicroPath(storeID))
	}
}
