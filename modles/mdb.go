package modles

import (
	//"gopkg.in/mgo.v2"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"io/ioutil"
	"time"
)
import "gopkg.in/yaml.v2"

var (
	//_db 作为单例
	_db  *mongo.Database
	_ctx context.Context
)

type Conf struct {
	Mongo Mongo `yaml:"mongo"`
}

type Mongo struct {
	Url      string `yaml:"url"`
	DataBase string `yaml:"database"`
}

func (conf *Conf) getConf() *Conf {
	yamlFile, err := ioutil.ReadFile("../conf/mongo.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		panic(err)
	}
	return conf
}

//GetDatabase get mongodb database
func GetDatabase() (*mongo.Database, context.Context) {
	if _db != nil {
		return _db, _ctx
	}

	var conf Conf
	uri := conf.getConf().Mongo.Url
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ctx = ctx

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
	database := client.Database(conf.Mongo.DataBase)
	_db = database
	return _db, ctx
}
