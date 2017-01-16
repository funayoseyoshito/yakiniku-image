package lib

import (
	"os"
	"path/filepath"

	"strconv"

	"github.com/BurntSushi/toml"
)

const (
	//CmdInsert insert mode 定数
	CmdInsert = "insert"
	//CmdUpdate update mode 定数
	CmdUpdate = "update"
	//CmdDelete delete mode 定数
	CmdDelete = "delete"
	//CmdOptionStore -store オプション
	CmdOptionStore = "store"
	//CmdOptionImage -image オプション
	CmdOptionImage = "image"
	//AssetsDirName assets dir name
	AssetsDirName = "assets"
	//ImageDirName dir name
	ImageDirName = "image"
	//LogoDirName logo dir name
	LogoDirName = "logo"
	//ImageMicroName micro dir name
	ImageMicroName = "micro"
	//ImageSmallName small dir name
	ImageSmallName = "small"
	//ImageMediumName medium dir name
	ImageMediumName = "medium"
	//ImageLargeName large dir name
	ImageLargeName = "large"
	//ImageOriginName origin dir name
	ImageOriginName = "origin"
	//ImageOriginLogoName origin logo dir name
	ImageOriginLogoName = "origin_logo"
	//TypeCookingName cooking type name
	TypeCookingName = "cooking"
	//TypeOtherName eother type name
	TypeOtherName = "other"
	//SaveImageExt save image extension
	SaveImageExt = "jpg"
)

var (
	//Config 設定ファイルのグローバル変数
	Config Configs
	//CurrentBasePath current path
	CurrentBasePath string
)

//Configs 設定ファイル
type Configs struct {
	Database DatabaseConfig
	Logo     LogoConfig
	Cooking  CookingConfig
	Other    OtherConfig
}

//DatabaseConfig database 設定ファイル
type DatabaseConfig struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
}

//LogoConfig logo画像の設定
type LogoConfig struct {
	MediumName string
	LargeName  string
	OriginName string
}

//CookingConfig 料理画像の設定
type CookingConfig struct {
	OriginID int
	LargeID  int
	MediumID int
	SmallID  int
	MicroID  int
}

//OtherConfig その他画像の設定
type OtherConfig struct {
	OriginID int
	LargeID  int
	MediumID int
	SmallID  int
	MicroID  int
}

//GetAssetsPath assets path
func (con *Configs) GetAssetsPath() string {
	return filepath.Join(CurrentBasePath, AssetsDirName)
}

//GetLogoPath logo path
func (con *Configs) GetLogoPath() string {
	return filepath.Join(con.GetAssetsPath(), LogoDirName)
}

//GetMediumLogoPath mediumLogo path
func (con *Configs) GetMediumLogoPath() string {
	return filepath.Join(con.GetLargeLogoPath(), con.Logo.MediumName)
}

//GetLargeLogoPath largeLogo path
func (con *Configs) GetLargeLogoPath() string {
	return filepath.Join(con.GetLargeLogoPath(), con.Logo.LargeName)
}

//GetOriginLogoPath originLogo path
func (con *Configs) GetOriginLogoPath() string {
	return filepath.Join(con.GetLargeLogoPath(), con.Logo.OriginName)
}

//GetImagePath image dir path
func (con *Configs) GetImagePath() string {
	return filepath.Join(con.GetAssetsPath(), ImageDirName)
}

//GetStoreImagePath store image path
func (con *Configs) GetStoreImagePath(storeID int) string {
	return filepath.Join(con.GetImagePath(), strconv.Itoa(storeID))
}

//GetImageMicroPath micro image path
func (con *Configs) GetImageMicroPath(storeID int, typeName string) string {
	return filepath.Join(con.GetStoreImagePath(storeID), typeName, ImageMicroName)
}

//GetImageSmallPath small image path
func (con *Configs) GetImageSmallPath(storeID int, typeName string) string {
	return filepath.Join(con.GetStoreImagePath(storeID), typeName, ImageSmallName)
}

