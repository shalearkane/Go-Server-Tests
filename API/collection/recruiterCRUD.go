package collection

import (
	constant "API/constant"
	model "API/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func RecruiterCreate(r map[string]interface{}, client *mongo.Client, sessionContext mongo.SessionContext) (model.DBRef, error) {
	var recruiter model.Recruiter

	company := r["company"].(map[string]interface{})
	recruiter.Company, _ = companyCreate(company, client, sessionContext)

	if FirstName, ok := r["firstName"].(string); ok {
		recruiter.FirstName = FirstName
	}

	recruiterCollection := client.Database(constant.DB).Collection(constant.RecruiterCollection)
	result, err := recruiterCollection.InsertOne(sessionContext, recruiter)

	cDBRef := model.DBRef{}

	if InsertedID, ok := result.InsertedID.(primitive.ObjectID); ok {
		cDBRef.ID = InsertedID
		cDBRef.DB = constant.DB
		cDBRef.Ref = constant.RecruiterCollection

	}
	return cDBRef, err
}
