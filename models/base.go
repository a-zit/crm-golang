package models

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Model struct {
	ID        uint           `gorm:"primary_key"`
	CreatedAt *time.Time     `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time     `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `sql:"index"`
}

func init() {
	e := godotenv.Load()

	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Print(err)
	}

	db = conn
	//Automatically create migration as per model
	db.Debug().AutoMigrate(
		&Admin{}, &Branch{}, &Brand{},
		&BrandLevel{}, &BroadcastLog{}, &ReadMessageLog{},
		&Reward{}, &RewardBranch{}, &RewardLog{}, &User{},
		&GoAdminMenu{}, &GoAdminOperationLog{}, &GoAdminPermission{},
		&GoAdminRole{}, &GoAdminRoleMenu{}, &GoAdminRolePermission{},
		&GoAdminSession{}, &GoAdminRoleUser{}, &GoAdminSite{},
		&GoAdminUserPermission{},
	)
}
