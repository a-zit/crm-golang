package models

import "time"

type RewardLog struct {
	Model
	Detail     *string
	Code       *string
	RedeemedAt *time.Time
	BranchID   int    `gorm:"not null"`
	Branch     Branch `gorm:"foreignKey:BranchID"`
	RewardID   int    `gorm:"not null"`
	Reward     Reward `gorm:"foreignKey:RewardID"`
	UserID     int    `gorm:"not null"`
	User       User   `gorm:"foreignKey:UserID"`
}

func (rl *RewardLog) TableName() string {
	return "reward_logs"
}
