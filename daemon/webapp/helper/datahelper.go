package helper

import (
	"log"

	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	appConfigFile = "conf/app.conf"
)

var (
	Host string = ""
	DB   string = ""
)

func GetDb() (*mgo.Session, error) {
	Host = beego.AppConfig.String("db_server")
	DB = beego.AppConfig.String("db_name")

	uri := "mongodb://" + Host + "/"
	if uri == "" {
		log.Fatal("no connection string provided")
	}
	sess, err := mgo.Dial(uri)
	if err != nil {
		log.Fatal("ERROR When Connecting to host : %s")
	}
	sess.SetSafe(&mgo.Safe{})
	return sess, err
}

func SelectedColumn(columnName ...string) bson.M {
	result := make(bson.M, len(columnName))
	for _, d := range columnName {
		result[d] = 1
	}
	return result

}

func Populate(collectionName string, query map[string]interface{}, column map[string]interface{}, skip int, limit int, sort ...string) ([]bson.D, error) {
	sess, err := GetDb()
	defer sess.Close()
	collection := sess.DB(DB).C(collectionName)
	var result []bson.D
	err = collection.Find(query).Select(column).Skip(skip).Limit(limit).Sort(sort...).All(&result)
	if err != nil {
		log.Panic("ERROR when querying : %s", err)
	}
	return result, err
}

func PopulateAsObject(result interface{}, collectionName string, query map[string]interface{}, column map[string]interface{}, skip int, limit int, sort ...string) {
	sess, err := GetDb()
	defer sess.Close()
	collection := sess.DB(DB).C(collectionName)
	err = collection.Find(query).Select(column).Skip(skip).Limit(limit).Sort(sort...).All(result)

	if err != nil {
		log.Fatalf("ERROR when querying : %s", err)
	}
}

func PopulateOneRow(collectionName string, query map[string]interface{}, column map[string]interface{}) (bson.D, error) {
	sess, err := GetDb()
	defer sess.Close()
	collection := sess.DB(DB).C(collectionName)
	var result bson.D
	err = collection.Find(query).Select(column).One(&result)
	return result, err
}

func Aggregate(collectionName string, pipe interface{}) ([]bson.D, error) {
	var result []bson.D
	sess, err := GetDb()
	defer sess.Close()
	collection := sess.DB(DB).C(collectionName)
	err = collection.Pipe(pipe).AllowDiskUse().All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

func AggregateAsObject(result interface{}, collectionName string, pipe interface{}) {
	sess, err := GetDb()
	defer sess.Close()
	collection := sess.DB(DB).C(collectionName)
	err = collection.Pipe(pipe).AllowDiskUse().All(result)
	if err != nil {
		log.Fatal(err)
	}
}

func Save(collectionName string, docs ...interface{}) {
	sess, _ := GetDb()
	defer sess.Close()
	collection := sess.DB(DB).C(collectionName)
	err := collection.Insert(docs...)
	if err != nil {
		log.Fatal(err)
	}
}

func Update(collectionName string, query map[string]interface{}, update map[string]interface{}) {
	sess, _ := GetDb()
	defer sess.Close()
	collection := sess.DB(DB).C(collectionName)
	_, err := collection.UpdateAll(query, update)
	if err != nil {
		log.Fatal(err)
	}
}

func Delete(collectionName string, query map[string]interface{}) {
	sess, _ := GetDb()
	defer sess.Close()
	collection := sess.DB(DB).C(collectionName)
	_, err := collection.RemoveAll(query)
	if err != nil {
		log.Fatal(err)
	}
}

func Count(collectionName string, query map[string]interface{}) int {
	sess, _ := GetDb()
	defer sess.Close()
	collection := sess.DB(DB).C(collectionName)
	result, err := collection.Find(query).Count()
	if err != nil {
		log.Fatal(err)
	}

	return result
}
