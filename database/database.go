package database

import (
	"gopkg.in/mgo.v2"
)

/*
	Database interface exposing the method necessary to querying, inserting, and updating records
*/
type Database interface {
	Connect(host string) error
	FindOne(collection_name string, query interface{}, result interface{}) error
	FindAll(collection_name string, query interface{}, result interface{}) error
	Insert(collection_name string, item interface{}) error
	Update(collection_name string, selector interface{}, update interface{}) error
}

/*
	MongoDatabase struct which implements the Database interface for a mongo database
*/
type MongoDatabase struct {
	global_session *mgo.Session
	name           string
}

/*
	Initialize connection to mongo database
*/
func InitMongoDatabase(host string, db_name string) (MongoDatabase, error) {
	session, err := mgo.Dial(host)

	db := MongoDatabase{
		global_session: session,
		name:           db_name,
	}

	return db, err
}

/*
	Returns a copy of the global session for use by a connection
*/
func (db MongoDatabase) GetSession() *mgo.Session {
	return db.global_session.Copy()
}

/*
	Find one element matching the given query parameters
*/
func (db MongoDatabase) FindOne(collection_name string, query interface{}, result interface{}) error {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	err := collection.Find(query).One(result)

	return err
}

/*
	Find all elements matching the given query parameters
*/
func (db MongoDatabase) FindAll(collection_name string, query interface{}, result interface{}) error {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	err := collection.Find(query).All(result)

	return err
}

/*
	Insert the given item into the collection
*/
func (db MongoDatabase) Insert(collection_name string, item interface{}) error {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	err := collection.Insert(item)

	return err
}

/*
	Finds an item based on the given selector and updates it with the data in update
*/
func (db MongoDatabase) Update(collection_name string, selector interface{}, update interface{}) error {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	err := collection.Update(selector, update)

	return err
}
