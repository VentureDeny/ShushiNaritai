package models

type Dpm struct {
	DpmID            uint   `gorm:"primaryKey"`
	DpmName          string `gorm:"type:varchar(100);not null"`
	DpmDetail        string `gorm:"type:varchar(255)"`
	RelationalSchool uint   `gorm:"not null"` // 外键关联到 School
	RelationalField  uint   `gorm:"not null"` // 外键关联到 Field
	RelationalExam   uint   `gorm:"not null"` // 外键关联到 Exam
}
