package dbbuild

import (
	"context"

	"github.com/openimsdk/tools/log"

	"github.com/openimsdk/open-im-server/v3/pkg/common/config"
	"github.com/openimsdk/tools/db/mongoutil"
	"github.com/openimsdk/tools/db/redisutil"
	"github.com/redis/go-redis/v9"
)

type microservices struct {
	mongo *config.Mongo
	redis *config.Redis
}

func (x *microservices) Mongo(ctx context.Context) (*mongoutil.Client, error) {
	log.ZDebug(ctx, "microservices", "mongo hosts %s", x.mongo.ReplicaSet.Hosts)
	return mongoutil.NewMongoDB(ctx, x.mongo.Build())
}

func (x *microservices) Redis(ctx context.Context) (redis.UniversalClient, error) {
	if x.redis.Disable {
		return nil, nil
	}
	return redisutil.NewRedisClient(ctx, x.redis.Build())
}
