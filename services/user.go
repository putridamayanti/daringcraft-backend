package services

import (
	"context"
	"daringcraft-backend/database"
	"daringcraft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const UserCollection = "users"

func GetUsers(filters bson.M, opt *options.FindOptions, query models.Query) models.Result {
	results := make([]models.User, 0)

	cursor := database.Find(UserCollection, filters, opt)
	count := database.Count(UserCollection, filters)

	if cursor == nil {
		return models.Result{
			Data: results,
		}
	}
	for cursor.Next(context.Background()) {
		var user models.User

		if cursor.Decode(&user) == nil {
			results = append(results, user)
		}
	}

	pagination := query.GetPagination(count)

	result := models.Result{
		Data: results,
		Pagination: pagination,
		Query: query,
	}

	return result
}

func CreateUser(user models.User) (bool, error) {
	_, err := database.InsertOne(UserCollection, user)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetUser(filter bson.M, opts bson.D) *models.User {
	option := bson.D{}
	if opts != nil {
		option = opts
	}

	cursor := database.FindOne(UserCollection, filter, options.FindOne().SetProjection(option))

	if cursor == nil {
		return nil
	}

	var user models.User

	err := cursor.Decode(&user)
	if err != nil {
		return nil
	}

	return &user
}

func GetUserByEmail(email string) *models.User {
	cursor := database.FindOne(UserCollection, bson.M{"email": email}, nil)
	if cursor == nil {
		return nil
	}

	var user models.User
	err := cursor.Decode(&user)
	if err != nil {
		return nil
	}

	user.Password = ""

	return &user
}

func UpdateUser(id string, user models.User) (*mongo.UpdateResult, error) {
	filters := bson.M{
		"id": id,
	}

	res, err := database.UpdateOne(UserCollection, filters, user)

	if res == nil {
		return nil, err
	}

	return res, nil
}

func DeleteUser(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"id": id,
	}

	res, err := database.DeleteOne(UserCollection, filter)

	if res == nil {
		return nil, err
	}

	return res, nil
}
