package models

type AfterAdvice struct {
	AfterAdviceID     uint   `gorm:"primaryKey"`
	AfterAdviceType   string `gorm:"type:varchar(100);not null"`
	AfterAdviceDetail string `gorm:"type:varchar(255)"`
}
