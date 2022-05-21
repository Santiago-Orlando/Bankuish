To run the project:
```
docker-compose build
docker-compose up
```
To run the unit tests:
```
docker-compose build
docker-compose up
go test ./test
```
To run the End-To-End test:     // Please read the CAUTION at the bottom!
```
docker-compose build
docker-compose up
cd e2e
npm install
npm test
```


API Endpoints description:

- register:

    Check that the HTTP method is "POST"

    Get a firebase client instance

    Decode the email and password from the request body 

    Expect something like:
            ```
            {
                "email": "your_email@gmail.com",
                "password": "12345678"
            }
            ```
    Use the firebase instance to create a user

    Create the user in the PostgreSQL database

    Send an HTTP header with 201 ( created )


- courseOrder:

    Check that the HTTP method is "PUT"

    Decode the courses data from de body

    Order the courses

    Send the courses ordered with an status code of 200


- courseInscription:

    Check that the HTTP method is "PATCH"

    Get the firebase id from the context ( I get it through the middleware )

    Check that the user is not inscribed in any course

    Decode the course data from de body

    Expect something like:
                ```
                {
                    "course": "Investment"
                }
                ```
    Store the course inside the user register in the database

    Send the 202 status code 


- courseFinished:

    Check that the HTTP method is "PATCH"

    Get the firebase id from the context ( I get it through the middleware )

    Check that the user is inscribed in a course

    Clean the user register in the database To allow the user to re-enroll in another course

    Store the course and the user id in a database table 

    Send the 200 status code 


- courseCompleted:

    Check that the HTTP method is "GET"

    Get the firebase id from the context ( I get it through the middleware )

    Get a list of all courses made by the user

    Send the courses and a status code of 200



Things I would have done if I had more time:

    -Add validations in the register, for example, check that the email is valid and the password is at least 8 characters long.

    -With a table of the courses in the database, I could have made a validation when the user tried to inscribe in a course like if the course does not exist in the database send a 401 status code.

    -Refactor the course ordered function to admit multiple courses with the same order, to let the user decide which course wants to take.

    -Document the API using Swagger ( I have made this here: https://github.com/Santiago-Orlando/Guardian/blob/main/proxy/swagger.yaml)

    -Use ENV Variables in the docker-compose file.

    -Do the API with Node.js, I read the blogs that Adrian send to me and the "supertest" library looks pretty cool!

    -Make more tests.



CAUTION:

To perform the End-To-End test needs a valid token, I tried really hard to find a way to get a firebase token but I don't find any way to log in from the backend.


Trying to solve that I made a super simple react app that shows a valid token. So before running the e2e test load the app in the "Token Getter".


To load the app in the "Token Getter" folder run:

```bash
    cd "Token Getter"
    npm i
    npm start
```

Also if you are in a system with bash run in the root of the project:

```
    sed -i 's/HOST: DATABASE_URL/HOST: TEST_DATABASE_URL/' docker-compose.yml
    sed -i 's/DB_NAME: Bankuish/DB_NAME: test/' docker-compose.yml 
```
If you have no access to bash, please replace manually the __HOST__ and __DB_NAME__ variables values inside the docker-compose.yml file with __TEST_DATABASE_URL__ and __test__.

Copy the token into the first variable in the ./e2e/fakeData.js and then run in the "npm i" and "npm test" command.


I would have liked to give a valid token but only last 1 hour.
