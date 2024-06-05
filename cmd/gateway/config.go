package main

import "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/internal"

type config struct {
	Service   internal.ServiceCfg `envPrefix:"SERVICE_"`
	JwtSecret string              `env:"JWT_SECRET" envDefault:"secret"`
	Address   string              `env:"ADDRESS" envDefault:":55688"`
}
