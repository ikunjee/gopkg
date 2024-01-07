package jsons

import (
	"github.com/bytedance/sonic"
)

var (
	UnMarshal       = sonic.Unmarshal
	UnMarshalString = sonic.UnmarshalString
)

func ToString(a any) string {
	res, _ := sonic.MarshalString(a)
	return res
}

func Marshal(a any) []byte {
	res, _ := sonic.Marshal(a)
	return res
}
