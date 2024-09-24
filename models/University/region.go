package models

type Region struct {
	RegionID   uint   `gorm:"primaryKey"`
	RegionName string `gorm:"type:varchar(100);not null"`
}
