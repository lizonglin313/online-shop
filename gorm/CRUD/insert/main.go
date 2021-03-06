package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string	// 使用指针和null string都可以解决空字符串的问题
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}


func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	// 设置全局logger 打印每一行sql
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.Lshortfile),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// 进行数据库的 配置和连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// 定义表结构 直接生成对应的表 migration
	// _ = db.AutoMigrate(&User{})

	// db.Create(&User{Name: "lzl"})
	// Update会更新0值 而 Updates不会更新0值 使用 Updates 需要考虑 NullString这种类型
	// db.Model(&User{ID: 1}).Update("Name", "")
	// db.Model(&User{ID: 1}).Updates(User{Name: ""})

	// Updates更新0值
	//empty := ""
	//db.Model(&User{ID: 1}).Updates(User{Email: &empty})

	// ============= 进行批量插入 =============
	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(&users)

	for _, user := range users {
		fmt.Println(user.ID) // 1,2,3
	}

	// 如果数据量过大 超出一条 sql 语句的限制 那么可以使用 指定一次批量插入的数量
	// 拆分成多条 sql 语句来执行
	// db.CreateBatchSize(users, 20)

}
