package models

type BrandLevel struct {
	Model
	LevelName       string `gorm:"not null"`
	NumberCondition int    `gorm:"not null"`
	BrandID         int    `gorm:"not null"`
	Brand           Brand  `gorm:"foreignKey:BrandID"`
}

func (bl *BrandLevel) TableName() string {
	return "brand_levels"
}
