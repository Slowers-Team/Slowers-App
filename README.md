# Slowers-App

Project done for the Software Lab course (TKT20007) at the University of Helsinki. Helps connecting local flower growers to local flower shops in the spirit of the [Slow Flowers](https://en.wikipedia.org/wiki/Slow_Flowers) movement.

- [Staging server](https://slowers-app.ext.ocp-test-0.k8s.it.helsinki.fi)
- [Changelog](documentation/changelog.md)
- [Product backlog](https://github.com/orgs/Slowers-Team/projects/17)
- [Sprint backlogs](https://github.com/orgs/Slowers-Team/projects?query=is%3Aopen+Sprint)
- [Coding conduct](documentation/coding_conduct.md)
- [Definition of done](documentation/definition_of_done.md)

## Starting the application locally

There are two ways to run the application locally: the first one uses Docker containers and the second one is more traditional way of installing and building application locally.

### Run using Docker development containers

TLDR:
```
docker compose -f docker-compose.dev.yml up
[ application runs in http://localhost:8080 ]
[ Ctrl+C to exit ]

docker compose -f docker-compose.dev.yml up -d
docker compose -f docker-compose.dev.yml down

docker compose -f docker-compose.dev.yml up --build

docker exec slowers-backend-dev go test ./...
docker exec slowers-backend-dev go test ./... -tags=api
docker exec slowers-frontend-dev npm run test

docker exec -it slowers-mongo-dev mongosh -u root -p example
docker exec -it slowers-psql-dev psql -U Slowers
[ "exit" to exit ]
```

You need installation of Docker in your machine. [Fullstack MOOC part 12](https://fullstackopen.com/en/part12) has [links and some basic terminology](https://fullstackopen.com/en/part12/introduction_to_containers#installing-everything-required-for-this-part) you will need. Docker Desktop is also useful.

You can run development containers of either databases, backend, frontend or all three simultaneously (recommended). Changing files restarts backend server container and updates pages served from frontend container automatically. Note that you might have to remove folder "psql_data_dev" in [dbs](dbs/) folder if files used in Postgres setup ([schema.sql](backend/databases/sql/schema.sql) and [functions.sql](backend/databases/sql/functions.sql)) are modified.

To run whole application inside containers, build and run containers from repository root directory with command `docker compose -f docker-compose.dev.yml up`. Now you can use address http://localhost:8080 to access frontend and address http://localhost:8080/api to access backend. You can exit and close containers with Ctrl+C.

To run unit tests for frontend, you can use command `docker exec slowers-frontend-dev npm run test` while slowers-frontend-dev -container is running. Unit tests for backend can similarly be run using command `docker exec slowers-backend-dev go test ./...`. Another option is to go inside containers e.g. in VSCode and use normal test commands there.

API tests are currently broken, those can be run using command `docker exec slowers-backend-dev go test ./... -tags=api`.

More instructions can be found [here](documentation/docker_instructions.md).

### Run by installing and building locally

1. Run MongoDB either online in MongoDB Atlas ([Instructions in Finnish](https://fullstackopen.com/osa3/tietojen_tallettaminen_mongo_db_tietokantaan#mongo-db)/[English](https://fullstackopen.com/en/part3/saving_data_to_mongo_db#mongo-db)) or locally on your computer ([Instructions in English](https://www.mongodb.com/docs/manual/administration/install-community))
2. Inside the `frontend` directory, install the dependencies using the command `npm install`
3. Pick one:
    - Run the frontend in the development mode:
        1. Inside the `frontend` directory, run the development server for the frontend using the command `npm run dev` (By default the development server stays running in the foreground, so you have to either use another terminal for the rest of the commands or use the suitable command/syntax in your system to run the development server in the background. The development server can be stopped by using the command `q`)
    - Run the frontend in the production mode:
        1. Inside the `frontend` directory, build the frontend using the command `npm run build`
        2. Move the `frontend/dist` directory to `backend/client` (If the directory `backend/client` does not exist yet create it now)
4. Inside the `backend` directory, create a `.env` file and set the `MONGODB_URI` environment variable to the MongoDB URI of your MongoDB server there (If you are using MongoDB Atlas, see the instructions in step 1 to find your MongoDB URI. If you are running MongoDB locally, the MongoDB URI is by default `mongodb://localhost:27017`).
5. Add `SECRET_KEY=<your-secret-key>` into `.env` file. Replace `<your-secret-key>` with a secret key of your choice. This is needed for JWT authentication to work.
6. Inside the `backend` directory, start the app with the command `go run .` (the app can be stopped by pressing Ctrl+C in the terminal where it was started)
7. If you chose to run the frontend in the development mode, the application is now running at http://localhost:5173. If you chose to run it in the production mode, the application is running at http://localhost:5001.

## Running unit tests for the backend

1. Go to the `backend` directory
2. Make sure you have a `.env` file as above
3. Install Mockery if you haven't already ([instructions here](https://vektra.github.io/mockery/latest/installation)). (Do not use the version in the Ubuntu package manager, because it's not up-to-date.)
4. Generate the mocks using the command `mockery`. (You might need to use the full path (e.g. `/home/user/go/bin/mockery`) depending on how you installed Mockery.)
5. Run unit tests using the command `go test ./...`

## Running unit tests for the frontend

1. Go to the `frontend` directory
2. Install dependencies using `npm install`
3. Run units tests using the command `npm test`

## Running End-to-End tests

[Instructions for running E2E tests](documentation/e2e_tests.md)
