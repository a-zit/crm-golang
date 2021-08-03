package models

type Brand struct {
	Model
	Name       string `gorm:"not null"`
	BrandImage *string
}

func (b *Brand) TableName() string {
	return "brands"
}
