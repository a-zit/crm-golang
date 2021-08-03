package models

type GoAdminOperationLog struct {
	Model
	UserID int    `gorm:"not null"`
	Path   string `gorm:"not null"`
	Method string `gorm:"not null"`
	IP     string `gorm:"not null"`
	Input  string `gorm:"not null"`
	Admin  Admin  `gorm:"foreignKey:UserID"`
}

func (gaol *GoAdminOperationLog) TableName() string {
	return "goadmin_operation_log"
}
