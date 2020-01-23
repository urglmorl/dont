package store

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"models"
	"utils"
)

var session *mgo.Session

const databaseName = "dont"
const databaseUrl = "mongodb://127.0.0.1"

var tables Tables

type Tables struct {
	subjects string
	tests    string
	settings string
}

func Init() {
	tables = initTables()
	session = openSession()
}

func openSession() (session *mgo.Session) {
	session, err := mgo.Dial(databaseUrl)
	if err != nil {
		session = nil
		panic(err)
	}
	return
}

func initTables() (tables Tables) {
	tables.subjects = "subjects"
	tables.tests = "tests"
	return
}

func GetSubjects() (subjects models.SubjectsBson) {
	collection := session.DB(databaseName).C(tables.subjects)
	query := bson.M{}
	err := collection.Find(query).All(&subjects)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func CreateSubject(subject models.SubjectBson) error {
	subject.Id = utils.UUIDs().String()
	collection := session.DB(databaseName).C(tables.subjects)
	return collection.Insert(subject)
}
