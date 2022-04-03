package main

import (
	"database/sql"
	"errors"
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
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// 通过take，first、last获取数据
func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取参数详情
	// 想要正确的处理 time.Time ，您需要带上 parseTime 参数， (更多参数) 要支持完整的 UTF-8 编码，
	// 您需要将 charset=utf8 更改为 charset=utf8mb4
	dsn := "admin:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 启用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("open db failed, err: %v", err)
	}

	fmt.Printf("connect %s success\n", db.Name())

	// 迁移 schema
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("AutoMigrate failed, err: %s", err.Error())
	}

	// 会生成单独一条SQL语句，而不是分多条插入
	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(&users)

	var firstUser User
	db.First(&firstUser)
	fmt.Println(firstUser.ID)

	var takeUser User
	db.Take(&takeUser)
	fmt.Println(takeUser.ID)

	var lastUser User
	db.Last(&lastUser)
	fmt.Println(lastUser.ID)

	var notExitUser User
	result := db.First(&notExitUser, 10)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("can not find user: ", 10)
	}
}
