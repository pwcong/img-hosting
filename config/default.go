package config

import (
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"img-hosting/utils"
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
	Server      ServerConfig
	Auth        AuthConfig
	Databases   map[string]DatabaseConfig
	Middlewares map[string]MiddlewareConfig
}

type AuthConfig struct {
	Secret      string
	ExpiredTime int
}

type ServerConfig struct {
	Host string
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	DBPath   string
}

type MiddlewareConfig struct {
	Active bool
}

func check(conf *Config) error {
	return nil
}

func Initialize() (Config, error) {

	var conf Config

	rootDir := utils.GetRootDir()

	_, err := toml.DecodeFile(filepath.Join(rootDir, "config/default.toml"), &conf)

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
