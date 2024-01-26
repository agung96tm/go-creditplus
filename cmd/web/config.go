package main

type Config struct {
	Addr      string
	SecretKey string
	DB        DBConfig
}

type DBConfig struct {
	dsn string
}

func DefaultConfig() Config {
	return Config{}
}
