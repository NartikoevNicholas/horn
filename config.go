package horn

import (
	"github.com/NartikoevNicholas/horn/types"
	"github.com/NartikoevNicholas/horn/utils"
)

type Config struct {
	App             types.App
	Host            string
	Port            int
	Protocol        string
	Transport       string
	WorkerCount     int
	MaxBuffSize     int
	ProtocolHandler types.ProtocolHandler
}

func (c *Config) ApplyDefault() {
	if c.Host == "" {
		c.Host = "127.0.0.1"
	}
	if c.Port == 0 {
		c.Port = 8080
	}
	if c.Protocol == "" {
		c.Protocol = "http"
	}
	if c.Transport == "" {
		c.Transport = utils.TCP
	}
	if c.WorkerCount <= 0 {
		c.WorkerCount = 1
	}
	if c.MaxBuffSize == 0 {
		c.MaxBuffSize = 2048
	}
	if c.ProtocolHandler != (types.ProtocolHandler{}) {
		c.ProtocolHandler = getProtocolHandler(c.Protocol)
	}

}
