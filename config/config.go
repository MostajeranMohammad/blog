package config

import (
	"fmt"

	"github.com/MostajeranMohammad/blog/pkg/utils"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
	}

	// App -.
	App struct {
		Name    string `validate:"required"   yaml:"name"    env:"APP_NAME"`
		Version string `validate:"required"   yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		JwtSecret string `validate:"required" yaml:"jwt_secret" env:"JWT_SECRET"`
		Port      string `validate:"required" yaml:"port"       env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `validate:"required"     yaml:"log_level" env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		DSN string ` validate:"required"      yaml:"pg_dsn"    env:"PG_DSN"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := parseConfigFiles([]string{"./config/config.yml", "./.env"}, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.UpdateEnv(cfg)
	if err != nil {
		return nil, err
	}

	err = utils.ValidateDto(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func parseConfigFiles(files []string, cfg *Config) error {
	for _, path := range files {
		err := cleanenv.ReadConfig(path, cfg)
		if err != nil {
			return fmt.Errorf("config error: %w", err)
		}
	}
	return nil
}
