package models

type GoAdminPermission struct {
	Model
	Name       string `gorm:"not null;unique"`
	Slug       string `gorm:"not null"`
	HttpMethod *string
	HttpPath   string `gorm:"not null"`
}

func (gap *GoAdminPermission) TableName() string {
	return "goadmin_permissions"
}
