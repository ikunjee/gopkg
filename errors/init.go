package errors

var (
    unknownCode int32
    unknownMsg  string
)

var msgMap map[int32]string

func init() {
    unknownCode = -1
    unknownMsg = "unknown error"
}

func SetUnknownCode(defaultCode int32) {
    unknownCode = defaultCode
}

func SetUnknownMsg(defaultMsg string) {
    unknownMsg = defaultMsg
}

func SetCode2Msg(code2msg map[int32]string) {
    msgMap = code2msg
}
