package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const DEFAULT_CONFIG = `
[server]
host = "0.0.0.0"
port = 7001

[auth]
secret = "IMG_HOSTING"
expiredTime = 86400

[middlewares]

    [middlewares.cors]
    active = true
    [middlewares.logger]
    active = true

[databases]

    [databases.mysql]
    host = "127.0.0.1"
    port = 3306
    username = "root"
    password = "root"
    dbname = "img_hosting"
`

type Config struct {
	Server      serverConfig
	Auth        AuthConfig
	Databases   map[string]databaseConfig
	Middlewares map[string]middlewareConfig
}

type AuthConfig struct {
	Secret      string
	ExpiredTime int
}

type serverConfig struct {
	Host string
	Port int
}

type databaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

type middlewareConfig struct {
	Active bool
}

func check(conf *Config) error {
	return nil
}

func Initialize() (Config, error) {

	var conf Config

	root, err := os.Getwd()
	if err != nil {
		root = filepath.Dir(os.Args[0])
	}

	_, err := toml.DecodeFile(filepath.Join(root, "config/default.toml"), &conf)

	if err == nil {

		log.Print("Custom configutation has been loaded successfully.")

	} else {
		_, err := toml.Decode(DEFAULT_CONFIG, &conf)

		if err != nil {
			return Config{}, err
		}

		log.Print("Failed to load custom configuration. Use default.")

	}

	err = nil
	err = check(&conf)

	return conf, err

}
