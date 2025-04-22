package jsonx

import (
	"encoding/json"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

var (
	validJsonStr = `{"int": 1,"int64":8888888888888888888,"intList":[1,2,3],"str":"hello world","bool":true,"null":null,"obj":{"int":2}, "html": "</>"}`
	validJsonMap = map[string]any{
		"int":     1,
		"int64":   int64Number,
		"intList": []int{1, 2, 3},
		"str":     "hello world",
		"bool":    true,
		"null":    nil,
		"obj": map[string]any{
			"int": 2,
		},
		"html": "</>",
	}
	invalidJsonStr       = "invalid"
	nullJsonStr          = "null"
	emptyStr             = ""
	int64Number    int64 = 8888888888888888888
	int64NumberStr       = "8888888888888888888"
)

func TestDefault(t *testing.T) {
	convey.Convey("ToString", t, func() {
		convey.Convey("null", func() {
			var obj map[string]any
			str := ToString(obj)
			convey.So(str, convey.ShouldEqual, "null")
		})
		convey.Convey("string to string", func() {
			convey.So(ToString("string"), convey.ShouldEqual, "\"string\"")
		})
		convey.Convey("int64 to string", func() {
			convey.So(ToString(int64Number), convey.ShouldEqual, int64NumberStr)
		})
		convey.Convey("html to string", func() {
			convey.So(ToString("</>"), convey.ShouldEqual, "\"</>\"")
		})
		convey.Convey("map", func() {
			str := ToString(validJsonMap)
			t.Log(str) // 结果是无序的，不做数据 assert
			convey.So(len(str), convey.ShouldBeGreaterThan, 0)
		})
	})

	convey.Convey("UnmarshalStringWithType", t, func() {
		convey.Convey("empty str", func() {
			obj, err := UnmarshalStringWithType[map[string]any](emptyStr)
			convey.So(err, convey.ShouldBeNil)
			convey.So(len(obj), convey.ShouldEqual, 0)
		})
		convey.Convey("null json", func() {
			obj, err := UnmarshalStringWithType[map[string]any](nullJsonStr)
			convey.So(err, convey.ShouldBeNil)
			convey.So(len(obj), convey.ShouldEqual, 0)
		})
		convey.Convey("invalid str", func() {
			obj, err := UnmarshalStringWithType[map[string]any](invalidJsonStr)
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(len(obj), convey.ShouldEqual, 0)
		})
		convey.Convey("valid str", func() {
			obj, err := UnmarshalStringWithType[map[string]any](validJsonStr)
			convey.So(err, convey.ShouldBeNil)
			convey.So(len(obj), convey.ShouldBeGreaterThan, 0)
			numberInt64Any := obj["int64"]
			numberInt64, ok := numberInt64Any.(float64) // 数字默认 unmarshal 为 float64
			convey.So(ok, convey.ShouldBeTrue)
			convey.So(numberInt64, convey.ShouldEqual, float64(8888888888888888888))
			numberIntAny := obj["int"]
			numberInt, ok := numberIntAny.(float64)
			convey.So(ok, convey.ShouldBeTrue)
			convey.So(numberInt, convey.ShouldEqual, float64(1))
			convey.So(obj["intList"], convey.ShouldHaveLength, 3)
			convey.So(obj["str"], convey.ShouldEqual, "hello world")
			convey.So(obj["bool"], convey.ShouldEqual, true)
			convey.So(obj["null"], convey.ShouldBeNil)
			convey.So(obj["obj"], convey.ShouldHaveLength, 1)
		})
	})
}

func TestConfigDefault(t *testing.T) {
	ConfigDefault(WithEscapeHTML(true), WithUseNumber(true))

	convey.Convey("ToString", t, func() {
		convey.Convey("html to string", func() {
			convey.So(ToString("</>"), convey.ShouldEqual, "\"\\u003c/\\u003e\"")
		})
	})

	convey.Convey("UnmarshalStringWithType", t, func() {
		convey.Convey("valid str", func() {
			obj, err := UnmarshalStringWithType[map[string]any](validJsonStr)
			convey.So(err, convey.ShouldBeNil)
			convey.So(len(obj), convey.ShouldBeGreaterThan, 0)
			numberInt64Any := obj["int64"]
			numberInt64, ok := numberInt64Any.(json.Number) // use number
			convey.So(ok, convey.ShouldBeTrue)
			convey.So(numberInt64.String(), convey.ShouldEqual, int64NumberStr)
			i, err := numberInt64.Int64()
			convey.So(err, convey.ShouldBeNil)
			convey.So(i, convey.ShouldEqual, int64Number)
		})
	})
}

func TestConfigOpt(t *testing.T) {
	convey.Convey("ToString", t, func() {
		convey.Convey("html to string", func() {
			convey.So(ToString("</>", WithEscapeHTML(true)), convey.ShouldEqual, "\"\\u003c/\\u003e\"")
		})
	})

	convey.Convey("UnmarshalStringWithType", t, func() {
		convey.Convey("valid str", func() {
			obj, err := UnmarshalStringWithType[map[string]any](validJsonStr, WithUseNumber(true))
			convey.So(err, convey.ShouldBeNil)
			convey.So(len(obj), convey.ShouldBeGreaterThan, 0)
			numberInt64Any := obj["int64"]
			numberInt64, ok := numberInt64Any.(json.Number) // use number
			convey.So(ok, convey.ShouldBeTrue)
			convey.So(numberInt64.String(), convey.ShouldEqual, int64NumberStr)
			i, err := numberInt64.Int64()
			convey.So(err, convey.ShouldBeNil)
			convey.So(i, convey.ShouldEqual, int64Number)
		})
	})
}
