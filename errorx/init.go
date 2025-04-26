package errorx

import "sync"

var (
	codeUnknown int64  = -1
	msgUnknown  string = "unknown error"
	codeMsgMap         = map[int64]string{
		codeUnknown: msgUnknown,
	}

	codeUnknownLock = sync.RWMutex{}
	msgUnknownLock  = sync.RWMutex{}
	codeMsgMapLock  = sync.RWMutex{}
)

func SetUnknownCode(code int64) {
	codeUnknownLock.Lock()
	codeMsgMapLock.Lock()
	defer codeUnknownLock.Unlock()
	defer codeMsgMapLock.Unlock()

	codeUnknown = code
	codeMsgMap[codeUnknown] = msgUnknown
}

func SetUnknownMsg(msg string) {
	codeUnknownLock.Lock()
	codeMsgMapLock.Lock()
	defer codeUnknownLock.Unlock()
	defer codeMsgMapLock.Unlock()

	msgUnknown = msg
	codeMsgMap[codeUnknown] = msgUnknown
}

func SetCodeMsgMap(newCodeMsgMap map[int64]string) {
	codeMsgMapLock.Lock()
	defer codeMsgMapLock.Unlock()

	for code, msg := range newCodeMsgMap {
		codeMsgMap[code] = msg
	}
}
