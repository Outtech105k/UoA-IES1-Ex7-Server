package cmd

import (
	"fmt"
	"os"

	"github.com/Outtech105k/UoA-IES1-Ex7-Server/models"
	"gopkg.in/yaml.v3"
)

func ReadAppConfig(path string) (models.AppConfig, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return models.AppConfig{}, fmt.Errorf("openFile: %w", err)
	}

	var conf models.AppConfig
	if err := yaml.Unmarshal(f, &conf); err != nil {
		return models.AppConfig{}, fmt.Errorf("decode yaml: %w", err)
	}

	return conf, nil
}
