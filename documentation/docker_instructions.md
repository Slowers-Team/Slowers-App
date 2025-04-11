# Docker instructions

To run whole application inside containers, build and run containers from repository root directory with command `docker compose -f docker-compose.dev.yml up`. Now you can use address http://localhost:8080 to access frontend and address http://localhost:8080/api to access backend. You can exit and close containers with Ctrl+C.

If you want to run container in detached mode, add flag `-d` to the end of the command and close containers with command `docker compose -f docker-compose.dev.yml down`.

## Run only databases inside container
Inside the repository root directory, build and run container from frontend image using command `docker compose -f dbs/docker-compose.dev.yml up`. See files in [dbs](../dbs/) directory and [backend/docker-compose.dev.yml](../backend/docker-compose.dev.yml) for settings needed to connect to the databases. You can exit and close containers with Ctrl+C.

If you get error "Error response from daemon: Conflict. The container name "/slowers-mongo-dev" is already in use by container "\<container id\>". You have to remove (or rename) that container to be able to reuse that name.", you have to remove old container with the same name. This can be done either through Docker Desktop or with `docker container rm slowers-mongo-dev`. There is another conflict with container slowers-psql-dev and that has to also be removed.

If you want to run container in detached mode, add flag `-d` to the end of the command and close containers with command `docker compose -f dbs/docker-compose.dev.yml down`.

## Run only development backend inside containers
Inside the repository root directory, build and run container from backend image using command `MONGODB_URI=... SECRET_KEY=... SQLDATABASEURI=... docker compose -f backend/docker-compose.dev.yml up`. Now you can use address http://localhost:5001/ to access backend. You can exit and close containers with Ctrl+C.

If you get error "Error response from daemon: Conflict. The container name "/slowers-backend-dev" is already in use by container...", you have to remove old container with the same name. This can be done either through Docker Desktop or with `docker container rm slowers-backend-dev`.

If you want to run container in detached mode, add flag `-d` to the end of the command and close containers with command `docker compose -f backend/docker-compose.dev.yml down`.

## Run only development frontend inside container
Inside the repository root directory, build and run container from frontend image using command `docker compose -f frontend/docker-compose.dev.yml up`. Now you can use address http://localhost:5173/ to access frontend. You can exit and close containers with Ctrl+C.

If you get error "Error response from daemon: Conflict. The container name "/slowers-frontend-dev" is already in use by container...", you have to remove old container with the same name. This can be done either through Docker Desktop or with `docker container rm slowers-frontend-dev`.

If you want to run container in detached mode, add flag `-d` to the end of the command and close containers with command `docker compose -f frontend/docker-compose.dev.yml down`.

## Inspecting MongoDB inside container

You can inspect MongoDB with command `docker exec -it slowers-mongo-dev mongosh -u root -p example` while datebase containers are running. You can exit the container with command `exit`. MongoDB data is saved to directory "mongo-data-dev", which is created inside directory [dbs](../dbs/) when the container is created for the first time.

## Inspecting PostgreSQL-database inside container

You can inspect MongoDB with command `docker exec -it slowers-psql-dev psql -U Slowers` while datebase containers are running. You can access correct database with command `\c slowers`. You can exit the container with command `\q`. Postgres-data is saved to directory "psql-data-dev", which is created inside directory [dbs](../dbs/) when the container is created for the first time.
