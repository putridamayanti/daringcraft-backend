package services

import (
	"context"
	"daringcraft-backend/database"
	"daringcraft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ProductCollection = "products"

func GetProducts(filters bson.M, opt *options.FindOptions, query models.Query) models.Result {
	results := make([]models.Product, 0)

	cursor := database.Find(ProductCollection, filters, opt)
	count := database.Count(ProductCollection, filters)

	if cursor == nil {
		return models.Result{
			Data: results,
		}
	}
	for cursor.Next(context.Background()) {
		var Product models.Product

		if cursor.Decode(&Product) == nil {
			results = append(results, Product)
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

func CreateProduct(Product models.Product) (bool, error) {
	_, err := database.InsertOne(ProductCollection, Product)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetProduct(filter bson.M, opts bson.D) *models.Product {
	option := bson.D{}
	if opts != nil {
		option = opts
	}

	cursor := database.FindOne(ProductCollection, filter, options.FindOne().SetProjection(option))

	if cursor == nil {
		return nil
	}

	var Product models.Product

	err := cursor.Decode(&Product)
	if err != nil {
		return nil
	}

	return &Product
}

func UpdateProduct(id string, Product models.Product) (*mongo.UpdateResult, error) {
	filters := bson.M{
		"id": id,
	}

	res, err := database.UpdateOne(ProductCollection, filters, Product)

	if res == nil {
		return nil, err
	}

	return res, nil
}

func DeleteProduct(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"id": id,
	}

	res, err := database.DeleteOne(ProductCollection, filter)

	if res == nil {
		return nil, err
	}

	return res, nil
}