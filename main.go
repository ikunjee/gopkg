package main

import (
	"fmt"

	"github.com/ikunjee/gopkg/jsonx"
)

func main() {
	fmt.Println(jsonx.ToString(map[string]string{"a": "123"}))
	fmt.Println(jsonx.ToString("</>"))
	fmt.Println(jsonx.ToString("</>", jsonx.WithEscapeHTML(true)))
	data, _ := jsonx.UnmarshalStringWithType[map[string]any]("{\"a\":8888888888888888888}", jsonx.WithUseNumber(true))
	fmt.Println(jsonx.ToString(data))
	data, _ = jsonx.UnmarshalStringWithType[map[string]any]("{\"a\":8888888888888888888}")
	fmt.Println(jsonx.ToString(data))
	jsonx.ConfigDefault(jsonx.WithEscapeHTML(true), jsonx.WithUseNumber(true))
	fmt.Println(jsonx.ToString("</>"))
	data, _ = jsonx.UnmarshalStringWithType[map[string]any]("{\"a\":8888888888888888888}")
	fmt.Println(jsonx.ToString(data))
}
