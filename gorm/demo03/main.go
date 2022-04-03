package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// 简单的增删改查
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
	if err := db.AutoMigrate(&Product{}); err != nil {
		log.Fatalf("AutoMigrate failed, err: %s", err.Error())
	}
	
	// Create
	// INSERT INTO `products` (`created_at`,`updated_at`,`deleted_at`,`code`,`price`) VALUES ('2022-03-27 22:19:29.825','2022-03-27 22:19:29.825',NULL,'D42',100)
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	// SELECT * FROM `products` WHERE `products`.`id` = 1 AND `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1
	db.First(&product, 1) // 根据整型主键查找
	// SELECT * FROM `products` WHERE code = 'D42' AND `products`.`deleted_at` IS NULL AND `products`.`id` = 1 ORDER BY `products`.`id` LIMIT 1
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	// UPDATE `products` SET `price`=200,`updated_at`='2022-03-27 22:19:29.828' WHERE `products`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	// UPDATE `products` SET `updated_at`='2022-03-27 22:19:29.829',`code`='F42',`price`=200 WHERE `products`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// UPDATE `products` SET `code`='F42',`price`=200,`updated_at`='2022-03-27 22:19:29.83' WHERE `products`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//  UPDATE `products` SET `deleted_at`='2022-03-27 22:19:29.832' WHERE `products`.`id` = 1 AND `products`.`id` = 1 AND `products`.`deleted_at` IS NULL
	db.Delete(&product, 1)
}
