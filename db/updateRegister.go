package db

import (
	"context"
	"time"

	"github.com/lucassingh/go-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*UpdateRegister*/
func UpdateRegister(u models.User, ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twittor")

	col := db.Collection("usuarios")

	register := make(map[string]interface{})

	if len(u.Name) > 0 {
		register["name"] = u.Name
	}

	if len(u.LastName) > 0 {
		register["lastName"] = u.LastName
	}

	register["dateBorn"] = u.DateBorn

	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}

	if len(u.Bio) > 0 {
		register["bio"] = u.Bio
	}

	if len(u.Location) > 0 {
		register["location"] = u.Location
	}

	if len(u.WebSite) > 0 {
		register["webSite"] = u.WebSite
	}

	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}

	if len(u.Email) > 0 {
		register["email"] = u.Email
	}

	uptdateString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, uptdateString)

	if err != nil {
		return false, err
	}

	return true, nil
}
