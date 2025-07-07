package models

type AppConfig struct {
	Slack   SlackConfig   `yaml:"slack"`
	Sqlite3 Sqlite3Config `yaml:"sqlite3"`
}

type SlackConfig struct {
	Webhook WebhookConfig `yaml:"webhook"`
}

type WebhookConfig struct {
	Url string `yaml:"url"`
}

type Sqlite3Config struct {
	Path string `yaml:"path"`
}
