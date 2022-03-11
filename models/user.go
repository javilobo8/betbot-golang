package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	TwitchId int                `bson:"twitchId" json:"twitchId"`
	UserName string             `bson:"userName" json:"userName"`
}
