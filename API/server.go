package main

import (
	collection "API/collection"
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Stores a handle to the database being used by the Lambda function
type ConnectionAWS struct {
	client *mongo.Client
}

func (client ConnectionAWS) handleRequestAWS(i context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var ijson map[string]interface{}

	log.Println(request.Body)
	json.Unmarshal([]byte(request.Body), &ijson)

	cDBRef := collection.Transaction("recruiter", ijson, client.client)

	data := map[string]interface{}{"data": cDBRef}
	jsonStr, _ := json.Marshal(data)

	return events.APIGatewayProxyResponse{Body: string(jsonStr), StatusCode: 200}, nil

}

func mainAWS() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("ATLAS_URI")))
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(ctx)

	connection := ConnectionAWS{
		client: client,
	}

	lambda.Start(connection.handleRequestAWS)
}
