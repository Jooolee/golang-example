package test

import (
	"log"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func TestGormGen(*testing.T) {
	// 连接 MySQL 数据库
	dsn := "root:root@tcp(127.0.0.1:3306)/public?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		os.Exit(1)
	}
	log.Println("数据库连接成功")

	// 创建代码生成器实例
	g := gen.NewGenerator(gen.Config{
		OutPath:      "../dal/query",
		ModelPkgPath: "../dal/model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// 设置数据库连接
	g.UseDB(db)

	// 生成所有表的模型
	allTables := g.GenerateAllTable()

	g.ApplyBasic(allTables...)

	// 执行生成操作
	g.Execute()
}
