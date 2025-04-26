package rediscluster

import (
	"context"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type Client struct {
	redis.UniversalClient
}

func NewRedisClusterClient(envPrefix string) *Client {
	c := Config{}
	envconfig.MustProcess(envPrefix, &c)
	log.Debug().Msgf("Redis Config: %+v", c)

	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        c.ClusterAddrs,
		MaxRedirects: c.ClusterMaxRedirects,
		ReadTimeout:  c.ReadTimeout,
		PoolSize:     c.PoolSize,
		Password:     c.Password,
		TLSConfig:    nil,
	})

	// ping to check if the connection is successful
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatal().Err(err).Msg("failed to establish connection with Redis cluster")
	}

	log.Info().Msg("Successfully connected to Redis cluster")

	if err := redisotel.InstrumentTracing(client); err != nil {
		log.Fatal().Err(err).Msg("failed to instrument OpenTelemetry tracing for Redis cluster")
	}
	if err := redisotel.InstrumentMetrics(client); err != nil {
		log.Fatal().Err(err).Msg("failed to instrument OpenTelemetry metrics for Redis cluster")
	}

	return &Client{client}
}
