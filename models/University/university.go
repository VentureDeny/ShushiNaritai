package models

type University struct {
	UniversityID     uint           `gorm:"primaryKey"`
	UniversityName   string         `gorm:"type:varchar(100);not null"`
	UniversityDetail string         `gorm:"type:varchar(255)"`
	UniversityTypeID uint           `gorm:"not null"` // 外键关联到 UniversityType
	UniversityType   UniversityType `gorm:"foreignKey:UniversityTypeID"`
}
