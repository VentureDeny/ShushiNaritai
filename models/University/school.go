package models

type School struct {
	SchoolID             uint   `gorm:"primaryKey"`
	SchoolName           string `gorm:"type:varchar(100);not null"`
	CampusID             uint   `gorm:"not null"`
	SchoolDetail         string `gorm:"type:varchar(255)"`
	RelationalUniversity uint   `gorm:"not null"` // 外键关联到 University
	RelationalRegion     uint   `gorm:"not null"` // 外键关联到 Region
}
