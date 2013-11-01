package mongo

import "labix.org/v2/mgo"
import "labix.org/v2/mgo/bson"
import "testing"

func check(t *testing.T, err error) {
	if err != nil {
		t.Fail()
	}
}

func testStorage(t *testing.T) *MongoStorage {
	s, err := mgo.Dial("localhost")
	if err != nil {
		t.Fail()
	}
	return NewMongoStorage(s, "test")
}

func TestCreateEntity(t *testing.T) {
	e := bson.M{"name": "Steven"}
	storage := testStorage(t)
	id, resp := storage.Create("tag", e)
	t.Log("id: ", id)
	t.Log("resp: ", resp)
	retrieved, resp := storage.Get("tag", id)
	if "Steven" != retrieved.(bson.M)["name"] {
		t.Fail()
	}
	t.Log("retrieved: ", retrieved)
}

func TestRemoveAll(t *testing.T) {
	storage := testStorage(t)
	storage.DeleteAll("tag")
	all, _ := storage.GetAll("tag")
	if len(all) != 0 {
		t.Fail()
	}
}

func TestGetAll(t *testing.T) {
	TestCreateEntity(t)
	storage := testStorage(t)
	all, _ := storage.GetAll("tag")
	if len(all) != 1 {
		t.Fail()
	}
	if all[0].(bson.M)["name"] != "Steven" {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	e := bson.M{"name": "Steven"}
	storage := testStorage(t)
	id, _ := storage.Create("tag", e)
	resp := storage.Delete("tag", id)
	if resp.StatusCode != 200 {
		t.Fail()
	}
}

func TestDeleteNonExistent(t *testing.T) {
	e := bson.M{"name": "Steven"}
	storage := testStorage(t)
	id, _ := storage.Create("tag", e)
	resp := storage.Delete("tag", id)
	if resp.StatusCode != 200 {
		t.Fail()
	}
	resp = storage.Delete("tag", id)
	if resp.StatusCode != 500 {
		t.Fail()
	}
	t.Log("ErrorMessage: ", resp.ErrorMessage)
}

func TestUpdate(t *testing.T) {
	e := bson.M{"name": "Steven"}
	storage := testStorage(t)
	id, _ := storage.Create("tag", e)
	resp := storage.Update("tag", id, bson.M{"occupation": "hacker"})
	if resp.StatusCode != 200 {
		t.Fail()
	}
  retrieved, _ := storage.Get("tag", id)
  if retrieved.(bson.M)["occupation"] != "hacker" {
    t.Fail()
  }
}
