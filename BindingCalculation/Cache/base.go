package cache

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// BaseClient 最基本的结构体,只实现了基本的方法，后续所有的结构体都通过装饰者的模式持有此对象及其拓展对象.
type BaseClient struct {
	items map[string]Item
}

type Item struct {
	Object     interface{}
	Expiration int64
}

// New 初始化方法
func New() *BaseClient {
	return &BaseClient{
		items: make(map[string]Item),
	}
}

// Set 所有方法均使用全部采用拷贝指针的方式传递对象
func (b *BaseClient) Set(ctx context.Context, key string, value interface{}, d time.Duration) {
	b.items[key] = Item{
		Object:     value,
		Expiration: int64(d),
	}
}

// Get bool标志有没有找到对象
func (b *BaseClient) Get(ctx context.Context, k string) (interface{}, bool) {
	item, found := b.items[k]
	if !found {
		return nil, false
	}
	return item.Object, true
}

// Replace 手动更新对象
func (b *BaseClient) Replace(ctx context.Context, k string, x interface{}, d time.Duration) error {
	_, found := b.Get(ctx, k)
	if !found {
		return fmt.Errorf("item %s doesn't exist", k)
	}
	b.Set(ctx, k, x, d)
	return nil
}

// Delete 从缓存中删除项目。如果key不在缓存中，则不执行任何操作。
func (b *BaseClient) Delete(ctx context.Context, k string) (interface{}, bool) {
	if v, found := b.items[k]; found {
		delete(b.items, k)
		return v.Object, true
	}
	return nil, false
}

// ItemCount 返回map中元素个数
func (b *BaseClient) ItemCount(ctx context.Context) int {
	n := len(b.items)
	return n
}

// Flush 刷新缓存
func (b *BaseClient) Flush(ctx context.Context) {
	b.items = map[string]Item{}
}

// DeleteExpired 删除过期缓存，初始不做任何操作
func (b *BaseClient) DeleteExpired(ctx context.Context) {
	return
}

func (b *BaseClient) RandomKey() string {
	keys := make([]string, 0, len(b.items))
	for key := range b.items {
		fmt.Printf(key)
		keys = append(keys, key)
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(keys))
	return keys[randomIndex]
}
