package services

import (
	"context"
	"daringcraft-backend/database"
	"daringcraft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const PageCollection = "pages"

func GetPages(filters bson.M, opt *options.FindOptions, query models.Query) models.Result {
	results := make([]models.Page, 0)

	cursor := database.Find(PageCollection, filters, opt)
	count := database.Count(PageCollection, filters)

	if cursor == nil {
		return models.Result{
			Data: results,
		}
	}
	for cursor.Next(context.Background()) {
		var Page models.Page

		if cursor.Decode(&Page) == nil {
			results = append(results, Page)
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

func CreatePage(Page models.Page) (bool, error) {
	_, err := database.InsertOne(PageCollection, Page)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetPage(filter bson.M, opts bson.D) *models.Page {
	option := bson.D{}
	if opts != nil {
		option = opts
	}

	cursor := database.FindOne(PageCollection, filter, options.FindOne().SetProjection(option))

	if cursor == nil {
		return nil
	}

	var Page models.Page

	err := cursor.Decode(&Page)
	if err != nil {
		return nil
	}

	return &Page
}

func UpdatePage(id string, Page models.Page) (*mongo.UpdateResult, error) {
	filters := bson.M{
		"id": id,
	}

	res, err := database.UpdateOne(PageCollection, filters, Page)

	if res == nil {
		return nil, err
	}

	return res, nil
}

func DeletePage(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"id": id,
	}

	res, err := database.DeleteOne(PageCollection, filter)

	if res == nil {
		return nil, err
	}

	return res, nil
}
