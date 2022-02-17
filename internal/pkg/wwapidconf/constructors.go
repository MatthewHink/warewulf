package wwapidconf

import (
	"github.com/hpcng/warewulf/internal/pkg/buildconfig"
	"path"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

// WwapidConf is the structure of the wwapid config file.
type WwapidConf struct {
	ApiPrefix string	`yaml:"apiPrefix"`
	ApiVersion string	`yaml:"apiVersion"`
	Port int			`yaml:"port"`
	PublicKey string	`yaml:"publicKey"`
}

// New loads the wwapid config from the given configFilePath or the default if empty.
func New(configFilePath string) (conf WwapidConf, err error) {

	if configFilePath == "" {
		configFilePath = path.Join(buildconfig.SYSCONFDIR(), "warewulf/wwapid.conf")
	}
	log.Printf("Loading wwapid configuration from: %v\n", configFilePath)

	var fileBytes []byte
	fileBytes, err = ioutil.ReadFile(configFilePath)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(fileBytes, &conf)
	if err != nil {
		return
	}

	log.Printf("wwapid config: %#v\n", conf)
	return
}
