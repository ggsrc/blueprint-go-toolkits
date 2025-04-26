package rediscluster

import (
	"time"
)

type Config struct {
	ClusterAddrs        []string      `required:"true"`
	Password            string        `required:"true"`
	ClusterMaxRedirects int           `default:"3"`
	ReadTimeout         time.Duration `default:"3s"`
	PoolSize            int           `default:"50"`
}
