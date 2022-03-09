package models

type User struct {
	TwitchId int    `json:"twitchId"`
	UserName string `json:"userName"`
}
