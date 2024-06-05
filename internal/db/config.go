package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Type string `env:"TYPE" envDefault:"sqlite"`
	DSN  string `env:"DSN" envDefault:"file::memory:?cache=shared"`
}

func (cfg Config) GetGormDB() (*gorm.DB, error) {
	if cfg.Type == "sqlite" {
		return gorm.Open(sqlite.Open(cfg.DSN), &gorm.Config{})
	} else if cfg.Type == "mysql" {
		return gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
	}

	panic("Unsupported database type: " + cfg.Type)
}
