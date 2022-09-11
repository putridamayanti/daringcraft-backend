package services

import (
	"context"
	"daringcraft-backend/database"
	"daringcraft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const BannerCollection = "banners"

func GetBanners(filters bson.M, opt *options.FindOptions, query models.Query) models.Result {
	results := make([]models.Banner, 0)

	cursor := database.Find(BannerCollection, filters, opt)
	count := database.Count(BannerCollection, filters)

	if cursor == nil {
		return models.Result{
			Data: results,
		}
	}
	for cursor.Next(context.Background()) {
		var Banner models.Banner

		if cursor.Decode(&Banner) == nil {
			results = append(results, Banner)
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

func CreateBanner(banner models.Banner) (bool, error) {
	_, err := database.InsertOne(BannerCollection, banner)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetBanner(filter bson.M, opts bson.D) *models.Banner {
	option := bson.D{}
	if opts != nil {
		option = opts
	}

	cursor := database.FindOne(BannerCollection, filter, options.FindOne().SetProjection(option))

	if cursor == nil {
		return nil
	}

	var Banner models.Banner

	err := cursor.Decode(&Banner)
	if err != nil {
		return nil
	}

	return &Banner
}

func UpdateBanner(id string, banner models.Banner) (*mongo.UpdateResult, error) {
	filters := bson.M{
		"id": id,
	}

	res, err := database.UpdateOne(BannerCollection, filters, banner)

	if res == nil {
		return nil, err
	}

	return res, nil
}

func DeleteBanner(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"id": id,
	}

	res, err := database.DeleteOne(BannerCollection, filter)

	if res == nil {
		return nil, err
	}

	return res, nil
}
