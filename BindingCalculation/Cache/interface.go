package cache

import (
	"context"
	"time"
)

// Client 内存缓存的基本接口
type Client interface {
	Set(ctx context.Context, k string, x interface{}, d time.Duration)
	Get(ctx context.Context, k string) (interface{}, bool)
	Replace(ctx context.Context, k string, x interface{}, d time.Duration) error
	Delete(ctx context.Context, k string) (interface{}, bool)
	ItemCount(ctx context.Context) int
	Flush(ctx context.Context)
	DeleteExpired(ctx context.Context)
}
