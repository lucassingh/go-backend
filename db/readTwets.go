package db

import (
	"context"
	"log"
	"time"

	"github.com/lucassingh/go-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ReadTweets*/
func ReadTweets(ID string, page int64) ([]*models.ResponseTweets, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twittor")

	col := db.Collection("tweet")

	var results []*models.ResponseTweets

	condition := bson.M{
		"userid": ID,
	}

	variants := options.Find()

	variants.SetLimit(20)

	variants.SetSort(bson.D{{Key: "date", Value: -1}})

	variants.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, variants)

	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {

		var register models.ResponseTweets

		err := cursor.Decode(&register)

		if err != nil {
			return results, false
		}

		results = append(results, &register)
	}

	return results, true
}
