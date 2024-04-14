package caches

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/rueidis"
)

type Cache struct {
	client rueidis.Client
}

func (c *Cache) Set(key string, value []byte, ttl time.Duration) error {
	return c.client.Do(context.Background(), c.client.B().Set().Key(key).Value(rueidis.BinaryString(value)).ExSeconds(int64(ttl.Seconds())).Build()).Error()
}

func (c *Cache) Get(key string) (string, error) {
	return c.client.Do(context.Background(), c.client.B().Get().Key(key).Build()).ToString()
}

func Init(url string) (*Cache, error) {
	option := rueidis.ClientOption{
		InitAddress: []string{url},
	}

	cli, err := rueidis.NewClient(option)
	if err != nil {
		return nil, fmt.Errorf("failed to new redis: %w", err)
	}

	ctx := context.Background()

	if err := cli.Do(ctx, cli.B().Ping().Build()).Error(); err != nil {
		return nil, fmt.Errorf("failed to new redis: %w", err)
	}

	return &Cache{
		client: cli,
	}, nil
}
