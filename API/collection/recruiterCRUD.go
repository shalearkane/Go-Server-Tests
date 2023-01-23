package collection

import (
	constant "API/constant"
	model "API/model"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func RecruiterCreate(r map[string]interface{}, db *mongo.Database) model.DBRef {
	var recruiterCollection = db.Collection(constant.RecruiterCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var recruiter model.Recruiter

	company := r["company"].(map[string]interface{})
	recruiter.Company = companyCreate(company, db)

	if FirstName, ok := r["FirstName"].(string); ok {
		recruiter.FirstName = FirstName
	}

	result, err := recruiterCollection.InsertOne(ctx, recruiter)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.InsertedID)
	}

	cDBRef := model.DBRef{}

	if InsertedID, ok := result.InsertedID.(primitive.ObjectID); ok {
		cDBRef.ID = InsertedID
		cDBRef.DB = constant.DB
		cDBRef.Ref = constant.RecruiterCollection

	}

	return cDBRef
}
