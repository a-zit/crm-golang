package models

type ReadMessageLog struct {
	Model
	Count          *int
	BroadcastLogID int          `gorm:"not null"`
	BroadcastLog   BroadcastLog `gorm:"foreignKey:BroadcastLogID"`
	UserID         int          `gorm:"not null"`
	User           User         `gorm:"foreignKey:UserID"`
}

func (rml *ReadMessageLog) TableName() string {
	return "read_message_logs"
}
