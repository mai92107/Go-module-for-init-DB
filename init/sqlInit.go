package init

import (
	constVar "batchLog/const"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	// 環境變數設定（或直接寫死在程式中）
	dbconst := constVar.DBConst
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",dbconst.DBUser,dbconst.DBPassword,dbconst.DBHost,dbconst.DBPort,dbconst.DBName)

	fmt.Println(dsn)
	// 連接資料庫
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ 無法連接資料庫:", err)
	}

	log.Println("✅ 資料庫連接成功")
	constVar.DB = db
}