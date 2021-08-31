package modles

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
	"time"
)

func TestGetDatabase(t *testing.T) {
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
	collection := client.Database(conf.Mongo.DataBase).Collection("idGen")
	fmt.Println(collection)
	//one, err := collection.InsertOne(ctx, bson.D{{"_id","url"},{"maxId",0}})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(one.InsertedID)
	collection.UpdateOne(ctx, bson.D{{"_id", "url"}}, bson.D{{"$inc", bson.D{{"maxId", 1}}}})
}
