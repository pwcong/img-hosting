package init

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type config struct {
	Server struct {
		Domain string
		Host   string
		Port   int
		Jwt    struct {
			SigningKey string `yaml:"signingKey"`
		}
		Image struct {
			Limit       int
			SupportExts []string `yaml:"supportExts"`
		}
		Middlewares struct {
			Cors struct {
				Active       bool
				AllowOrigins []string `yaml:"allowOrigins"`
				AllowMethods []string `yaml:"allowMethods"`
			}
			Gzip struct {
				Active bool
				Level  int
			}
			Log struct {
				Active bool
				Format string
				Output string
			}
			Limit struct {
				Active bool
				Size   string
			}
		}
	}
	Database struct {
		MySQL struct {
			User     string
			Password string
			Address  string
			DBName   string
		}
	}
}

// DEFAULT_CONFIG is default configuration for img-hosting
const DEFAULT_CONFIG = `
server:
  domain: 'localhost'
  host: 'localhost'
  port: 80
  jwt:
    signingKey: 'img-hosting'
  image:
    limit: 12
    supportExts: ['image/jpeg', 'image/gif', 'image/png', 'image/tiff', 'image/x-icon']
  middlewares:
    cors:
      active: true
      allowOrigins: ['*']
      allowMethods: ['GET', 'HEAD', 'PUT', 'PATCH', 'POST', 'DELETE']
    gzip:
      active: true
      level: 5
    log:
      active: true
      format: '${time_rfc3339_nano} ${remote_ip} ${host} ${method} ${uri} ${status} ${latency_human} ${bytes_in} ${bytes_out}'
      output: 'stdout'
    limit:
      active: true
      size: '4MB'
database:
  mysql:
    user: 'root'
    password: ''
    address: 'localhost:3306'
    dbname: 'img_hosting'
`

var Config config

func initConfig() {

	data, err := ioutil.ReadFile(filepath.Join(filepath.Dir(os.Args[0]), "config", "config.yaml"))

	if err == nil {

		err := yaml.Unmarshal(data, &Config)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			log.Print("has loaded configuration successfully.")
		}

	} else {
		err := yaml.Unmarshal([]byte(DEFAULT_CONFIG), &Config)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			log.Print("failed to load custom configuration, use default.")
		}
	}

}

func init() {
	initConfig()
}
