package models

/*Tweet*/
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
