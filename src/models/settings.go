package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Settings struct {
	Id         bson.ObjectId `bson:"_id"`
	SessionKey string        `bson:"sessionKey"`
}
