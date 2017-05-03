package init

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type config struct {
	Mode   string
	Server struct {
		Domain        string
		Host          string
		Port          int
		CompressLevel int
	}
}

const DEFAULT_CONFIG = `
mode: prod
server:
  domain: pwcong.me
  host: localhost
  port: 80
  compressLevel: 5
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
