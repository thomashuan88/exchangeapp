package dao

type Role struct {
	ID   int    `gorm:"column:id;primary_key" json:"id"`
	Role string `gorm:"column:role" json:"role"`
	BaseModel
}
