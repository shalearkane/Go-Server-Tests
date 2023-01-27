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

func RecruiterCreate(r map[string]interface{}, client *mongo.Client) model.DBRef {
	var recruiter model.Recruiter

	company := r["company"].(map[string]interface{})
	recruiter.Company = companyCreate(company, client)

	if FirstName, ok := r["firstName"].(string); ok {
		recruiter.FirstName = FirstName
	}

	result := recruiterCreateTransact(recruiter, client)
	cDBRef := model.DBRef{}

	if InsertedID, ok := result.InsertedID.(primitive.ObjectID); ok {
		cDBRef.ID = InsertedID
		cDBRef.DB = constant.DB
		cDBRef.Ref = constant.RecruiterCollection

	}

	return cDBRef
}

func recruiterCreateTransact(r model.Recruiter, client *mongo.Client) *mongo.InsertOneResult {
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

		recruiterCollection := client.Database(constant.DB).Collection(constant.RecruiterCollection)

		result, err := recruiterCollection.InsertOne(ctx, r)

		return result, err
	}

	result, err := session.WithTransaction(context.Background(), callback, txnOpts)
	if err != nil {
		panic(err)
	} else {
		return result.(*mongo.InsertOneResult)
	}
}
