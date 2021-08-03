package models

import "time"

type GoAdminRoleUser struct {
	CreatedAt *time.Time  `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time  `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	RoleID    int         `gorm:"not null"`
	Role      GoAdminRole `gorm:"foreignKey:RoleID"`
	UserID    int         `gorm:"not null"`
	User      Admin       `gorm:"foreignKey:UserID"`
}

func (garu *GoAdminRoleUser) TableName() string {
	return "goadmin_role_users"
}
