package config

type Config struct {
	Host     string
	Port     string
	DbName   string
	DbHost   string
	DbPort   string
	DbUser   string
	DbPass   string
	CertFile string
	KeyFile  string
}

func (cfg *Config) Setup() {
	cfg.Host = "localhost"
	cfg.Port = "8081"
	cfg.DbName = "postgres"
	cfg.DbHost = "localhost"
	cfg.DbPort = "2229"
	cfg.DbUser = "mt"
	cfg.DbPass = "password"
	cfg.CertFile = "localhost.crt"
	cfg.KeyFile = "localhost.key"
}
