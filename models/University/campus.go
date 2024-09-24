package models

type Campus struct {
	CampusID             uint   `gorm:"primaryKey"`
	CampusName           string `gorm:"type:varchar(100);not null"`
	RelationalUniversity uint   `gorm:"not null"` // 外键关联到 University
}
