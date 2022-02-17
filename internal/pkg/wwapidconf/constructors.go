package wwapidconf

import (
	"github.com/hpcng/warewulf/internal/pkg/buildconfig"
	"path"
	"io/ioutil"
	"fmt" // TODO: log
	"gopkg.in/yaml.v2"
)

// WwapidConf is the structure of the wwapid config file.
type WwapidConf struct {
	ApiPrefix string	`yaml:"apiPrefix"`
	ApiVersion string	`yaml:"apiVersion"`
	Port int			`yaml:"port"`
	PublicKey string	`yaml:"publicKey"`
}

func New() (conf WwapidConf, err error) {
	configFilePath := path.Join(buildconfig.SYSCONFDIR(), "warewulf/wwapid.conf")
	fmt.Printf("configFilePath: %v\n", configFilePath)

	var fileBytes []byte
	fileBytes, err = ioutil.ReadFile(configFilePath)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(fileBytes, &conf)
	if err != nil {
		return
	}

	fmt.Printf("wwapid config: %#v\n", conf)
	return
}
