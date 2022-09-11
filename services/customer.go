package services

import (
	"context"
	"daringcraft-backend/database"
	"daringcraft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CustomerCollection = "customers"

func GetCustomers(filters bson.M, opt *options.FindOptions, query models.Query) models.Result {
	results := make([]models.Customer, 0)

	cursor := database.Find(CustomerCollection, filters, opt)
	count := database.Count(CustomerCollection, filters)

	if cursor == nil {
		return models.Result{
			Data: results,
		}
	}
	for cursor.Next(context.Background()) {
		var Customer models.Customer

		if cursor.Decode(&Customer) == nil {
			results = append(results, Customer)
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

func CreateCustomer(Customer models.Customer) (bool, error) {
	_, err := database.InsertOne(CustomerCollection, Customer)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetCustomer(filter bson.M, opts bson.D) *models.Customer {
	option := bson.D{}
	if opts != nil {
		option = opts
	}

	cursor := database.FindOne(CustomerCollection, filter, options.FindOne().SetProjection(option))

	if cursor == nil {
		return nil
	}

	var Customer models.Customer

	err := cursor.Decode(&Customer)
	if err != nil {
		return nil
	}

	return &Customer
}

func UpdateCustomer(id string, Customer models.Customer) (*mongo.UpdateResult, error) {
	filters := bson.M{
		"id": id,
	}

	res, err := database.UpdateOne(CustomerCollection, filters, Customer)

	if res == nil {
		return nil, err
	}

	return res, nil
}

func DeleteCustomer(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"id": id,
	}

	res, err := database.DeleteOne(CustomerCollection, filter)

	if res == nil {
		return nil, err
	}

	return res, nil
}
