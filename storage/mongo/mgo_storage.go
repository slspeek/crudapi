package mongo

import "labix.org/v2/mgo"
import "github.com/sauerbraten/crudapi"
import "labix.org/v2/mgo/bson"

type MongoStorage struct {
	session *mgo.Session
	db      string
}

func NewMongoStorage(s *mgo.Session, db string) *MongoStorage {
	return &MongoStorage{s, db}
}

func (self *MongoStorage) Create(kind string, entity interface{}) (string, crudapi.StorageResponse) {
  id := bson.NewObjectId().Hex()
  entity.(map[string]interface{})["_id"] = id
  err := self.collection(kind).Insert(entity)
  if err != nil {
    return "", crudapi.StorageResponse{500, err.Error()}
  }

	return id,  crudapi.StorageResponse{200, ""}
}

func (self *MongoStorage) Get(kind string, id string) (interface{}, crudapi.StorageResponse) {
  var result interface{}
  err := self.collection(kind).FindId(id).One(&result)
  if err != nil {
    return result, crudapi.StorageResponse{404, err.Error()}
  }
	return result, crudapi.StorageResponse{200, ""}
}

func (self *MongoStorage) GetAll(kind string) ([]interface{}, crudapi.StorageResponse) {
  var result []interface{}
  err := self.collection(kind).Find(nil).All(&result)
  if err != nil {
    return result, crudapi.StorageResponse{500, err.Error()}
  }
	return result, crudapi.StorageResponse{200, ""}
}

func (self *MongoStorage) Update(kind string, id string, value interface{}) crudapi.StorageResponse {
  err := self.collection(kind).UpdateId(id, value)
  if err != nil {
    return crudapi.StorageResponse{500, err.Error()}
  }
	return crudapi.StorageResponse{200, ""}
}

func (self *MongoStorage) Delete(kind string, id string) crudapi.StorageResponse {
  err := self.collection(kind).RemoveId(id)
  if err != nil {
    return crudapi.StorageResponse{500, err.Error()}
  }
	return crudapi.StorageResponse{200, ""}
}

func (self *MongoStorage) DeleteAll(kind string) crudapi.StorageResponse {
  _, err := self.collection(kind).RemoveAll(nil)
  if err != nil {
    return crudapi.StorageResponse{500, err.Error()}
  }
	return crudapi.StorageResponse{200, ""}
}

func (self *MongoStorage) collection(kind string) *mgo.Collection {
  return self.session.DB(self.db).C(kind)
}
