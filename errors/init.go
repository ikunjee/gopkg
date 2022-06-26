package errors

var (
    codeUnknown int32
    msgUnknown  string
)

var code2msg map[int32]string

func init() {
    codeUnknown = -1
    msgUnknown = "unknown error"
}

func SetUnknownCode(code int32) {
    codeUnknown = code
}

func SetUnknownMsg(msg string) {
    msgUnknown = msg
}

func SetCode2MsgMap(msgMap map[int32]string) {
    code2msg = msgMap
}
