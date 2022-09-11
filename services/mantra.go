package services

import (
	"context"
	"daringcraft-backend/database"
	"daringcraft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MantraCollection = "mantras"

func GetMantras(filters bson.M, opt *options.FindOptions, query models.Query) models.Result {
	results := make([]models.Mantra, 0)

	cursor := database.Find(MantraCollection, filters, opt)
	count := database.Count(MantraCollection, filters)

	if cursor == nil {
		return models.Result{
			Data: results,
		}
	}
	for cursor.Next(context.Background()) {
		var Mantra models.Mantra

		if cursor.Decode(&Mantra) == nil {
			results = append(results, Mantra)
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

func CreateMantra(Mantra models.Mantra) (bool, error) {
	_, err := database.InsertOne(MantraCollection, Mantra)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetMantra(filter bson.M, opts bson.D) *models.Mantra {
	option := bson.D{}
	if opts != nil {
		option = opts
	}

	cursor := database.FindOne(MantraCollection, filter, options.FindOne().SetProjection(option))

	if cursor == nil {
		return nil
	}

	var Mantra models.Mantra

	err := cursor.Decode(&Mantra)
	if err != nil {
		return nil
	}

	return &Mantra
}

func UpdateMantra(id string, Mantra models.Mantra) (*mongo.UpdateResult, error) {
	filters := bson.M{
		"id": id,
	}

	res, err := database.UpdateOne(MantraCollection, filters, Mantra)

	if res == nil {
		return nil, err
	}

	return res, nil
}

func DeleteMantra(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"id": id,
	}

	res, err := database.DeleteOne(MantraCollection, filter)

	if res == nil {
		return nil, err
	}

	return res, nil
}
