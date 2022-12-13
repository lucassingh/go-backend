package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*User is the structure of user type*/
type User struct {
	ID       primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Name     string             `bson:"name, omitempty" json:"name"`
	LastName string             `bson:"lastName, omitempty" json:"lastName"`
	DateBorn time.Time          `bson:"dateBorn, omitempty" json:"dateBorn"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password, omitempty" json:"password"`
	Avatar   string             `bson:"avatar, omitempty" json:"avatar"`
	Banner   string             `bson:"banner, omitempty" json:"banner"`
	Bio      string             `bson:"bio, omitempty" json:"bio"`
	Location string             `bson:"location, omitempty" json:"location"`
	WebSite  string             `bson:"webSite, omitempty" json:"webSite"`
}
