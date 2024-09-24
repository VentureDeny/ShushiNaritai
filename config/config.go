package config

import (
	"os"
)

type Config struct {
	Port  string
	DBDsn string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // 默认端口为8080
	}

	return &Config{
		Port:  port,
		DBDsn: os.Getenv("DB_DSN"), // 从环境变量中读取数据库连接字符串
	}
}
