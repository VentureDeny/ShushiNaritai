package models

type PlanAdvice struct {
	PlanAdviceID     uint   `gorm:"primaryKey"`
	PlanAdviceType   string `gorm:"type:varchar(100);not null"`
	PlanAdviceDetail string `gorm:"type:varchar(255)"`
}
