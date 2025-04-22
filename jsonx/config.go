package jsonx

import (
	"sync"
)

type jsonLibType int32

const (
	JsonLibTypeSonic jsonLibType = iota
	JsonLibTypeEncodingJson
)

type config struct {
	jsonLibType jsonLibType // 底层使用的 json 库

	// 公共 json 行为控制
	escapeHTML bool // marshal 时是否转义 html 字符
	useNumber  bool // unmarshal 时对于数字类型是否转成 json.number 类型而不是 float64，避免精度丢失

	isInit bool // 是否初始配置
}

var (
	defaultConfig = &config{
		jsonLibType: JsonLibTypeSonic,
		useNumber:   false,
		escapeHTML:  false,
		isInit:      true,
	}
	defaultConfigLock = sync.RWMutex{}
)

func (c *config) Clone() *config {
	return &config{
		jsonLibType: c.jsonLibType,
		useNumber:   c.useNumber,
		escapeHTML:  c.escapeHTML,
		isInit:      false,
	}
}

func configDefault(opts ...Option) {
	defaultConfigLock.Lock()
	defer defaultConfigLock.Unlock()
	for _, opt := range opts {
		opt(defaultConfig)
	}
	defaultConfig.isInit = false
}

func getConfig(opts ...Option) *config {
	if len(opts) == 0 {
		return defaultConfig
	}
	config := defaultConfig.Clone()
	for _, opt := range opts {
		opt(config)
	}
	return config
}
