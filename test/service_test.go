package test

import (
	"Bankuish/api/services"
	"database/sql"
	"fmt"
	"testing"
)


var DB *sql.DB = GetConnection()

const querySetUser string = "INSERT INTO users ( firebase_id ) VALUES ( $1 )"
const queryGetStudying string = `SELECT studying FROM users WHERE firebase_id = $1`

func TestCourseOrderer(t *testing.T) {

	orderedCourses := services.CourseOrdererService(FakeCourses)
	
	for i := 0; i < len(FakeOrderedCourses.Courses); i++ {

		orderedCourse := orderedCourses.Courses[i].Course
		fakeOrderedCourse := FakeOrderedCourses.Courses[i].Course

		orderedOrden := orderedCourses.Courses[i].Order
		fakeOrderedOrden := FakeOrderedCourses.Courses[i].Order

		if orderedCourse != fakeOrderedCourse {
			t.Fail()
		}

		if orderedOrden != fakeOrderedOrden {
			t.Fail()
		}
	}
	
}

func TestStudyingSetter(t *testing.T) {

	_, err := DB.Query(querySetUser, FirebaseID)
	if err != nil {
		t.FailNow()
	}

	course := "Investment"

	err = services.StudyingSetter(DB, course, FirebaseID)
	if err != nil {
		t.FailNow()
	}

	var studying sql.NullString

	err = DB.QueryRow(queryGetStudying, FirebaseID).Scan(&studying)
	if err != nil {
		t.Fail()
	}

	if studying.String != "Investment" {
		t.Fail()
	}

}

func TestStudyingGetter(t *testing.T) {

	studying, err := services.StudyingGetter(DB, FirebaseID)
	if err != nil {
		t.Fail()
	}
	if studying != "Investment" {
		t.Fail()
	}


	studyingEmpty, err := services.StudyingGetter(DB, FakeID)
	if err == nil {
		t.Fail()
	}
	if studyingEmpty != "" {
		t.Fail()
	}

}

func TestCourseFinished(t *testing.T) {

	course := "Investment"
	noCourse := ""

	err := services.CourseFinished(DB, course, FirebaseID)
	if err != nil {
		t.Fail()
	}

	var studying sql.NullString

	err = DB.QueryRow(queryGetStudying, FirebaseID).Scan(&studying)
	if err != nil {
		t.Fail()
	}

	if studying.String != "" {
		t.Fail()
	}

	err = services.CourseFinished(DB, noCourse, FirebaseID)
	if err == nil {
		t.Fail()
	}

}

func TestCoursesGetter(t *testing.T) {

	completedCourses, err := services.CoursesGetter(DB, FirebaseID)
	if err != nil {
		fmt.Println(120, err)
		t.Fail()
	}
	if completedCourses.Courses[0].Course != "Investment"{
		t.Fail()
	}

	completedCoursesEmpty, err := services.CoursesGetter(DB, FakeID)
	if err != nil {
		t.Fail()
	}
	if len(completedCoursesEmpty.Courses) != 0 {
		t.Fail()
	}

}

func TestDropTableContent(t *testing.T) {

	queryDropUsers := `DELETE FROM users`
	queryDropCourses := `DELETE FROM course_made_by_user`

	_, err := DB.Query(queryDropUsers)
	if err != nil {
		t.Fail()
	}
	_, err = DB.Query(queryDropCourses)
	if err != nil {
		t.Fail()
	}
	
}