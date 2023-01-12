package global

import (
	"CSAwork/model/config"
	"context"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/juju/ratelimit"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var (
	Config        *config.Config
	Logger        *zap.Logger
	GlobalDb1     *sqlx.DB
	RedisDb       *redis.Client
	MongoDBClient *mongo.Client
	Ctx           context.Context
	Bucket        *ratelimit.Bucket
)
