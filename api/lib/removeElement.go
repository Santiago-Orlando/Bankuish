package lib

import "Bankuish/api/models"


func RemoveElement(courses models.Courses, index int) models.Courses {

	courses.Courses[index] = courses.Courses[len(courses.Courses)-1]
	courses.Courses[len(courses.Courses)-1] = models.CourseInfo{}
	courses.Courses = courses.Courses[:len(courses.Courses)-1] 

	return courses
}