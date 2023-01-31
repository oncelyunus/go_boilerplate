package internal

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/oncelyunus/go_boilerplate/config"
)

type BaseMongoRepo struct {
	Config           *config.Config
	MongodbConnector MongodbConnector
}

func NewBaseMongoRepo(cfg *config.Config, mongodbConnector MongodbConnector) *BaseMongoRepo {
	base := &BaseMongoRepo{
		Config:           cfg,
		MongodbConnector: mongodbConnector,
	}

	return base
}

func (base *BaseMongoRepo) GenerateID(ctx context.Context) string {
	return base.Config.MongoDB.CustomIDPrefix + "-" + strings.ToUpper(strconv.FormatInt(time.Now().UnixNano(), 36))
}
