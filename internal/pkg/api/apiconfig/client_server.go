package apiconfig

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

// ClientServerConfig is the full client server configuration.
type ClientServerConfig struct {
	ClientApiConfig ClientApiConfig `yaml:"clientapi"`
	ServerApiConfig ServerApiConfig `yaml:"serverapi"`
	ClientTlsConfig TlsConfig `yaml:"clienttls"`
	ServerTlsConfig TlsConfig `yaml:"servertls"`
}

// NewClientServer loads the client config from the given configFilePath.
// TODO: parameters will be oneof these:
// path.Join(buildconfig.SYSCONFDIR(), "warewulf/wwapic.conf")
func NewClientServer(configFilePath string) (config ClientServerConfig, err error) {

	log.Printf("Loading api client server configuration from: %v\n", configFilePath)

	var fileBytes []byte
	fileBytes, err = ioutil.ReadFile(configFilePath)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(fileBytes, &config)
	if err != nil {
		return
	}

	log.Printf("api client server config: %#v\n", config)
	return
}