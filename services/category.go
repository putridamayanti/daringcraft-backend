package services

import (
	"context"
	"daringcraft-backend/database"
	"daringcraft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CategoryCollection = "categories"

func GetCategories(filters bson.M, opt *options.FindOptions, query models.Query) models.Result {
	results := make([]models.Category, 0)

	cursor := database.Find(CategoryCollection, filters, opt)
	count := database.Count(CategoryCollection, filters)

	if cursor == nil {
		return models.Result{
			Data: results,
		}
	}
	for cursor.Next(context.Background()) {
		var category models.Category

		if cursor.Decode(&category) == nil {
			results = append(results, category)
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

func CreateCategory(category models.Category) (bool, error) {
	_, err := database.InsertOne(CategoryCollection, category)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetCategory(filter bson.M, opts bson.D) *models.Category {
	option := bson.D{}
	if opts != nil {
		option = opts
	}

	cursor := database.FindOne(CategoryCollection, filter, options.FindOne().SetProjection(option))

	if cursor == nil {
		return nil
	}

	var Category models.Category

	err := cursor.Decode(&Category)
	if err != nil {
		return nil
	}

	return &Category
}

func UpdateCategory(id string, category models.Category) (*mongo.UpdateResult, error) {
	filters := bson.M{
		"id": id,
	}

	res, err := database.UpdateOne(CategoryCollection, filters, category)

	if res == nil {
		return nil, err
	}

	return res, nil
}

func DeleteCategory(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"id": id,
	}

	res, err := database.DeleteOne(CategoryCollection, filter)

	if res == nil {
		return nil, err
	}

	return res, nil
}