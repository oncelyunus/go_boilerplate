package internal

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/oncelyunus/go_boilerplate/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

type MongodbConnector interface {
	DB(context.Context) *mongo.Database
	Client(context.Context) *mongo.Client
	Config() config.MongodbConfig
}

type mongodbConnector struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
	db     *mongo.Database
	client *mongo.Client
}

func NewMongodbConnector(cfg *config.Config, logger *zap.SugaredLogger) (MongodbConnector, error) {
	mongodbConnector := &mongodbConnector{
		cfg:    cfg,
		logger: logger,
	}

	err := mongodbConnector.connect()
	if err != nil {
		return mongodbConnector, err
	}
	return mongodbConnector, nil
}

func (this *mongodbConnector) connect() error {
	var (
		connectOnce sync.Once
		err         error
		client      *mongo.Client
	)

	connectOnce.Do(func() {
		connStr := getConnectionString(&this.cfg.MongoDB)
		client, err = mongo.NewClient(options.Client().ApplyURI(connStr))
		if err != nil {
			this.logger.Errorf("Failed to connect to database: %s, %v", this.cfg.MongoDB.DatabaseName, err)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(this.cfg.MongoDB.DialTimeOut))
		defer cancel()
		err = client.Connect(ctx)
		if err != nil {
			this.logger.Errorf("Failed to connect to database: %s, %v", this.cfg.MongoDB.DatabaseName, err)
			return
		}
	})
	if err != nil {
		return err
	}
	this.client = client
	this.db = this.client.Database(this.cfg.MongoDB.DatabaseName)
	return nil
}

func (this *mongodbConnector) DB(ctx context.Context) *mongo.Database {
	var rp readpref.ReadPref
	err := this.client.Ping(ctx, &rp)
	if err != nil {
		this.logger.Errorf("fail to ping %s, %v", this.cfg.MongoDB.DatabaseHosts, err)
	}
	return this.db
}

func (this *mongodbConnector) Client(ctx context.Context) *mongo.Client {
	return this.client
}

func (this *mongodbConnector) Config() config.MongodbConfig {
	return this.cfg.MongoDB
}

func getConnectionString(config *config.MongodbConfig) string {
	if config.URI != "" {
		return config.URI
	}

	var b bytes.Buffer
	b.WriteString("mongodb://")
	if config.Username != "" {
		b.WriteString(config.Username)
		b.WriteString(":")
	}
	if config.Password != "" {
		b.WriteString(config.Password)
		b.WriteString("@")
	}
	b.WriteString(config.DatabaseHosts)
	b.WriteString("/")

	var urlQueryString []string

	if config.PoolSize != 0 {
		urlQueryString = append(urlQueryString, fmt.Sprintf("maxPoolSize=%d", config.PoolSize))
	}

	if config.ReplicaSet != "" {
		urlQueryString = append(urlQueryString, fmt.Sprintf("replicaSet=%s", config.ReplicaSet))
	}

	if config.AuthSource != "" {
		urlQueryString = append(urlQueryString, fmt.Sprintf("authSource=%s", config.AuthSource))
	}

	if len(urlQueryString) > 0 {
		s := strings.Join(urlQueryString, "&")
		s = "?" + s
		b.WriteString(s)
	}

	return b.String()
}
