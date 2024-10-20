package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
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
}

func (dbCfg DBConfig) ConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Database)
}

func Parse() (Config, error) {
	fmt.Println(os.Getwd())
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
