package config

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		Server   ServerConfig
		Tracker  TrackerConfig
		Info     InfoConfig
		Postgres PostgresConfig
	}

	ServerConfig struct {
		Port            int    `yaml:"server_port"`
		Host            string `yaml:"server_host"`
		ReadTimeout     int    `yaml:"server_read_timeout"`
		WriteTimeout    int    `yaml:"server_write_timeout"`
		ShutdownTimeout int    `yaml:"server_shutdown_timeout"`
	}

	TrackerConfig struct {
		Port int    `yaml:"tracker_port"`
		Host string `yaml:"tracker_host"`
	}

	InfoConfig struct {
		Port int    `yaml:"info_port"`
		Host string `yaml:"info_host"`
	}

	PostgresConfig struct {
		IP       string `yaml:"pg_ip"`
		Port     int    `yaml:"pg_port"`
		User     string `yaml:"pg_user"`
		Password string `yaml:"pg_password"`
		DB       string `yaml:"pg_db"`
		Timeout  int    `yaml:"pg_timeout"`
	}
)

func Init() (Config, error) {
	configPath := "./config/config.yaml"

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return Config{}, nil
	}

	return cfg, nil
}
