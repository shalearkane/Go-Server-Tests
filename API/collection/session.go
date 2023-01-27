package collection

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func Transaction(collectionType string, m map[string]interface{}, client *mongo.Client) interface{} {
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(context.Background())

	switch collectionType {
	case "recruiter":
		{
			callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
				return RecruiterCreate(m, client, sessionContext)
			}
			result, err := session.WithTransaction(context.Background(), callback, txnOpts)
			if err != nil {
				panic(err)
			}
			return result
		}

	default:
		{
			panic("Method not implemented")

		}
	}
}
