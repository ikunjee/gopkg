package logs

import (
	"context"
)

const (
	keyCtxKVs = "k_kvs"
)

type structCtxKVs struct {
	kvs []interface{}
	pre *structCtxKVs
}

func CtxAddKVs(ctx context.Context, kvs ...interface{}) context.Context {
	if len(kvs) == 0 || (len(kvs)&1 == 1) {
		return ctx
	}

	return context.WithValue(ctx, keyCtxKVs, &structCtxKVs{
		kvs: kvs,
		pre: getCurrentKVsStruct(ctx),
	})
}

func getCurrentKVsStruct(ctx context.Context) *structCtxKVs {
	if ctx == nil {
		return nil
	}
	i := ctx.Value(keyCtxKVs)
	if i == nil {
		return nil
	}
	if kvs, ok := i.(*structCtxKVs); ok {
		return kvs
	}

	return nil
}

func GetAllKVs(ctx context.Context) []interface{} {
	if ctx == nil {
		return nil
	}
	kVsStruct := getCurrentKVsStruct(ctx)
	if kVsStruct == nil {
		return nil
	}

	var result []interface{}
	recursiveAllKVs(&result, kVsStruct, 0)

	return result
}

func recursiveAllKVs(result *[]interface{}, kvsStruct *structCtxKVs, total int) {
	if kvsStruct == nil {
		*result = make([]interface{}, 0, total)
		return
	}
	recursiveAllKVs(result, kvsStruct.pre, total+len(kvsStruct.kvs))
	*result = append(*result, kvsStruct.kvs...)
}

func makeKVs(ctx context.Context, kvs []interface{}) []interface{} {
	kvList := GetAllKVs(ctx)
	if kvList != nil {
		kvList = append(kvList, kvs...)
	} else {
		kvList = kvs
	}

	return kvList
}

func CtxDebugKvs(ctx context.Context, kvs ...interface{}) {
	logger.Debugw("", makeKVs(ctx, kvs)...)
}

func CtxInfoKvs(ctx context.Context, kvs ...interface{}) {
	logger.Infow("", makeKVs(ctx, kvs)...)
}

func CtxWarnKvs(ctx context.Context, kvs ...interface{}) {
	logger.Warnw("", makeKVs(ctx, kvs)...)
}

func CtxErrorKvs(ctx context.Context, kvs ...interface{}) {
	logger.Errorw("", makeKVs(ctx, kvs)...)
}

func CtxPanicKvs(ctx context.Context, kvs ...interface{}) {
	logger.Panicw("", makeKVs(ctx, kvs)...)
}
