package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recruiter struct {
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	IsActive   bool   `json:"isActive"`
	CreatedAt  primitive.DateTime
	UpdatedAt  primitive.DateTime
	Company    DBRef
}

type RecruiterFetch struct {
	ID primitive.ObjectID `json:"_id"`
	Recruiter
}

type DBRef struct {
	Ref interface{} `bson:"$ref"`
	ID  interface{} `bson:"$id"`
	DB  interface{} `bson:"$db"`
}

type Company struct {
	Name            string `json:"name"`
	Address         string `json:"address"`
	CompanyTurnover string `json:"companyTurnover"`
	CompanyType     string `json:"companyType"`
	Sector          string `json:"sector"`
	IsActive        bool   `json:"isActive"`
}

type CompanyFetch struct {
	ID primitive.ObjectID `json:"_id"`
	Company
}