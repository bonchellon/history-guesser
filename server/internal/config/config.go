package config

import "os"

type Config struct { Port, DatabaseURL string }

func Load() Config {
	return Config{
		Port: get("SERVER_PORT", "8080"),
		DatabaseURL: get("DATABASE_URL", "postgres://history:history@localhost:5432/history?sslmode=disable"),
	}
}

func get(k, d string) string { if v:=os.Getenv(k); v!="" { return v }; return d }
