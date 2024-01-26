package main

type Config struct {
	Addr      string
	SecretKey string
	DB        DBConfig
	cors      CorsConfig
}

type DBConfig struct {
	dsn string
}

type CorsConfig struct {
	trustedOrigins []string
}

func DefaultConfig() Config {
	return Config{
		cors: CorsConfig{
			trustedOrigins: []string{
				"http://localhost:3000",
				"http://localhost:4000",
				"http://localhost:5000",
			},
		},
	}
}
