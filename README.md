# Slowers-App

Project done for the Software Lab course (TKT20007) at the University of Helsinki. Helps connecting local flower growers to local flower shops in the spirit of the [Slow Flowers](https://en.wikipedia.org/wiki/Slow_Flowers) movement.

- [Staging server](https://slowers.ext.ocp-test-0.k8s.it.helsinki.fi)
- [Product backlog](https://github.com/orgs/Slowers-Team/projects/3)
- [Sprint backlogs](https://github.com/orgs/Slowers-Team/projects?query=is%3Aopen+Sprint)
- [Coding conduct](https://github.com/Slowers-Team/Slowers-App/blob/main/documentation/coding_conduct.md)
- [Definition of done](https://github.com/Slowers-Team/Slowers-App/blob/main/documentation/definition_of_done.md)

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
