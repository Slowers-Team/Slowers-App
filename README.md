# Slowers-App

Repository for Software Engineering Project Autumn 2024 course.

## Starting the application locally

1. Install and start MongoDB ([See MongoDB documentation for details](https://www.mongodb.com/docs/manual/administration/install-community))
2. Inside the `frontend` directory, install the dependencies and build the frontend using the commands
```
npm install
npm run build
```
3. Move the `frontend/dist` directory to `backend/client` (If the directory `backend/client` does not exist yet create it now)
4. Inside the `backend` directory, create `.env` file with the `MONGODB_URI` environment variable set as below (though the exact MongoDB URI might vary depending on your configuration)
```
MONGODB_URI=mongodb://localhost:27017
```
5. Inside the `backend` directory, start the app with the command `go run main.go`
6. Now the application is running at http://localhost:5001. It can be stopped by pressing Ctrl+C in the terminal where it was started.
