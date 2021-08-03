package models

import "time"

type GoAdminUserPermission struct {
	CreatedAt    *time.Time        `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt    *time.Time        `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UserID       int               `gorm:"not null"`
	User         Admin             `gorm:"foreignKey:UserID"`
	PermissionID int               `gorm:"not null"`
	Permission   GoAdminPermission `gorm:"foreignKey:PermissionID"`
}

func (gaup *GoAdminUserPermission) TableName() string {
	return "goadmin_user_permissions"
}
