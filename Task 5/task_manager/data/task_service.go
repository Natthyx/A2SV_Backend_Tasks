package data

import (
	"context"
	"log"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var taskCollection *mongo.Collection

func InitMongoDB(uri string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	taskCollection = client.Database("taskdb").Collection("tasks")
}

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := taskCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskByID(id string) (models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, err
	}
	var task models.Task
	err = taskCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func CreateTask(task models.Task) (models.Task, error) {
	res, err := taskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		return models.Task{}, err
	}
	task.ID = res.InsertedID.(primitive.ObjectID)
	return task, nil
}

func UpdateTask(id string, updated models.Task) (models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, err
	}
	updated.ID = objID
	_, err = taskCollection.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": updated})
	if err != nil {
		return models.Task{}, err
	}
	return updated, nil
}

func DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = taskCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}
