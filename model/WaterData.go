package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type WaterData struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Date   time.Time     `json:"date" bson:"date"`
	Type   string        `json:"type" bson:"type"`
	Amount int           `json:"amount" bson:"amount"`
}
