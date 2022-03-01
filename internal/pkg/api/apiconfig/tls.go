package apiconfig

import (
    //"github.com/hpcng/warewulf/internal/pkg/buildconfig"
	//"path"
	//"io/ioutil"
	//"log"
	//"gopkg.in/yaml.v2"
)

// TlsConfig contains TLS configuration parameters for a client or server.
type TlsConfig struct {
	// Enabled is true when secure.
	Enabled bool `yaml:"enabled"`
	// Cert is the path to the client or server certificate file.
	Cert string `yaml:"cert,omitempty"`
	// Key is the path to the client or server key file.
	Key string `yaml:"key,omitempty"`
	// CaCert is the path the CA certificate file.
	CaCert string `yaml:"cacert,omitempty"`
}
