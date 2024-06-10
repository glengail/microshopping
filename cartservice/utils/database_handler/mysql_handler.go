package databasehandler

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewMysqlDB(conString string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{
		PrepareStmt: true, //为 SQL 语句启用预编译，以提高执行效率
		NamingStrategy: schema.NamingStrategy{ //数据表命名策略的配置
			SingularTable: true, //表名使用单数形式
			NoLowerCase:   true, //不将表名转换为小写
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		panic(fmt.Sprintf("不能连接到数据库：%s", err.Error()))
	}
	return db
}
