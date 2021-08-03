package models

type GoAdminSite struct {
	Model
	Key         *string
	Value       string
	Description *string
	State       int `gorm:"type:bool;default:0;not null"`
}

func (gas *GoAdminSite) TableName() string {
	return "goadmin_site"
}
