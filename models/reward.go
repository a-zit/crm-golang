package models

import (
	"database/sql/driver"
	"time"
)

type Reward struct {
	Model
	Name          string `gorm:"not null"`
	Detail        *string
	Type          *RewardType `gorm:"type:ENUM('DBP', 'DBA', 'O')"`
	DiscountValue *int
	Quantity      int `gorm:"not null;default:0"`
	RequirePoint  int `gorm:"not null;default:0"`
	RequireLevel  *string
	BeginDate     *time.Time
	EndDate       *time.Time
	BrandID       *int
	Brand         Brand `gorm:"foreignKey:BrandID"`
}

type RewardType string

const (
	DiscountByPercent RewardType = "DBP"
	DiscountByAmount  RewardType = "DBA"
	Object            RewardType = "O"
)

func (rt *RewardType) Scan(value interface{}) error {
	*rt = RewardType(value.([]byte))
	return nil
}

func (rt RewardType) Value() (driver.Value, error) {
	return string(rt), nil
}

func (r *Reward) TableName() string {
	return "rewards"
}
