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

func companyCreate(c map[string]interface{}, db *mongo.Database) model.DBRef {
	var companyCollection = db.Collection(constant.CompanyCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var company model.Company
	if Name, ok := c["name"].(string); ok {
		company.Name = Name
	}

	result, err := companyCollection.InsertOne(ctx, company)
	if err != nil {
		fmt.Println(err)
	}

	cDBRef := model.DBRef{}

	if InsertedID, ok := result.InsertedID.(primitive.ObjectID); ok {
		cDBRef.ID = InsertedID
		cDBRef.DB = constant.DB
		cDBRef.Ref = constant.CompanyCollection

	}

	return cDBRef

}
