package logx

import (
	"context"
	"fmt"
)

const (
	keyCtxKVs = "k_kvs"
)

type structCtxKVs struct {
	kvs []any
	pre *structCtxKVs
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

func getAllKVs(ctx context.Context) []any {
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
	kvList := getAllKVs(ctx)
	if kvList != nil {
		kvList = append(kvList, kvs...)
	} else {
		kvList = kvs
	}

	return kvList
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

func CtxDebug(ctx context.Context, template string, args ...any) {
	defaultLogger.Debugw(fmt.Sprintf(template, args...), makeKVs(ctx, nil)...)
}

func CtxInfo(ctx context.Context, template string, args ...any) {
	defaultLogger.Infow(fmt.Sprintf(template, args...), makeKVs(ctx, nil)...)
}

func CtxWarn(ctx context.Context, template string, args ...any) {
	defaultLogger.Warnw(fmt.Sprintf(template, args...), makeKVs(ctx, nil)...)
}

func CtxError(ctx context.Context, template string, args ...any) {
	defaultLogger.Errorw(fmt.Sprintf(template, args...), makeKVs(ctx, nil)...)
}

func CtxPanic(ctx context.Context, template string, args ...any) {
	defaultLogger.Panicw(fmt.Sprintf(template, args...), makeKVs(ctx, nil)...)
}

func CtxDebugKvs(ctx context.Context, msg string, kvs ...any) {
	defaultLogger.Debugw(msg, makeKVs(ctx, kvs)...)
}

func CtxInfoKvs(ctx context.Context, msg string, kvs ...any) {
	defaultLogger.Infow(msg, makeKVs(ctx, kvs)...)
}

func CtxWarnKvs(ctx context.Context, msg string, kvs ...any) {
	defaultLogger.Warnw(msg, makeKVs(ctx, kvs)...)
}

func CtxErrorKvs(ctx context.Context, msg string, kvs ...any) {
	defaultLogger.Errorw(msg, makeKVs(ctx, kvs)...)
}

func CtxPanicKvs(ctx context.Context, msg string, kvs ...any) {
	defaultLogger.Panicw(msg, makeKVs(ctx, kvs)...)
}
