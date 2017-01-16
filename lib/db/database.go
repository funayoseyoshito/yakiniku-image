package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//gormで依存がある為

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type DatabaseSet struct {
	User, Password, Host, Port, Name string
}

func NewDatabaseSet(user string, password string, host string, port string, name string) *DatabaseSet {
	return &DatabaseSet{
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
		Name:     name}
}

//GetConnection データベースコネクションを返却します
func (d *DatabaseSet) Connection() *gorm.DB {
	if db == nil {
		databaseInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
			d.User,
			d.Password,
			d.Host,
			d.Port,
			d.Name)

		var err error
		db, err = gorm.Open("mysql", databaseInfo)
		if err != nil {
			fmt.Println(err)
			panic("failed to connect database")
		}
	}
	return db
}
