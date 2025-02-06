# Slowers-App

Project done for the Software Lab course (TKT20007) at the University of Helsinki. Helps connecting local flower growers to local flower shops in the spirit of the [Slow Flowers](https://en.wikipedia.org/wiki/Slow_Flowers) movement.

- [Staging server](https://slowers.ext.ocp-test-0.k8s.it.helsinki.fi)
- [Product backlog](https://github.com/orgs/SlowersTeamSpring2025/projects/1/views/1)
- [Sprint backlogs](https://github.com/orgs/SlowersTeamSpring2025/projects?query=is%3Aopen+Sprint)
- [Coding conduct](documentation/coding_conduct_spring2025.md)
- [Definition of done](documentation/definition_of_done.md)

## Starting the application locally

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