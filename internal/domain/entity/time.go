package entity

type Time struct {
	ID   uint   `json:"id" gorm:"autoIncreament;primaryKey"`
	Time string `json:"time"`
}
