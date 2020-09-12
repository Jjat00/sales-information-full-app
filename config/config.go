package config

type Config struct {
	DB     *DBConfig
	Server *ServerConfig
}

type DBConfig struct {
	Dialect string
	Host    string `required:"true"`
	Port    string `required:"true"`
}

type ServerConfig struct {
	Port string `required:"true"`
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect: "dgraph",
			Host:    "127.0.0.1",
			Port:    ":9080",
		},
		Server: &ServerConfig{
			Port: ":3000",
		},
	}
}
