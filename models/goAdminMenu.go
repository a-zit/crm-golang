package models

type GoAdminMenu struct {
	Model
	ParentID   int    `gorm:"not null;default:0"`
	Type       int    `gorm:"type:bool;default:0;not null"`
	Order      int    `gorm:"not null;default:0"`
	Title      string `gorm:"not null"`
	Icon       string `gorm:"not null"`
	URI        string `gorm:"not null"`
	Header     *string
	PluginName string `gorm:"not null"`
	UUID       *string
}

func (gam *GoAdminMenu) TableName() string {
	return "goadmin_menu"
}
