package db

import (
	"context"
	"fmt"
	"time"

	"github.com/lucassingh/go-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ReadRelation*/
func ReadRelation(t models.Relation) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := MongoCN.Database("twittor")

	col := db.Collection("relacion")

	condition := bson.M{
		"userid":         t.UserID,
		"userrelationid": t.UserRelation,
	}

	var result models.Relation

	fmt.Println(result)

	err := col.FindOne(ctx, condition).Decode(&result)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
