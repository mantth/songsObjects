package model

// 数据库自动迁移；
func migration() {
	Db.Set("gorm:table_option", "charset=utf8mb4").AutoMigrate(&User{})
}
