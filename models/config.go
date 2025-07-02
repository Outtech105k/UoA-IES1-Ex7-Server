package models

type AppConfig struct {
	Slack SlackConfig `yaml:"slack"`
}

type SlackConfig struct {
	Webhook WebhookConfig `yaml:"webhook"`
}

type WebhookConfig struct {
	Url string `yaml:"url"`
}
