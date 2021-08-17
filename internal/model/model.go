package model

import (
	"fmt"

	"github.com/goprogramming/blog/global"
	"github.com/goprogramming/blog/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"Modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB,error){
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
	databaseSetting.UserName,
	databaseSetting.Password,
	databaseSetting.Host,
	databaseSetting.DBName,
	databaseSetting.Charset,
	databaseSetting.ParseTime,
)
	db,err := gorm.Open(databaseSetting.DBType,s)
	if err != nil{
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug"{
		db.LogMode(true)
	}
	db.SingularTable(true)//默认使用单独的一个表  
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db,nil
}
