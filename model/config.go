package model

type Config struct {
	Database DatabaseConfig `json:"database"`
	Jwt      JwtConfig      `json:"jwt"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type JwtConfig struct {
	SecretKey string `json:"secret_key"`
	ExpTime   int    `json:"exp_time"`
}
