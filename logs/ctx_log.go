package logs

import (
	"context"
)

const (
	keyCtxKVs = "k_kvs"
)

type structCtxKVs struct {
	kvs []any
	pre *structCtxKVs
}

func CtxAddKVs(ctx context.Context, kvs ...any) context.Context {
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

func GetAllKVs(ctx context.Context) []any {
	if ctx == nil {
		return nil
	}
	kVsStruct := getCurrentKVsStruct(ctx)
	if kVsStruct == nil {
		return nil
	}

	var result []any
	recursiveAllKVs(&result, kVsStruct, 0)

	return result
}

func recursiveAllKVs(result *[]any, kvsStruct *structCtxKVs, total int) {
	if kvsStruct == nil {
		*result = make([]any, 0, total)
		return
	}
	recursiveAllKVs(result, kvsStruct.pre, total+len(kvsStruct.kvs))
	*result = append(*result, kvsStruct.kvs...)
}

func makeKVs(ctx context.Context, kvs []any) []any {
	kvList := GetAllKVs(ctx)
	if kvList != nil {
		kvList = append(kvList, kvs...)
	} else {
		kvList = kvs
	}

	return kvList
}

func CtxDebugKvs(ctx context.Context, msg string, kvs ...any) {
	logger.Debugw(msg, makeKVs(ctx, kvs)...)
}

func CtxInfoKvs(ctx context.Context, msg string, kvs ...any) {
	logger.Infow(msg, makeKVs(ctx, kvs)...)
}

func CtxWarnKvs(ctx context.Context, msg string, kvs ...any) {
	logger.Warnw(msg, makeKVs(ctx, kvs)...)
}

func CtxErrorKvs(ctx context.Context, msg string, kvs ...any) {
	logger.Errorw(msg, makeKVs(ctx, kvs)...)
}

func CtxPanicKvs(ctx context.Context, msg string, kvs ...any) {
	logger.Panicw(msg, makeKVs(ctx, kvs)...)
}
