package config

type Config struct {
	HTTPAddr    string `json:"HTTP_ADDR"`
	MirationDir string `json:"MIG_DIR"`
	PG_DSN      string `json:"PG_DSN"`
}

// TODO: read from env / dotfiles
func FromEnv() Config {
	return Config{
		HTTPAddr:    ":8080",
		MirationDir: "internal/db/migrations/",
		PG_DSN:      "postgres://username:password@localhost:5432/wallet_restapi",
	}
}
