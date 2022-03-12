package services

import (
	"betbot/constants"
	"betbot/models"
	"betbot/util"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	DBClient       mongo.Client
	UserCollection *mongo.Collection
}

type UserRepositoryInterface interface {
	FindByUsername(username string) models.User
	Create(username string, password string) models.User
}

func NewUserRepository(DBClient mongo.Client) UserRepositoryInterface {
	return &UserRepository{
		DBClient:       DBClient,
		UserCollection: DBClient.Database(constants.DBName).Collection("users"),
	}
}

func (repo UserRepository) FindByUsername(username string) models.User {
	var user models.User
	query := bson.D{primitive.E{Key: "username", Value: username}}
	repo.UserCollection.FindOne(context.TODO(), query).Decode(&user)
	return user
}

func (repo UserRepository) Create(username string, password string) models.User {
	passwordHash, _ := util.HashPassword(password)
	user := models.User{
		Username: username,
		Password: passwordHash,
	}
	repo.UserCollection.InsertOne(context.TODO(), user)
	return repo.FindByUsername(username)
}
