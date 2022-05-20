package test

import (
	"Bankuish/api/models"
)

var FakeCourses models.Courses = models.Courses{
	UserID: "y1EzrxQLgWRNFX5xtLTsHdowMmr2",
	Courses: []models.CourseInfo{
		{
			DesiredCourse: "PortfolioConstruction",
			RequiredCourse: "PortfolioTheories",
			},
			{
			DesiredCourse: "InvestmentManagement",
			RequiredCourse: "Investment",
			},
			{
			DesiredCourse: "Investment",
			RequiredCourse: "Finance",
			},
			{
			DesiredCourse: "PortfolioTheories",
			RequiredCourse: "Investment",
			},
			{
			DesiredCourse: "InvestmentStyle",
			RequiredCourse: "InvestmentManagement",
			},
	},
}

var FakeOrderedCourses models.OrderedCourses = models.OrderedCourses{

	Courses: []models.OrderedCourse{
		{
			Course: "Finance",
			Order: 0,
		},
		{
			Course:"Investment",
			Order:1,
		},
		{
			Course:"InvestmentManagement",
			Order:2,
		},
		{
			Course:"PortfolioTheories",
			Order:3,
		},
		{
			Course:"PortfolioConstruction",
			Order:4,
		},
		{
			Course:"InvestmentStyle",
			Order:5,
		},
	},
}

var CompletedCourses models.CompletedCourses = models.CompletedCourses{
	Courses: []models.Course{
		{
			Course: "Investment",
		},
	},
}

const FirebaseID string = 	"y1EzrxQLgWRNFX5xtLTsHdowMmr2"
const FakeID string = 		"e1EzrxQLgWRNFX5xtLTsHdowMmr2"