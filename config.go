package proxy

import (
	"github.com/BurntSushi/toml"
	"io"
	"net"
)

type Host struct {
	Backends map[string]*Backend `toml:"backends"`
	Dns      string              `toml:"dns"`
	Log      string              `toml:"log"`
}

type Backend struct {
	Proto            string `toml:"proto"`
	IP               net.IP `toml:"ip"`
	Port             int    `toml:"port"`
	Query            string `toml:"query"`
	Records          string `toml:"records"`
	MaxConcurrent    int    `toml:"max_concurrent"`
	ConnectionBuffer int    `toml:"connection_buffer"`
}

func LoadConfig(r io.Reader) (*Host, error) {
	var config *Host
	if _, err := toml.DecodeReader(r, &config); err != nil {
		return nil, err
	}
	return config, nil
}
