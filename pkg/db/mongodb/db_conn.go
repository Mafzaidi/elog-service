package mongodb

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/mafzaidi/elog/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	DB *mongo.Database
}

var (
	once       sync.Once
	dbInstance *MongoDatabase
	initErr    error
)

func NewMongoDB(conf *config.Config) (*MongoDatabase, error) {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"mongodb://%s:%s",
			conf.DB.Host,
			conf.DB.Port,
		)
		fmt.Println(dsn)
		cred := options.Credential{
			Username:   conf.DB.User,
			Password:   conf.DB.Password,
			AuthSource: conf.DB.DBName,
		}

		clientOpt := options.Client().ApplyURI(dsn).SetAuth(cred)
		client, err := mongo.Connect(context.TODO(), clientOpt)
		if err != nil {
			initErr = err
			return
		}

		if err := client.Ping(context.Background(), nil); err != nil {
			initErr = err
			return
		}

		dbInstance = &MongoDatabase{
			DB: client.Database(conf.DB.DBName),
		}
	})

	if initErr != nil {
		return nil, initErr
	}
	if dbInstance == nil {
		return nil, errors.New("failed to initialize MongoDB")
	}

	return dbInstance, nil
}
