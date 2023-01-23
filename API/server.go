package main

import (
	collection "API/collection"
	constant "API/constant"
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Stores a handle to the database being used by the Lambda function
type DB struct {
	database *mongo.Database
}

func (db DB) handleRequest(i context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var ijson map[string]interface{}

	json.Unmarshal([]byte(request.Body), &ijson)

	cDBRef := collection.RecruiterCreate(ijson, db.database)

	data := map[string]interface{}{"data": cDBRef}
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
		database: client.Database(constant.DB),
	}

	lambda.Start(connection.handleRequest)
}
