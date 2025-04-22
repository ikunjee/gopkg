package asyncx

import (
	"context"
	"log/slog"
	"sync"
	"testing"
)

func TestRecover(t *testing.T) {
	t.Run("recoverWithCtx", func(t *testing.T) {
		ctx := context.Background()

		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer Recover(ctx)
			defer wg.Done()
			panic("test panic")
		}()
		wg.Wait()

		slog.InfoContext(ctx, "recoverWithCtx done")
	})

	t.Run("recoverWithoutCtx", func(t *testing.T) {
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer Recover()
			defer wg.Done()
			panic("test panic")
		}()
		wg.Wait()

		slog.Info("recoverWithoutCtx done")
	})
}
