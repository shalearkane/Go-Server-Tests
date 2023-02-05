package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recruiter struct {
	FirstName    string `json:"firstName"`
	MiddleName   string `json:"middleName"`
	LastName     string `json:"lastName"`
	Gender       string `json:"gender"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	IsActive     bool   `json:"isActive"`
	CreatedAt    primitive.DateTime
	UpdatedAt    primitive.DateTime
	Company      DBRef
	Selfie       string  `json:"selfie"`
	CompanyArray []DBRef `json:"companyArray"`
}

type RecruiterFetch struct {
	ID primitive.ObjectID `json:"_id"`
	Recruiter
}

type DBRef struct {
	Ref string             `bson:"$ref" json:"$ref"`
	ID  primitive.ObjectID `bson:"$id" json:"$id"`
	DB  string             `bson:"$db" json:"$db"`
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

type StudentApplication struct {
	StudentID  DBRef
	ResumeID   DBRef
	AppliedFor DBRef
}

type EligibilityCriteria struct {
	EligibleBranches []string
	Cgpa             float32
	XPercent         int
	XIIPercent       int
}

type InternshipBranchDetails struct {
	Eligibility            EligibilityCriteria
	InternshipDuration     int
	Stipend                string
	Accomodation           string
	RelocationCompensation string
	Perks                  string
	InternshipMode         int
	AgeLimit               int
}

type BranchDetails struct {
	BTech InternshipBranchDetails
	IDD   InternshipBranchDetails
	MTech InternshipBranchDetails
	PHd   InternshipBranchDetails
}

type JobDescription struct {
	NumberOfHires      int
	TentativeLocations []string
	Description        string
	Attachement        string
	BranchDetails      BranchDetails
}

type InternshipForm struct {
	RecruitmentType    string
	Company            DBRef
	Recruiter          []DBRef
	IntershipProfile   []JobDescription
	SelectionProcedure []string
	TermsAccepted      bool
	IsActive           bool
}
