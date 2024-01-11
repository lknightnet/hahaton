package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
	"time"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		API  `yaml:"api"`
		PG   `yaml:"pg"`
	}

	App struct {
		Name string `yaml:"name" env:"APP_NAME"`
	}
	HTTP struct {
		Port            string        `yaml:"port" env:"HTTP_PORT"`
		ReadTimeout     time.Duration `json:"read_timeout" env:"HTTP_READ_TIMEOUT"`
		WriteTimeout    time.Duration `json:"write_timeout" env:"HTTP_WRITE_TIMEOUT"`
		ShutdownTimeout time.Duration `json:"shutdown_timeout" env:"HTTP_SHUTDOWN_TIMEOUT"`
	}

	API struct {
		UrlAPI string `yaml:"url_api" env:"API_URL"`
	}

	PG struct {
		ConnURI string `yaml:"conn_uri" env:"PG_URI"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, errors.Wrap(err, "NewConfig: fail to read config.yaml file")
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "NewConfig: fail to read env")
	}
	return cfg, err
}
