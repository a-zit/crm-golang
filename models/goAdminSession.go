package models

type GoAdminSession struct {
	Model
	SID    string `gorm:"not null"`
	Values string `gorm:"not null"`
}

func (gas *GoAdminSession) TableName() string {
	return "goadmin_session"
}
