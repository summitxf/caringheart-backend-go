package service

import (
	"../db"
	"../model"

	"gopkg.in/mgo.v2/bson"

	"fmt"

	"time"
)

type HeartService struct {
}

// List GET /heart/:duration
func (service HeartService) List(duration int) (result []bson.M, err error) {
	Db := db.MgoDb{}
	Db.Init()

	startDate := time.Now().Add(-time.Duration(duration) * 24 * time.Hour)

	pipeline := []bson.M{
		// {$match: {date: { $gte: ISODate("2016-10-25T09:05:14.688Z") }}},
		bson.M{"$match": bson.M{"date": bson.M{"$gte": startDate}}},
		// {$group: {_id: {_id: '$_id', groupdate: {$substr: ['$date', 0, 10]}, date: '$date', low: '$low', high: '$high', heartbeat: '$heartbeat', weight: '$weight' }}},
		bson.M{"$group": bson.M{
			"_id": bson.M{
				"_id": "$_id",
				"groupdate": bson.M{
					"$substr": []interface{}{"$date", 0, 10},
				},
				"date":      "$date",
				"low":       "$low",
				"high":      "$high",
				"heartbeat": "$heartbeat",
				"weight":    "$weight",
			},
		}},
		// {$project: {groupdate: '$_id.groupdate', date: '$_id.date', info: {_id: '$_id._id', date: '$_id.date', low: '$_id.low', high: '$_id.high', heartbeat: '$_id.heartbeat', weight: '$_id.weight' }, _id: 0}},
		bson.M{"$project": bson.M{
			"groupdate": "$_id.groupdate",
			"date":      "$_id.date",
			"info": bson.M{
				"id":        "$_id._id",
				"date":      "$_id.date",
				"low":       "$_id.low",
				"high":      "$_id.high",
				"heartbeat": "$_id.heartbeat",
				"weight":    "$_id.weight",
			},
			"_id": 0,
		}},
		// {$group: {_id: {date: '$groupdate'}, groupdate: {$max: '$date'}, listdata: {$push: '$info'}}},
		bson.M{"$group": bson.M{
			"_id":       bson.M{"date": "$groupdate"},
			"groupdate": bson.M{"$max": "$date"},
			"listdata":  bson.M{"$push": "$info"},
		}},
		// {$sort: {_id: -1}}
		bson.M{"$sort": bson.M{"_id": -1}},
	}

	pipe := Db.C("heartEntity").Pipe(pipeline)
	err = pipe.All(&result)

	defer Db.Close()

	return
}

func (service HeartService) Delete(heartData model.HeartData) (result model.ResultMsg, err error) {
	Db := db.MgoDb{}
	Db.Init()

	if err := Db.C("heartEntity").Remove(heartData); err != nil {
		fmt.Println(err)
		result = model.RsMsg("1")
	} else {
		result = model.RsMsg("0")
	}
	defer Db.Close()

	return
}

func (service HeartService) Save(heartData model.HeartData) (result model.ResultMsg, err error) {
	Db := db.MgoDb{}
	Db.Init()

	heartData.Id = bson.NewObjectId()
	if err := Db.C("heartEntity").Insert(heartData); err != nil {
		fmt.Println(err)
		result = model.RsMsg("1")
	} else {
		result = model.RsMsg("0")
	}
	defer Db.Close()

	return
}
