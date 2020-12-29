package mongomanager

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type DatabaseManager struct {
	client *mongo.Client
	database *mongo.Database
}

func NewDatabaseManager(databaseURI string) (MongoManager, error) {
	client, _ := mongo.NewClient(options.Client().ApplyURI(databaseURI))
	err := client.Connect(context.Background())
	if err != nil {
		log.Println("MongoManager error to connected database", err)
		return nil, err
	}
	return &DatabaseManager{
		client: client,
	}, nil
}

func (d *DatabaseManager) SetDatabase(name string) {
	d.database = d.client.Database(name)
}

func (d *DatabaseManager) GetItemById(table, id string, result interface{}) {
	idB,_ := primitive.ObjectIDFromHex(id)
	d.database.Collection(table).FindOne(context.Background(), bson.M{"_id": idB}).Decode(result)
}

func (d *DatabaseManager) GetOnlyItem(table string, filter interface{}, result interface{}) {
	if filter == nil {
		filter = bson.D{}
	}
	d.database.Collection(table).FindOne(context.Background(), filter).Decode(result)
}

func (d *DatabaseManager) GetManyItems(table string, filter interface{}, result interface{}) {
	if filter == nil {
		filter = bson.D{}
	}
	c, _ := d.database.Collection(table).Find(context.Background(), filter)
	c.All(context.Background(), result)
}

func (d *DatabaseManager) InsertItem(table string, item interface{}) (string, error) {
	res, err := d.database.Collection(table).InsertOne(context.Background(), item)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (d *DatabaseManager) Close() {
	d.client.Disconnect(context.Background())
}
