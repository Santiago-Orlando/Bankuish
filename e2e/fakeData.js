const token = "";
const fakeToken = "zyJhbGciOiJSUzI1NiIsImtpZCI6IjY5N2Q3ZmI1ZGNkZThjZDA0OGQzYzkxNThiNjIwYjY5MTA1MjJiNGQiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vYmFua3Vpc2gtYTQ0YWQiLCJhdWQiOiJiYW5rdWlzaC1hNDRhZCIsImF1dGhfdGltZSI6MTY1Mjk4OTUzMSwidXNlcl9pZCI6InkxRXpyeFFMZ1dSTkZYNXh0TFRzSGRvd01tcjIiLCJzdWIiOiJ5MUV6cnhRTGdXUk5GWDV4dExUc0hkb3dNbXIyIiwiaWF0IjoxNjUyOTg5NTMxLCJleHAiOjE2NTI5OTMxMzEsImVtYWlsIjoic2VvcmxhbmRvMzNAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7ImVtYWlsIjpbInNlb3JsYW5kbzMzQGdtYWlsLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6InBhc3N3b3JkIn19.ZdW4ePPkYHeadpB5cur1PQU8OhzWpHgQJ31CCGny4OZnlqp5pdlZVppe5v25qpNr3XlfroEH8L0WCX-9hYHvKdTWjjcI8PHKPld0xBlBmprMD2KUEFHQJUI8VYSNr1NwI2FboYwMOhBAvWb0IMRzPLtF08RLx63S8uKhSTRf847nulOb20VwAxAcnWyJtwjOHMzIBuFUURpzy4TzAq4dLg5SGIRBjEuZOi-1EEJu5UtQq3LnuZtwh1C33uY7v_8uO9kkhS-5qzK1YZQYpJ_0eCIu1cJbGwMNzjjzHOUgJ6rvGayxpV4HlcIgUEGPRCtbE1kaXDFrV78K75Qs0xOJtQ";

const courses = {
  userId: "30ecc27b-9df7-4dd3-b52f-d001e79bd035",
  courses: [
    {
      desiredCourse: "PortfolioConstruction",
      requiredCourse: "PortfolioTheories",
    },
    {
      desiredCourse: "InvestmentManagement",
      requiredCourse: "Investment",
    },
    {
      desiredCourse: "Investment",
      requiredCourse: "Finance",
    },
    {
      desiredCourse: "PortfolioTheories",
      requiredCourse: "Investment",
    },
    {
      desiredCourse: "InvestmentStyle",
      requiredCourse: "InvestmentManagement",
    },
  ],
};

const orderedCourses = {
  orderedCourses: [
    { course: "Finance", order: 0 },
    { course: "Investment", order: 1 },
    { course: "InvestmentManagement", order: 2 },
    { course: "PortfolioTheories", order: 3 },
    { course: "PortfolioConstruction", order: 4 },
    { course: "InvestmentStyle", order: 5 },
  ],
};

const firebaseID = "y1EzrxQLgWRNFX5xtLTsHdowMmr2"

module.exports = {
    token,
    fakeToken,
    courses,
    orderedCourses,
    firebaseID
}