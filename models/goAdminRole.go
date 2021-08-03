package models

type GoAdminRole struct {
	Model
	Name string `gorm:"not null:unique"`
	Slug string `gorm:"not null"`
}

func (gar *GoAdminRole) TableName() string {
	return "goadmin_roles"
}
