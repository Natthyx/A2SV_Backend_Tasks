package data

import (
	"context"
	"errors"
	"task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var taskCollection *mongo.Collection


func CreateTask(task *models.Task) error {
	task.ID = primitive.NewObjectID()
	_, err := taskCollection.InsertOne(context.TODO(), task)
	return err
}

func GetTasks(userID primitive.ObjectID) ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := taskCollection.Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTask(id primitive.ObjectID) (models.Task, error) {
	var task models.Task
	err := taskCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return task, errors.New("task not found")
	}
	return task, nil
}

func UpdateTask(id primitive.ObjectID, updated models.Task) error {
	_, err := taskCollection.UpdateOne(context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": bson.M{
			"title":       updated.Title,
			"description": updated.Description,
			"due_date":    updated.DueDate,
			"status":      updated.Status,
		}},
	)
	return err
}

func DeleteTask(id primitive.ObjectID) error {
	_, err := taskCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
