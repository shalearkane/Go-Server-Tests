package collection

import (
	constant "API/constant"
	model "API/model"
	s3 "API/s3"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func RecruiterCreate(r map[string]interface{}, client *mongo.Client, sessionContext mongo.SessionContext) (model.DBRef, error) {
	var recruiter model.Recruiter

	if company, ok := r["company"].(map[string]interface{}); ok {
		recruiter.Company, _ = companyCreate(company, client, sessionContext)
	} else {
		log.Print("This is not map[string]interface : ")
		log.Println(r["company"])
	}

	if FirstName, ok := r["firstName"].(string); ok {
		recruiter.FirstName = FirstName
	} else {
		log.Print("This is not string : ")
		log.Println(r["firstname"])
	}

	if selfie, ok := r["selfie"].(string); ok {
		recruiter.Selfie, _ = s3.Upload(selfie)
	} else {
		log.Print("This is not string : ")
		log.Println(r["selfie"])
	}

	if CompanyArray, ok := r["companyArray"].([]interface{}); ok {
		for _, element := range CompanyArray {
			if c, ok := element.(map[string]interface{}); ok {
				dbref, _ := companyCreate(c, client, sessionContext)
				recruiter.CompanyArray = append(recruiter.CompanyArray, dbref)
			}
		}
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
