package collection

import (
	constant "API/constant"
	model "API/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func companyCreate(c map[string]interface{}, client *mongo.Client, sessionContext mongo.SessionContext) (model.DBRef, error) {
	var company model.Company
	if Name, ok := c["name"].(string); ok {
		company.Name = Name
	}

	companyCollection := client.Database(constant.DB).Collection(constant.CompanyCollection)

	result, err := companyCollection.InsertOne(sessionContext, company)
	cDBRef := model.DBRef{}

	if InsertedID, ok := result.InsertedID.(primitive.ObjectID); ok {
		cDBRef.ID = InsertedID
		cDBRef.DB = constant.DB
		cDBRef.Ref = constant.CompanyCollection

	}

	return cDBRef, err
}
