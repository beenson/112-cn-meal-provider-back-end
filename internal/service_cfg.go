package internal

type ServiceCfg struct {
	BillingTarget      string `env:"BILLING_TARGET"`
	NotificationTarget string `env:"NOTIFICATION_TARGET"`
	OrderingTarget     string `env:"ORDERING_TARGET"`
	RatingTarget       string `env:"RATING_TARGET"`
	UserMgmtTarget     string `env:"USER_MGMT_TARGET"`
}
