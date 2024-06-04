package model

import "time"

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Uid        string `gorm:"uniqueIndex"`
	Group      string
	Username   string
	Department string
	Email      string
	PublicKey  string
}

type UserAuth struct {
	ID        uint   `gorm:"primaryKey"`
	Uid       string `gorm:"uniqueIndex"`
	Challenge string
	AuthAt    time.Time
	AuthToken string
}

func init() {
	models = append(models, &User{})
	models = append(models, &UserAuth{})
}
