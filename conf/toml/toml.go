package toml

import (
	"bytes"
	"github.com/pelletier/go-toml"
)

type RawConfig map[string]interface{}

type TOML struct {
}

func Parser() *TOML {
	return &TOML{}
}
func (p *TOML) Unmarshal(b []byte) (RawConfig, error) {
	r, err := toml.LoadReader(bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	return r.ToMap(), err
}
func (p *TOML) Marshal(o RawConfig) ([]byte, error) {
	out, err := toml.TreeFromMap(o)
	if err != nil {
		return nil, err
	}
	return out.Marshal()
}
