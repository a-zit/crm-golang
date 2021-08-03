package models

type Branch struct {
	Model
	Name        string `gorm:"not null"`
	BranchImage *string
	Latitude    *float64
	Longitude   *float64
	BrandID     int   `gorm:"not null"`
	Brand       Brand `gorm:"foreignKey:BrandID"`
}

func (b *Branch) TableName() string {
	return "branches"
}
