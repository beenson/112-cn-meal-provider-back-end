package model

type Feedback struct {
	ID      uint `gorm:"primaryKey"`
	Foodid  uint
	Uid     string
	Rating  int32
	Comment string
}

func init() {
	models = append(models, &Feedback{})
}
