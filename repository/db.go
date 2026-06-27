package repository

import (
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"heritage-go/model"
)

var DB *gorm.DB

// InitDB 自动化初始化本地 SQLite 数据库
func InitDB() {
	dbDir := "storage"
	// 自动创建本地存储目录
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		_ = os.MkdirAll(dbDir, os.ModePerm)
	}

	dbPath := filepath.Join(dbDir, "genealogy.db")
	
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // 生产环境关闭冗余日志
	})
	if err != nil {
		log.Fatalf("SQLite 数据库连接失败: %v", err)
	}

	// 核心：自动物理建表与索引字段对齐（GORM AutoMigrate）
	err = DB.AutoMigrate(&model.Member{})
	if err != nil {
		log.Fatalf("数据库表自动初始化失败: %v", err)
	}
	log.Println("[💾 数据库] SQLite 本地底座初始化成功，表结构已对齐。")
}
