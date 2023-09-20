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
	FrontEnd string
}

func (cfg *Config) Setup() {
	//cfg.Host = "localhost"
	//cfg.Port = "8081"
	//cfg.DbName = "postgres"
	//cfg.DbHost = "localhost"
	//cfg.DbPort = "2229"
	//cfg.DbUser = "mt"
	//cfg.DbPass = "password"
	//cfg.CertFile = "localhost.crt"
	//cfg.KeyFile = "localhost.key"
	//cfg.FrontEnd = "frontend/dist"
	cfg.Host = "server"
	cfg.Port = "8081"
	cfg.DbName = "postgres"
	cfg.DbHost = "database"
	cfg.DbPort = "5432"
	cfg.DbUser = "mt"
	cfg.DbPass = "password"
	cfg.CertFile = "localhost.crt"
	cfg.KeyFile = "localhost.key"
	cfg.FrontEnd = "frontend/dist"
}
