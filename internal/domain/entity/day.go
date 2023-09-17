package entity

type Day struct {
	ID   uint   `json:"id" gorm:"autoIncreament;primaryKey"`
	Name string `json:"day"`
}
