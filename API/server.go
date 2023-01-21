package main

import (
	// getcollection "API/collection"
	model "API/model"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Stores a handle to the collection being used by the Lambda function

// A data structure representation of the collection schema

type DB struct {
	database *mongo.Database
}

func (db DB) handleRequest(i context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var recruiterCollection = db.database.Collection("Recruiter")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var recruiter model.Recruiter

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.

	var ijson map[string]interface{}

	json.Unmarshal([]byte(request.Body), &ijson)

	recruiter = model.Recruiter{}
	if FirstName, ok := ijson["FirstName"].(string); ok {
		recruiter.FirstName = FirstName
	}

	company := ijson["company"].(map[string]interface{})

	if Name, ok := company["Name"].(string); ok {
		fmt.Println(Name)
	}

	result, err := recruiterCollection.InsertOne(ctx, recruiter)

	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{Body: "failed!!", StatusCode: 200}, nil
	}

	data := map[string]interface{}{"data": result}
	jsonStr, _ := json.Marshal(data)

	return events.APIGatewayProxyResponse{Body: string(jsonStr), StatusCode: 200}, nil

}
func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("ATLAS_URI")))
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(ctx)

	connection := DB{
		database: client.Database("GoAPI"),
	}

	lambda.Start(connection.handleRequest)
}
