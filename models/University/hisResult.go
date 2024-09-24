package models

type HisResult struct {
	ID             uint `gorm:"primaryKey"`
	RgsNum         int  `gorm:"not null"` // 出愿人数
	RgsForeignNum  int  `gorm:"not null"` // 出愿外国人数
	ExamNum        int  `gorm:"not null"` // 参加考试人数
	ExamForeignNum int  `gorm:"not null"` // 参加考试外国人数
	PassNum        int  `gorm:"not null"` // 合格人数
	PassForeignNum int  `gorm:"not null"` // 合格外国人数
}
