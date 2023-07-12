// @Title  mysql.go
// @Description  Mysql初始化以及相关操作

package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewMysql(mysqlConfig string) (*gorm.DB, error) {
	fmt.Println("数据库开始连接")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		"dbUser",
		"dbPass",
		"dbAddr",
		"dbName",
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info), // 日志配置
	})
	if err != nil {
		fmt.Println("数据库正常连接")
		return nil, err
	}

	//db.Logger.LogMode(false)
	//db.SingularTable(true)
	//db.LogMode(false)

	return db, nil
}
