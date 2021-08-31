package modles

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Url struct {
	Id       int    `bson:"id"`
	ShortUrl string `bson:"short_url"`
	Url      string `bson:"url"`
}

func insert(url *Url) error {
	var conf Conf
	uri := conf.getConf().Mongo.Url
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	collection := client.Database(conf.Mongo.DataBase).Collection("short_url_collection")
	_, err = collection.InsertOne(ctx, bson.D{{"_id", url.Id}, {"url", url.Url}, {"short_url", url.ShortUrl}})
	return err
}
