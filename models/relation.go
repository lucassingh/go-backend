package models

type Relation struct {
	UserID       string `bson:"userid" json:"userId"`
	UserRelation string `bson:"userrelationid" json:"userRelationId"`
}
