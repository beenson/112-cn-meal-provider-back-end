package main

import "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/internal/db"

type config struct {
	DB          db.Config `envPrefix:"DB_"`
	GRPCAddress string    `env:"GRPC_ADDRESS"`
}
