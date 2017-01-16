package db

import "time"

const SelectLimit = 10

//Images テーブル定義
type Images struct {
	ID          int
	StoreID     int
	OriginID    int
	Kind        int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

////GetOriginImage オリジナル画像のimage.Imageを取得する
//func (img Images) GetOriginImage() image.Image {
//	originImg, _, err := image.Decode(bytes.NewReader(img.Source))
//	if err != nil {
//		fmt.Println(err)
//	}
//	return originImg
//}
//
////UpdateImage 合成画像で旧レコードを更新するメソッド
//func (img *Images) UpdateImage(imgRGBA image.Image) {
//
//	var opt jpeg.Options
//	opt.Quality = 100
//	buffer := new(bytes.Buffer)
//
//	if err := jpeg.Encode(buffer, imgRGBA, &opt); err != nil {
//		log.Println("unable to encode image.")
//	}
//	imageBytes := buffer.Bytes()
//
//	img.Source = imageBytes
//
//	//GetConnection().Save(img)
//}

// InsertDb テスト用の関数
//func InsertDb(img image.Image) int {
//
//	var opt jpeg.Options
//	opt.Quality = 100
//	buffer := new(bytes.Buffer)
//
//	if err := jpeg.Encode(buffer, img, &opt); err != nil {
//		log.Println("unable to encode image.")
//	}
//
//	imageBytes := buffer.Bytes()
//	row := Images{StoreID: 1111, Kind: 1111, Source: imageBytes, OriginID: 1111}
//	GetConnection().Create(&row)
//
//	return row.ID
//}

//SelectProcessingRows 処理対象のレコードを定数ずつ取得する
//func SelectProcessingRows(offset int) *sql.Rows {
//	rows, err := GetConnection().
//		Model(&Images{}).
//		Not("kind", []int{
//			lib.Config.Cooking.MicroID, lib.Config.Cooking.SmallID,
//			lib.Config.Other.MicroID, lib.Config.Other.SmallID}).
//		Offset(offset).Limit(SelectLimit).Rows()
//
//	if err != nil {
//		fmt.Println(err)
//		panic("オリジナル画像の取得に失敗")
//	}
//	return rows
//}
