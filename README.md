
# Simple GO CRUD API

### Getting Started

To start the project, you need to initialize the environment and start the server. Copy the .env.example file to .env and update the values as needed. The project uses PostgreSQL as the database, so you need to have a PostgreSQL server running. You can use the docker-compose.yaml file to start a PostgreSQL server. Here is the command to do that:

```bash
docker-compose up -d
```

Next, you need to run the following command to start the server:

```bash
go run cmd/main.go
```
This will start the server at `http://localhost:8080`.

## Architecture
The project follows a modular architecture with separate packages for different functionalities. The main packages include:

- environment: Handles the loading of environment variables from .env file and provides a global Config object. See environment.go.
- database: Manages the database connection. See connection.go.
- server: Creates and starts the server. See server.go.
- authentication: Handles user authentication using JWT. See jwt.
- user: Manages user-related functionalities. See user.

Each module will have its own sub-modules and will contain the following folders
- handlers: Contains the API handlers for the module.
- domain: Contains the data models and exceptions for the module.
- repository: Contains the database operations for the module.
- service: Contains the business logic for the module.
- routes: Contains the API routes for the module.

## Features
- User Authentication: The project uses JWT for user authentication.
- CRUD Operations: The project provides CRUD operations for users.
- Middleware: The project uses middleware to authenticate users, role controlled and log requests.
- Database Connection: The project uses PostgreSQL as the database and manages the connection using the database package.
- Environment Configuration: The project loads environment variables from a .env file using the godotenv package.

## Scalability
The project can become scalable by:
- Using Docker: The project already includes a docker-compose.yaml file which sets up a PostgreSQL database. This can be extended to include other services as well.

- Implementing Microservices: The modular architecture of the project makes it easy to split different functionalities into separate microservices.
- Using a Load Balancer: A load balancer can be used to distribute traffic among multiple instances of the application.
- Database Sharding: The database can be sharded to distribute data across multiple databases to improve performance.

Please note that these are just suggestions and implementing them would require additional work, however, the project is designed to be scalable from the start to the mature stage.


## API Endpoints
The project provides the following API endpoints:

auth:
- POST /register: Register a new user. Non-Authenticated.
- POST /login: Login and get a JWT token. Non-Authenticated.
users:
- GET /users: Get all users. Requires Admin role.
- GET /profile: Get the user profile. Requires Authenticated.
- PUT /profile: Update the user profile. Requires Authenticated.
- DELETE /profile: Delete the user profile. Requires Authenticated.