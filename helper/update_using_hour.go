package helper

import (
	"context"
	"graded-challange-1-ysmnababan/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 'transactions' collection contains 'using_hour' entity for describing rental using hour
// by default is empty
// using cron job for updating this field
func UpdateUsingHour(db *mongo.Database) {
	// get all the transactions that do not contain 'usinghour' field

	filter := bson.M{"using_hour": nil}
	cursor, err := db.Collection("transactions").Find(context.TODO(), filter)
	if err != nil {
		Logging(nil).Error(err)
	}
	Logging(nil).Info("BEFORE UPDATE:", cursor)

	for cursor.Next(context.TODO()) {
		//count the using time
		var t models.Transaction
		if err := cursor.Decode(&t); err != nil {
			Logging(nil).Error(err)
			continue
		}

		using_hour := hourDiff(t.StartTime, t.EndTime)
		Logging(nil).Info("USING HOUR:", using_hour)
		if using_hour == 0 {
			continue
		}

		//update to database
		filter = bson.M{"_id": t.ID}
		data := bson.M{"using_hour": using_hour}
		update := bson.M{"$set": data}
		res, err := db.Collection("transactions").UpdateOne(context.TODO(), filter, update)
		if err != nil {
			Logging(nil).Error(err)
			continue
		}
		Logging(nil).Error(res)
	}
}

func hourDiff(startTime, endTime string) float64 {
	// convert to time.Time object
	layout := "2006-01-02 15:04:05"

	// Parse the string according to the layout
	start, err := time.Parse(layout, startTime)
	if err != nil {
		Logging(nil).Error(err)
		return 0
	}

	end, err := time.Parse(layout, endTime)
	if err != nil {
		Logging(nil).Error(err)
		return 0
	}

	return end.Sub(start).Hours()
}
