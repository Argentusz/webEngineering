package config

type Config struct {
	Host   string
	Port   string
	DbName string
	DbHost string
	DbPort string
	DbUser string
	DbPass string
}

func (cfg *Config) Setup() {
	cfg.Host = "localhost"
	cfg.Port = "8081"
	cfg.DbName = "postgres"
	cfg.DbHost = "localhost"
	cfg.DbPort = "2229"
	cfg.DbUser = "mt"
	cfg.DbPass = "password"
}
