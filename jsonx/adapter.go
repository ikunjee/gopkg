package jsonx

type JsonAdapter interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
	MarshalString(v any) (string, error)
	UnmarshalString(data string, v any) error
	ToString(v any) string
	ToByte(v any) []byte
}

var libToAdapterMap = map[jsonLibType]func(config *config) JsonAdapter{
	JsonLibTypeSonic:        newSonicAdapter,
	JsonLibTypeEncodingJson: newEncodingJsonAdapter,
}

func getAdapter(config *config) JsonAdapter {
	adapterFunc, ok := libToAdapterMap[config.jsonLibType]
	if !ok {
		return newSonicAdapter(config)
	}
	return adapterFunc(config)
}
