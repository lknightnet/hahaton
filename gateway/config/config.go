package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type (
	Config struct {
		App          `yaml:"app"`
		ServerCfg    `yaml:"serverCfg"`
		ServicesPort `yaml:"servicesPort"`
	}

	App struct {
		Name string `yaml:"name"`
	}

	ServerCfg struct {
		Port            string        `yaml:"port"`
		ReadTimeout     time.Duration `yaml:"readTimeout"`
		WriteTimeout    time.Duration `yaml:"writeTimeout"`
		ShutdownTimeout time.Duration `yaml:"shutdownTimeout"`
	}

	ServicesPort struct {
		BellPort  string `yaml:"bellPort"`
		AudioPort string `yaml:"audioPort"`
		TransPort string `yaml:"transPort"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
