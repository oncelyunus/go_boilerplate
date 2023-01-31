package injectors

import (
	"github.com/oncelyunus/go_boilerplate/config"
	"github.com/oncelyunus/go_boilerplate/pkg/internal"
	"go.uber.org/zap"
)

func ProvideMongoDB(cfg *config.Config, logger *zap.SugaredLogger) (internal.MongodbConnector, error) {
	return internal.NewMongodbConnector(cfg, logger)
}

func ProvideBaseMongoRepo(config *config.Config,
	mongodbConnector internal.MongodbConnector) *internal.BaseMongoRepo {
	return internal.NewBaseMongoRepo(config, mongodbConnector)
}

