package collection

import (
	constant "API/constant"
	model "API/model"
	"context"

	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func companyCreate(c map[string]interface{}, client *mongo.Client) model.DBRef {
	var company model.Company
	if Name, ok := c["name"].(string); ok {
		company.Name = Name
	}

	result := companyCreateTransact(company, client)

	cDBRef := model.DBRef{}

	if InsertedID, ok := result.InsertedID.(primitive.ObjectID); ok {
		cDBRef.ID = InsertedID
		cDBRef.DB = constant.DB
		cDBRef.Ref = constant.CompanyCollection

	}

	return cDBRef

}

func companyCreateTransact(c model.Company, client *mongo.Client) *mongo.InsertOneResult {
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(context.Background())

	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		companyCollection := client.Database(constant.DB).Collection(constant.CompanyCollection)

		result, err := companyCollection.InsertOne(ctx, c)

		return result, err
	}

	result, err := session.WithTransaction(context.Background(), callback, txnOpts)
	if err != nil {
		panic(err)
	} else {
		return result.(*mongo.InsertOneResult)
	}
}
