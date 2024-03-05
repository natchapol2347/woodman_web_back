package config

type Database struct { // TODO: private type
	Host     string `env:"POSTGRESQL_URI"` // TODO: rename to URI
	Name     string `env:"POSTGRESQL_NAME"`
	Username string `env:"POSTGRESQL_USERNAME"`
	Password string `env:"POSTGRESQL_PASSWORD"`
	Port     string `env:"POSTGRESQL_PORT"`
}
