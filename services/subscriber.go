package services

import (
	"context"
	"daringcraft-backend/database"
	"daringcraft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const SubscriberCollection = "subscribers"

func GetSubscribers(filters bson.M, opt *options.FindOptions, query models.Query) models.Result {
	results := make([]models.Subscriber, 0)

	cursor := database.Find(SubscriberCollection, filters, opt)
	count := database.Count(SubscriberCollection, filters)

	if cursor == nil {
		return models.Result{
			Data: results,
		}
	}
	for cursor.Next(context.Background()) {
		var Subscriber models.Subscriber

		if cursor.Decode(&Subscriber) == nil {
			results = append(results, Subscriber)
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

func CreateSubscriber(Subscriber models.Subscriber) (bool, error) {
	_, err := database.InsertOne(SubscriberCollection, Subscriber)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetSubscriber(filter bson.M, opts bson.D) *models.Subscriber {
	option := bson.D{}
	if opts != nil {
		option = opts
	}

	cursor := database.FindOne(SubscriberCollection, filter, options.FindOne().SetProjection(option))

	if cursor == nil {
		return nil
	}

	var Subscriber models.Subscriber

	err := cursor.Decode(&Subscriber)
	if err != nil {
		return nil
	}

	return &Subscriber
}

func UpdateSubscriber(id string, Subscriber models.Subscriber) (*mongo.UpdateResult, error) {
	filters := bson.M{
		"id": id,
	}

	res, err := database.UpdateOne(SubscriberCollection, filters, Subscriber)

	if res == nil {
		return nil, err
	}

	return res, nil
}

func DeleteSubscriber(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"id": id,
	}

	res, err := database.DeleteOne(SubscriberCollection, filter)

	if res == nil {
		return nil, err
	}

	return res, nil
}
