package jsonx

import (
	"github.com/bytedance/sonic"
)

type sonicAdapter struct {
	config *config
}

func (a *sonicAdapter) Marshal(v any) ([]byte, error) {
	return a.getSonic(a.config).Marshal(v)
}

func (a *sonicAdapter) Unmarshal(data []byte, v any) error {
	return a.getSonic(a.config).Unmarshal(data, v)
}

func (a *sonicAdapter) MarshalString(v any) (string, error) {
	return a.getSonic(a.config).MarshalToString(v)
}

func (a *sonicAdapter) UnmarshalString(data string, v any) error {
	return a.getSonic(a.config).UnmarshalFromString(data, v)
}

func (a *sonicAdapter) ToString(v any) string {
	str, _ := a.getSonic(a.config).MarshalToString(v)
	return str
}

func (a *sonicAdapter) ToByte(v any) []byte {
	bytes, _ := a.getSonic(a.config).Marshal(v)
	return bytes
}

func (a *sonicAdapter) getSonic(config *config) sonic.API {
	if config.isInit {
		return sonic.ConfigDefault
	}
	return sonic.Config{
		UseNumber:  config.useNumber,
		EscapeHTML: config.escapeHTML,
	}.Froze()
}

func newSonicAdapter(config *config) JsonAdapter {
	return &sonicAdapter{
		config: config,
	}
}
