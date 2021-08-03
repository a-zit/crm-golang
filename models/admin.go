package models

import "database/sql/driver"

type Admin struct {
	Model
	FirstName     *string
	LastName      *string
	Name          string `gorm:"not null"`
	Username      string `gorm:"not null:unique"`
	Password      string `gorm:"not null"`
	Avatar        *string
	Email         string `gorm:"unique;not null"`
	Tel           string `gorm:"unique;not null"`
	Level         *Level `gorm:"type:ENUM('SA', 'BDA', 'BHM', 'BHS')"`
	RememberToken *string
	BrandID       *int
	Brand         Brand `gorm:"foreignKey:BrandID"`
	BranchID      *int
	Branch        Branch `gorm:"foreignKey:BranchID"`
}

type Level string

const (
	SuperAdmin    Level = "SA"
	BrandAdmin    Level = "BDA"
	BranchManager Level = "BHM"
	BranchStaff   Level = "BHS"
)

func (l *Level) Scan(value interface{}) error {
	*l = Level(value.([]byte))
	return nil
}

func (l Level) Value() (driver.Value, error) {
	return string(l), nil
}

func (a *Admin) TableName() string {
	return "goadmin_users"
}
