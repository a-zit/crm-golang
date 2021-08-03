package models

import "time"

type GoAdminRolePermission struct {
	CreatedAt    *time.Time        `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt    *time.Time        `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	RoleID       int               `gorm:"not null"`
	Role         GoAdminRole       `gorm:"foreignKey:RoleID"`
	PermissionID int               `gorm:"not null"`
	Permission   GoAdminPermission `gorm:"foreignKey:PermissionID"`
}

func (garp *GoAdminRolePermission) TableName() string {
	return "goadmin_role_permissions"
}
