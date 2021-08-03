package models

import (
	"database/sql/driver"
	"time"
)

type UserType string

const (
	Partner UserType = "PARTNER"
	Normal  UserType = "NORMAL"
)

func (e *UserType) Scan(value interface{}) error {
	*e = UserType(value.([]byte))
	return nil
}

func (e UserType) Value() (driver.Value, error) {
	return string(e), nil
}

type User struct {
	Model
	UID           *string  `gorm:"unique"`
	Type          UserType `gorm:"type:ENUM('PARTER', 'NORMAL')"`
	FirstName     string   `gorm:"not null"`
	LastName      string   `gorm:"not null"`
	Email         string   `gorm:"not null;unique"`
	Tel           string   `gorm:"size:10;not null;unique"`
	Level         *string
	Birthdate     *time.Time
	TotalPoint    int     `gorm:"default:0"`
	TotalScore    int     `gorm:"default:0"`
	OtpCode       *string `gorm:"size:6"`
	OtpVerifiedAt *time.Time
	OtpExpiredAt  *time.Time
	BrandID       int   `gorm:"not null"`
	Brand         Brand `gorm:"foreignKey:BrandID"`
}

func (u *User) UserList() (interface{}, error) {
	users := []User{}

	result := db.Joins("Brand").Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (u *User) CreateUser() error {
	result := db.Create(&u)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) TableName() string {
	return "users"
}
