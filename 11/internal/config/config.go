package config

type Config struct {
	Port string
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	cfg.Port = "8080"
	return cfg, nil
}