//GetImageMediumPath medium image path
func (con *Configs) GetImageMediumPath(storeID int, typeName string) string {
	return filepath.Join(con.GetStoreImagePath(storeID), typeName, ImageMediumName)
}

//GetImageLargePath large image path
func (con *Configs) GetImageLargePath(storeID int, typeName string) string {
	return filepath.Join(con.GetStoreImagePath(storeID), typeName, ImageLargeName)
}

//GetImagePath origin image path
func (con *Configs) GetImageOriginPath(storeID int, typeName string) string {
	return filepath.Join(con.GetStoreImagePath(storeID), typeName, ImageOriginName)
}

//GetImageOriginLogoPath originLogo image path
func (con *Configs) GetImageOriginLogoPath(storeID int, typeName string) string {
	return filepath.Join(con.GetStoreImagePath(storeID), typeName, ImageOriginLogoName)
}

//GetImageSrcPath cooking image source dir
func (con *Configs) GetImageSrcPath(storeID int, typeName string) string {
	return filepath.Join(con.GetStoreImagePath(storeID), typeName)
}

//GetImageCookingSrcPath cooking image source dir
func (con *Configs) GetImageCookingSrcPath(storeID int) string {
	return filepath.Join(con.GetStoreImagePath(storeID), TypeCookingName)
}

//GetImageCookingMicroPath cooking micro image path
func (con *Configs) GetImageCookingMicroPath(storeID int) string {
	return con.GetImageMicroPath(storeID, TypeCookingName)
}

//GetImageCookingSmallPath cooiking small image path
func (con *Configs) GetImageCookingSmallPath(storeID int) string {
	return con.GetImageSmallPath(storeID, TypeCookingName)
}

//GetImageCookingMediumPath cooking medium image path
func (con *Configs) GetImageCookingMediumPath(storeID int) string {
	return con.GetImageMediumPath(storeID, TypeCookingName)
}

//GetImageCookingLargePath cooking large image pa†h
func (con *Configs) GetImageCookingLargePath(storeID int) string {
	return con.GetImageLargePath(storeID, TypeCookingName)
}

//GetImageCookingOriginPath cooking origin image path
func (con *Configs) GetImageCookingOriginPath(storeID int) string {
	return con.GetImageOriginPath(storeID, TypeCookingName)
}

//GetImageCookingOriginLogoPath cooking origin logo path
func (con *Configs) GetImageCookingOriginLogoPath(storeID int) string {
	return con.GetImageOriginLogoPath(storeID, TypeCookingName)
}

//GetImageOtherSrcPath other image source dir
func (con *Configs) GetImageOtherSrcPath(storeID int) string {
	return filepath.Join(con.GetStoreImagePath(storeID), TypeCookingName)
}

//GetImageOtherMicroPath other micro image path
func (con *Configs) GetImageOtherMicroPath(storeID int) string {
	return con.GetImageMicroPath(storeID, TypeOtherName)
}

//GetImageOtherSmallPath other small image path
func (con *Configs) GetImageOtherSmallPath(storeID int) string {
	return con.GetImageSmallPath(storeID, TypeOtherName)
}

//GetImageOtherMediumPath other medium image path
func (con *Configs) GetImageOtherMediumPath(storeID int) string {
	return con.GetImageMediumPath(storeID, TypeOtherName)
}

//GetImageOtherLargePath other large image pa†h
func (con *Configs) GetImageOtherLargePath(storeID int) string {
	return con.GetImageLargePath(storeID, TypeOtherName)
}

//GetImageOtherOriginPath other origin image path
func (con *Configs) GetImageOtherOriginPath(storeID int) string {
	return con.GetImageOriginPath(storeID, TypeOtherName)
}

//GetImageOtherOriginLogoPath other origin logo path
func (con *Configs) GetImageOtherOriginLogoPath(storeID int) string {
	return con.GetImageOriginLogoPath(storeID, TypeOtherName)
}

func init() {
	_, err := toml.DecodeFile("./config.toml", &Config)
	if err != nil {
		FatalExit(err)
	}
	basePath, _ := os.Getwd()
	CurrentBasePath = basePath
}
