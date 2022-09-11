package services

import (
	"context"
	"daringcraft-backend/database"
	"daringcraft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MessageCollection = "messages"

func GetMessages(filters bson.M, opt *options.FindOptions, query models.Query) models.Result {
	results := make([]models.Message, 0)

	cursor := database.Find(MessageCollection, filters, opt)
	count := database.Count(MessageCollection, filters)

	if cursor == nil {
		return models.Result{
			Data: results,
		}
	}
	for cursor.Next(context.Background()) {
		var Message models.Message

		if cursor.Decode(&Message) == nil {
			results = append(results, Message)
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

func CreateMessage(Message models.Message) (bool, error) {
	_, err := database.InsertOne(MessageCollection, Message)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetMessage(filter bson.M, opts bson.D) *models.Message {
	option := bson.D{}
	if opts != nil {
		option = opts
	}

	cursor := database.FindOne(MessageCollection, filter, options.FindOne().SetProjection(option))

	if cursor == nil {
		return nil
	}

	var Message models.Message

	err := cursor.Decode(&Message)
	if err != nil {
		return nil
	}

	return &Message
}

func UpdateMessage(id string, Message models.Message) (*mongo.UpdateResult, error) {
	filters := bson.M{
		"id": id,
	}

	res, err := database.UpdateOne(MessageCollection, filters, Message)

	if res == nil {
		return nil, err
	}

	return res, nil
}

func DeleteMessage(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"id": id,
	}

	res, err := database.DeleteOne(MessageCollection, filter)

	if res == nil {
		return nil, err
	}

	return res, nil
}
