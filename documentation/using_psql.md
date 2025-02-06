# Using PostgreSQL

Using (local) Postgres-database in conjunction with MongoDB is currently possible with following steps:

1. You need local instance of Postgres running
    - One way to do this is to use Docker. There will be code which can be copied and used to set up a suitable container.
2. Add SQLDATABASEURI to your .env file(e.g. `SQLDATABASEURI=postgres://Slowers:salainensana@localhost:5432`)
3. Run backend using command `USESQL=true go run .`

## How to run Postgres-container

These instructions assume you have Docker installed and running. [Fullstack MOOC part12](https://fullstackopen.com/en/part12/) is a good place to start.

Create a new directory somewhere. Then add following files:

docker-compose.yml
```yml
services:
  slowers-dev-psql:
    image: postgres
    ports:
      - 5432:5432
    environment:
      POSTGRES_USER: Slowers
      POSTGRES_PASSWORD: salainensana
    volumes: 
      - ./init.sql:/docker-entrypoint-initdb.d/init.sql
```

init.sql
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    role VARCHAR(20) NOT NULL
)
```

While inside that directory, run command `docker compose up -d`. Database is now running in port 5432. To close the container, use command `docker compose down`.
