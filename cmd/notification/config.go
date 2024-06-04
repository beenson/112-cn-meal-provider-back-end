package main

import "gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/internal"

type config struct {
	GRPCAddress string              `env:"GRPC_ADDRESS"`
	Service     internal.ServiceCfg `envPrefix:"SERVICE_"`

	EmailFromAddress string `env:"EMAIL_FROM_ADDRESS" envDefault:"no-reply@example.com"`

	//SMTPAddress string `env:"SMTP_ADDRESS"`
	//SMTPUser    string `env:"SMTP_USER"`
	//SMTPPass    string `env:"SMTP_PASS"`
}
