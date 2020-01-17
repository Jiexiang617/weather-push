package model

func migration() {
	DB.AutoMigrate(&User{})
}