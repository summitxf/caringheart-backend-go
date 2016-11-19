package service

import (
	"../db"
	"../model"

	"gopkg.in/mgo.v2/bson"

	"fmt"

	"time"
)

type WaterService struct {
}

// List GET /water/:duration
func (service WaterService) List(duration int) (result []bson.M, err error) {
	Db := db.MgoDb{}
	Db.Init()

	startDate := time.Now().Add(-time.Duration(duration) * 24 * time.Hour)

	pipeline := []bson.M{
		// {$match: {date: { $gte: ISODate("2016-10-25T09:05:14.688Z") }}},
		bson.M{"$match": bson.M{"date": bson.M{"$gte": startDate}}},
		//{$project: {date: 1, type: 1, amount: 1, inamount: {$cond: {if: {$eq: ['$type', 'in']}, then: '$amount', else: 0}}, outamount: {$cond: {if: {$eq: ['$type', 'out']}, then: '$amount', else: 0}}, _id: 1}}
		bson.M{"$project": bson.M{
			"date":   1,
			"type":   1,
			"amount": 1,
			"inamount": bson.M{
				"$cond": bson.M{
					"if":   bson.M{"$eq": []interface{}{"$type", "in"}},
					"then": "$amount",
					"else": 0,
				}},
			"outamount": bson.M{
				"$cond": bson.M{
					"if":   bson.M{"$eq": []interface{}{"$type", "out"}},
					"then": "$amount",
					"else": 0,
				}},
			"_id": 1,
		}},
		// {$group: {_id: {_id: '$_id', groupdate: {$substr: ['$date', 0, 10]}, date: '$date', type: '$type', amount: '$amount', }, groupinamount: {$sum: '$inamount'}, groupoutamount: {$sum: '$outamount'}}},
		bson.M{"$group": bson.M{
			"_id": bson.M{
				"_id": "$_id",
				"groupdate": bson.M{
					"$substr": []interface{}{"$date", 0, 10},
				},
				"date":   "$date",
				"type":   "$type",
				"amount": "$amount",
			},
			"groupinamount":  bson.M{"$sum": "$inamount"},
			"groupoutamount": bson.M{"$sum": "$outamount"},
		}},
		// {$project: {groupdate: '$_id.groupdate', date: '$_id.date', groupinamount: 1, groupoutamount: 1, info: {_id: '$_id._id', date: '$_id.date', type: '$_id.type', amount: '$_id.amount', }, _id: 0}}
		bson.M{"$project": bson.M{
			"groupdate":      "$_id.groupdate",
			"date":           "$_id.date",
			"groupinamount":  1,
			"groupoutamount": 1,
			"info": bson.M{
				"id":     "$_id._id",
				"date":   "$_id.date",
				"type":   "$_id.type",
				"amount": "$_id.amount",
			},
			"_id": 0,
		}},
		// {$sort: {_id: -1}}
		bson.M{"$sort": bson.M{"date": -1}},
		// {$group: {_id: {date: '$groupdate'}, groupdate: {$max: '$date'}, groupinamount: {$sum: '$groupinamount'}, groupoutamount: {$sum: '$groupoutamount'}, listdata: {$push: '$info'}}}
		bson.M{"$group": bson.M{
			"_id":            bson.M{"date": "$groupdate"},
			"groupdate":      bson.M{"$max": "$date"},
			"groupinamount":  bson.M{"$sum": "$groupinamount"},
			"groupoutamount": bson.M{"$sum": "$groupoutamount"},
			"listdata":       bson.M{"$push": "$info"},
		}},
		// {$sort: {_id: -1}}
		bson.M{"$sort": bson.M{"_id": -1}},
	}

	pipe := Db.C("waterEntity").Pipe(pipeline)
	err = pipe.All(&result)

	defer Db.Close()

	return
}

func (service WaterService) Delete(waterData model.WaterData) (result model.ResultMsg, err error) {
	Db := db.MgoDb{}
	Db.Init()

	if err := Db.C("waterEntity").Remove(waterData); err != nil {
		fmt.Println(err)
		result = model.RsMsg("1")
	} else {
		result = model.RsMsg("0")
	}
	defer Db.Close()

	return
}

func (service WaterService) Save(waterData model.WaterData) (result model.ResultMsg, err error) {
	Db := db.MgoDb{}
	Db.Init()

	waterData.Id = bson.NewObjectId()
	if err := Db.C("waterEntity").Insert(waterData); err != nil {
		fmt.Println(err)
		result = model.RsMsg("1")
	} else {
		result = model.RsMsg("0")
	}
	defer Db.Close()

	return
}
