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
		Domain      string
		Host        string
		Port        int
		Middlewares struct {
			Cors struct {
				Active       bool
				AllowOrigins []string
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
		}
	}
}

const DEFAULT_CONFIG = `
server:
  domain: pwcong.me
  host: localhost
  port: 80

  middlewares:
    cors:
      active: true
      allowOrigins: ['*']
    gzip:
      active: true
      level: 5
    log:
      active: true
      format: '${time_rfc3339_nano} ${remote_ip} ${host} ${method} ${uri} ${status} ${latency_human} ${bytes_in} ${bytes_out}'
      output: stdout
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
