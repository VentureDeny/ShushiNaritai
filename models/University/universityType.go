package models

type UniversityType struct {
	UniversityTypeID     uint   `gorm:"primaryKey"`
	UniversityTypeName   string `gorm:"type:varchar(100);not null"`
	UniversityTypeDetail string `gorm:"type:varchar(255)"`
}
