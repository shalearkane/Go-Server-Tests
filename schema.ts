// Table
export interface Recruiter {
    _id: Types.ObjectId;
    firstName: String;
    middleName: String;
    lastName: String;
    gender: String;
    phone: String;
    email: String;
    company: Company;
    isActive: boolean;
    createdAt: Date;
    updatedAt: Date;
    access: Array<Access>
}
// Table
export interface Company {
    _id: Types.ObjectId;
    name: String;
    address: String;
    companyTurnover?: String;
    type: String;
    sector: String;
    isActive: boolean;
}
// Table
export interface StudentApplication {
    studentId: Student,
    resumeId: Resume,
    appliedFor: JobDescription,
}

export enum RecruitmentType = {"Internship", "Placement"}
export enum InternshipMode = {"Online", "Offline","Hybrid"}
export enum InternshipMode = {"Shortlist from resume", "Written test","Hybrid"}

export interface EligibilityCriteria {
    eligibleBranches: Array<String>;
    cgpa:Number;
    "X%": Number;
    "XII%": Number;
}


export interface InternshipBranchDetails {
    eligibility: EligibilityCriteria;
    internshipDuration: number;
    stipend: string;
    accommodation: string;
    relocationCompensation?: string;
    perks?: string;
    internshipMode:InternshipMode;
    ageLimit?:number;
}

// Table
export interface JobDescription {
    noOfHires: number;
    tentativeLocations: Array<string>;
    description?: String;
    attachment?: String;
    branchDetails: {
        btech?: InternshipBranchDetails,
        idd?: InternshipBranchDetails,
        mtech?: InternshipBranchDetails,
        phd?: InternshipBranchDetails
    },
    studentApplications?: Array<StudentApplication>
    venue?: Array<Venue>;
    schedule?: Array<Schedule>;
}

// Table
export interface InternshipForm {
    type: RecruitmentType;
    company: Company;
    recruiter: Array<Recruiter>;
    internshipProfile: Array<JobDescription>
    selectionProcedure: Array<SelectionProcedure>
    termsAccepted: boolean;
    isActive: boolean;
}