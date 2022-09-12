package services

import (
	"context"
	"daringcraft-backend/database"
	"daringcraft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMedias(filters bson.M, opt *options.FindOptions, query models.Query) models.Result {
	results := make([]models.Media, 0)

	cursor := database.Find(MediaCollection, filters, opt)
	count := database.Count(MediaCollection, filters)

	if cursor == nil {
		return models.Result{
			Data: results,
		}
	}
	for cursor.Next(context.Background()) {
		var Media models.Media

		if cursor.Decode(&Media) == nil {
			results = append(results, Media)
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

func CreateMedia(media models.Media) error {
	_, err := database.InsertOne(MediaCollection, media)
	if err != nil {
		return err
	}

	return nil
}

func DeleteMedia(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"id": id,
	}

	res, err := database.DeleteOne(MediaCollection, filter)

	if res == nil {
		return nil, err
	}

	return res, nil
}