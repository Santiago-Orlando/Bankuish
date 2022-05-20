package models


type CourseInfo struct {
	DesiredCourse		string		`json:"desiredCourse"`
	RequiredCourse		string		`json:"requiredCourse"`
}

type Courses struct {
	UserID		string				`json:"userId"`
	Courses		[]CourseInfo		`json:"courses"`
}

type Course struct {
	Course		string		`json:"course"`
}

type CompletedCourses struct {
	Courses		[]Course	`json:"courses"`
}