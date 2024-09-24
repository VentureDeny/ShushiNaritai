package models

type GuidanceMenu struct {
	ID               uint           `gorm:"primaryKey"`
	GuidanceMenuName string         `gorm:"type:varchar(100);not null"`
	SubMenus         []GuidanceMenu `gorm:"foreignKey:ParentID"`
	ParentID         *uint          `gorm:"default:null"`
}
