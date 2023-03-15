package main

import (
	collection "API/collection"
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Stores a handle to the database being used by the Lambda function
type ConnectionGCP struct {
	client *mongo.Client
}

func (client ConnectionGCP) handleRequestGCP(c *gin.Context) {
	var ijson map[string]interface{}

	jsonData, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal([]byte(jsonData), &ijson)

	cDBRef := collection.Transaction("recruiter", ijson, client.client)

	data := map[string]interface{}{"data": cDBRef}
	jsonStr, _ := json.Marshal(data)

	c.JSON(200, jsonStr)

}

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("ATLAS_URI")))
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(ctx)

	connection := ConnectionGCP{
		client: client,
	}

	router := gin.Default()
	router.POST("/", connection.handleRequestGCP)

	router.Run()
}
