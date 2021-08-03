package models

type RewardBranch struct {
	Model
	RewardID int    `gorm:"not null"`
	Reward   Reward `gorm:"foreignKey:RewardID"`
	BranchID int    `gorm:"not null"`
	Branch   Branch `gorm:"foreignKey:BranchID"`
}

func (rb *RewardBranch) TableName() string {
	return "reward_branches"
}
