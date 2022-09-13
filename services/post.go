package services

import (
	"context"
	"daringcraft-backend/database"
	"daringcraft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const PostCollection = "posts"

func GetPosts(filters bson.M, opt *options.FindOptions, query models.Query) models.Result {
	results := make([]models.Post, 0)

	cursor := database.Find(PostCollection, filters, opt)
	count := database.Count(PostCollection, filters)

	if cursor == nil {
		return models.Result{
			Data: results,
		}
	}
	for cursor.Next(context.Background()) {
		var Post models.Post

		if cursor.Decode(&Post) == nil {
			results = append(results, Post)
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

func CreatePost(Post models.Post) (bool, error) {
	_, err := database.InsertOne(PostCollection, Post)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetPost(filter bson.M, opts bson.D) *models.Post {
	option := bson.D{}
	if opts != nil {
		option = opts
	}

	cursor := database.FindOne(PostCollection, filter, options.FindOne().SetProjection(option))

	if cursor == nil {
		return nil
	}

	var Post models.Post

	err := cursor.Decode(&Post)
	if err != nil {
		return nil
	}

	return &Post
}

func UpdatePost(id string, Post models.Post) (*mongo.UpdateResult, error) {
	filters := bson.M{
		"id": id,
	}

	res, err := database.UpdateOne(PostCollection, filters, Post)

	if res == nil {
		return nil, err
	}

	return res, nil
}

func DeletePost(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"id": id,
	}

	res, err := database.DeleteOne(PostCollection, filter)

	if res == nil {
		return nil, err
	}

	return res, nil
}

