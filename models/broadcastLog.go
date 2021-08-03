package models

import "database/sql/driver"

type BroadcastLog struct {
	Model
	Title    *string
	ImageURL *string
	LinkURL  *string
	Detail   *string
	Type     *BroadcastLogType `gorm:"type:ENUM('TXT', 'IMG', 'C', 'CC', 'IMGM')"`
	AdminID  int               `gorm:"not null"`
	Admin    Admin             `gorm:"foreignKey:AdminID;not null"`
	RewardID *int
	Reward   Reward `gorm:"foreignKey:RewardID"`
}

type BroadcastLogType string

const (
	Text         BroadcastLogType = "TXT"
	Image        BroadcastLogType = "IMG"
	Card         BroadcastLogType = "C"
	CardCollusel BroadcastLogType = "CC"
	ImageMap     BroadcastLogType = "IMGM"
)

func (bclt *BroadcastLogType) Scan(value interface{}) error {
	*bclt = BroadcastLogType(value.([]byte))
	return nil
}

func (bclt BroadcastLogType) Value() (driver.Value, error) {
	return string(bclt), nil
}

func (bcl *BroadcastLog) TableName() string {
	return "broadcast_logs"
}
