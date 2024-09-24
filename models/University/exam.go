package models

type Exam struct {
	ExamID           uint   `gorm:"primaryKey"`
	ExamType         uint   `gorm:"not null"` // 外键关联到 ExamType
	RelationalDpm    uint   `gorm:"not null"` // 外键关联到 Dpm
	RegistrationTime string `gorm:"type:date"`
	WrittenTime      string `gorm:"type:date"`
	WrittenDetail    string `gorm:"type:varchar(255)"`
	InterviewTime    string `gorm:"type:date"`
	InterviewDetail  string `gorm:"type:varchar(255)"`
}
