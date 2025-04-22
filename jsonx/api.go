package jsonx

func Marshal(v any, opts ...Option) ([]byte, error) {
	return getAdapter(getConfig(opts...)).Marshal(v)
}

func Unmarshal(data []byte, v any, opts ...Option) error {
	return getAdapter(getConfig(opts...)).Unmarshal(data, v)
}

// ToString 序列化为 string，忽略错误
func ToString(a any, opts ...Option) string {
	return getAdapter(getConfig(opts...)).ToString(a)
}

// ToByte 序列化为 bytes，忽略错误
func ToByte(a any, opts ...Option) []byte {
	return getAdapter(getConfig(opts...)).ToByte(a)
}

func MarshalString(a any, opts ...Option) (string, error) {
	return getAdapter(getConfig(opts...)).MarshalString(a)
}

func UnmarshalString(str string, v any, opts ...Option) error {
	return getAdapter(getConfig(opts...)).UnmarshalString(str, v)
}

// UnmarshalStringWithType 反序列化 string 到指定类型，传入空字符串时返回零值而不是报错
func UnmarshalStringWithType[T any](str string, opts ...Option) (T, error) {
	var t T
	if str == "" {
		return t, nil
	}
	err := UnmarshalString(str, &t, opts...)
	return t, err
}

// UnmarshalWithType 反序列化 bytes 到指定类型，传入空 bytes 时返回零值而不是报错
func UnmarshalWithType[T any](bytes []byte, opts ...Option) (T, error) {
	var t T
	if len(bytes) == 0 {
		return t, nil
	}
	err := Unmarshal(bytes, &t, opts...)
	return t, err
}

func ConfigDefault(opts ...Option) {
	configDefault(opts...)
}
