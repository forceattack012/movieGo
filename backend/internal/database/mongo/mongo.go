package mongo

import (
	"backend/internal/database/mongo/seeds"
	"bytes"
	"context"
	"time"

	"github.com/naamancurtis/mongo-go-struct-to-bson/mapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Host         string
	Port         string
	DatabaseName string
}

func BuildConfig(host string, port string, databaseName string) *MongoConfig {
	return &MongoConfig{
		Host:         host,
		Port:         port,
		DatabaseName: databaseName,
	}
}

func Connect(mongoConfig *MongoConfig) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var b bytes.Buffer
	b.WriteString(mongoConfig.Host)
	b.WriteString(mongoConfig.Port)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(b.String()))

	if err != nil {
		return nil, err
	}

	return client.Database(mongoConfig.DatabaseName), err
}

func Seed(db *mongo.Database, collection string) {
	movies := seeds.GetMovies()

	for _, m := range movies {
		result := mapper.ConvertStructToBSONMap(m, nil)
		db.Collection(collection).InsertOne(context.Background(), result)
	}
}

func DeleteSeed(db *mongo.Database, collection string) {
	movies := seeds.GetMovies()

	for _, m := range movies {
		filter := bson.D{primitive.E{Key: "Name", Value: m.Name}}
		db.Collection(collection).DeleteMany(context.Background(), filter)
	}
}
