const fetch = require("node-fetch");
const pool = require("./testDatabase");

const {
  courses,
  firebaseID,
  fakeToken,
  orderedCourses,
  token,
} = require("./fakeData");

const querySetUser = "INSERT INTO users ( firebase_id ) VALUES ( $1 )";
const queryGetStudying = `SELECT studying FROM users WHERE firebase_id = $1`;

const URL = "http://localhost:3001/";

beforeAll(async () => {
  await pool.query(querySetUser, [firebaseID]);
});

describe("test API Endpoint", () => {
  test("PUT courseOrder unauthorized", async () => {
    const res = await fetch(URL + "courseOrder", {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: fakeToken,
      },
      body: courses,
    });

    expect(res.status).toEqual(401);
  });

  test("PUT courseOrder", async () => {
    const res = await fetch(URL + "courseOrder", {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: token,
      },
      body: JSON.stringify(courses),
    });
    let body = await res.json();

    expect(res.status).toEqual(200);

    expect(body.orderedCourses[0].course).toEqual(
      orderedCourses.orderedCourses[0].course
    );
    expect(body.orderedCourses[0].order).toEqual(
      orderedCourses.orderedCourses[0].order
    );

    expect(body.orderedCourses[1].course).toEqual(
      orderedCourses.orderedCourses[1].course
    );
    expect(body.orderedCourses[1].order).toEqual(
      orderedCourses.orderedCourses[1].order
    );

    expect(body.orderedCourses[5].course).toEqual(
      orderedCourses.orderedCourses[5].course
    );
    expect(body.orderedCourses[5].order).toEqual(
      orderedCourses.orderedCourses[5].order
    );
  });

  test("PATCH courseInscription", async () => {
    const res = await fetch(URL + "courseInscription", {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
        Authorization: token,
      },
      body: JSON.stringify({ course: "Investment" }),
    });

    let row = await pool.query(queryGetStudying, [firebaseID]);
    row = row.rows;
    console.log(row);

    expect(res.status).toEqual(202);
    expect(row[0]).toEqual({ studying: "Investment" });
  });

  test("PATCH courseFinished", async () => {

    const res = await fetch(URL + "courseFinished", {
      method: "PATCH",
      headers: {
        Authorization: token,
      },
    });

    expect(res.status).toEqual(200)
  });

  test("GET courseCompleted", async () => {

    const res = await fetch(URL + "courseCompleted", {
        method: "GET",
        headers: {
          Authorization: token,
        },
      });

    let { courses } = await res.json()
    
    expect(courses[0].course).toEqual("Investment")
  })
      
});

afterAll(async () => {
  const queryDropUsers = `DELETE FROM users`;
  const queryDropCourses = `DELETE FROM course_made_by_user`;

  await pool.query(queryDropUsers);
  await pool.query(queryDropCourses);
});
