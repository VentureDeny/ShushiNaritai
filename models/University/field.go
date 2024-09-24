package models

type Field struct {
	FieldID   uint   `gorm:"primaryKey"`
	FieldName string `gorm:"type:varchar(100);not null"`
}
