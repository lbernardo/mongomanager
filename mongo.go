package mongomanager

type MongoManager interface {
	SetDatabase(name string)
	GetItemById(table, id string, result interface{})
	GetOnlyItem(table string, filter interface{}, result interface{})
	GetManyItems(table string, filter interface{}, result interface{})
	InsertItem(table string, item interface{}) (string, error)
	Close()
}