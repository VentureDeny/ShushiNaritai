package models

type ExamType struct {
	ExamTypeID     uint   `gorm:"primaryKey"`
	ExamTypeName   string `gorm:"type:varchar(100);not null"`
	ExamTypeDetail string `gorm:"type:varchar(255)"`
}
