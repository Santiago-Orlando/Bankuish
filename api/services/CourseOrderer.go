package services

import (
	"Bankuish/api/lib"
	"Bankuish/api/models"
)

func CourseOrdererService(courses models.Courses) models.OrderedCourses {

	orderedCourses := models.OrderedCourses{}

	if len(courses.Courses) == 0 {
		return orderedCourses
	}

	var index int

	for {

		var lastOrderPlusOne uint8

		requiredCourse := courses.Courses[index].RequiredCourse
		desiredCourse := courses.Courses[index].DesiredCourse

		if len(orderedCourses.Courses) != 0 {
			lastOrderPlusOne = orderedCourses.Courses[len(orderedCourses.Courses)-1].Order
		} else {
			lastOrderPlusOne = 0
		}

		if len(courses.Courses) == 1 {

			orderedCourses.AddCourse(requiredCourse, lastOrderPlusOne)
			orderedCourses.AddCourse(desiredCourse, lastOrderPlusOne+1)

			break
		}

		for i := 0; i < len(courses.Courses); i++ {
			if requiredCourse == courses.Courses[i].DesiredCourse {
				requiredCourse = ""
				break
			}
		}

		if requiredCourse != "" {

			orderedCourses.AddCourse(requiredCourse, lastOrderPlusOne)
			orderedCourses.AddCourse(desiredCourse, lastOrderPlusOne+1)

			courses = lib.RemoveElement(courses, index)

			index = -1
		}

		index++
	}

	return orderedCourses
}
