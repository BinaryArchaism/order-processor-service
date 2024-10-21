package config

import (
	"context"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DBConfig DBConfig `yaml:"database"`
}

type DBConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`

	ConnMaxLifetime int `yaml:"conn_max_lifetime"`
	MaxOpenConns    int `yaml:"max_open_conns"`
	MaxIdleConns    int `yaml:"max_idle_conns"`
	ConnMaxIdleTime int `yaml:"conn_max_idle_time"`
}

func (dbCfg DBConfig) ConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Database)
}

func InitConfig(_ context.Context) (Config, error) {
	yamlFile, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return Config{}, err
	}
	var cfg Config
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
