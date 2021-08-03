package models

import "time"

type GoAdminRoleMenu struct {
	CreatedAt *time.Time  `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time  `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	RoleID    int         `gorm:"not null"`
	Role      GoAdminRole `gorm:"foreignKey:RoleID"`
	MenuID    int         `gorm:"not null"`
	Menu      GoAdminMenu `gorm:"foreignKey:MenuID"`
}

func (garm *GoAdminRoleMenu) TableName() string {
	return "goadmin_role_menu"
}
