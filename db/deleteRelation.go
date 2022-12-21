package db

import (
	"context"
	"time"

	"github.com/lucassingh/go-backend/models"
)

/*DeleteRelation*/
func DeleteRelation(t models.Relation) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := MongoCN.Database("twittor")

	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}

	return true, nil
}
