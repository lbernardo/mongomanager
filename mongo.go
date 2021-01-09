package mongomanager

import "go.mongodb.org/mongo-driver/bson/primitive"

type MongoManager interface {
	SetDatabase(name string)
	GetItemById(table, id string, result interface{})
	GetOnlyItem(table string, filter interface{}, result interface{})
	GetManyItems(table string, filter interface{}, result interface{})
	InsertItem(table string, item interface{}) (string, error)
	DeleteItem(table string, id primitive.ObjectID) error
	UpdateItem(table string, id primitive.ObjectID, item interface{}) error
	CreateFilterWithObjectID(name string, value primitive.ObjectID) interface{}
	Close()
}