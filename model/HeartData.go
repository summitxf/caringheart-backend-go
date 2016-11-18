package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type HeartData struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Date      time.Time     `json:"date" bson:"date"`
	Low       int           `json:"low" bson:"low"`
	High      int           `json:"high" bson:"high"`
	Heartbeat int           `json:"heartbeat" bson:"heartbeat"`
	Weight    float64       `json:"weight" bson:"weight"`
}
