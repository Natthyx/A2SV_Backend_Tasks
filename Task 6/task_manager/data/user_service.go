package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"

	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

var userCollection *mongo.Collection

func InitMongoDB(uri string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	taskCollection = client.Database("taskdb").Collection("tasks")
	userCollection = client.Database("taskdb").Collection("users")
}

func RegisterUser(user models.User) error {
	var existing models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existing)
	if err == nil {
		return errors.New("username already taken")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)
	user.Role = "user"
	_, err = userCollection.InsertOne(context.TODO(), user)
	return err
}

func AuthenticateUser(username, password string) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return user, errors.New("user not found")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return user, errors.New("invalid credentials")
	}
	return user, nil
}