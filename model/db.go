package model

import (
	"fake_twitter/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(
		utils.Db,
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			utils.DbUser,
			utils.DbPassWord,
			utils.DbHost,
			utils.DbPort,
			utils.DbName,
		))
	if err != nil {
		fmt.Printf("连接数据库失败，请检查参数：", err)
	}

	//创建出来的表名为单数形式。  假如有了 users 不会创建 users、user
	//要放在所有的表操作前面
	db.SingularTable(true)

	db.AutoMigrate(&User{}, &Article{}, &Category{})

	db.DB().SetMaxIdleConns(10)

	db.DB().SetMaxOpenConns(100)

	db.DB().SetConnMaxLifetime(10 * time.Second)

}
