package models


type OrderedCourse struct {
	Course		string		`json:"course"`
	Order		uint8		`json:"order"`
}

type OrderedCourses struct {
	Courses		[]OrderedCourse		`json:"orderedCourses"`
}


func (oc *OrderedCourses) AddCourse(name string, order uint8) {

	for i := 0; i < len(oc.Courses); i++ {
		if oc.Courses[i].Course == name {
			return
		}
	}

	course := OrderedCourse{
		Course: name,
		Order: order,
	}

	oc.Courses = append(oc.Courses, course)
}