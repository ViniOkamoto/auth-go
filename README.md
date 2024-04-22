
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
The project adopts a modular architecture, facilitating scalability, maintainability, and future refactoring into microservices. This architectural approach emphasizes clear boundaries between components, enhancing extensibility and comprehension.

### Modular Structure
The main packages encompass distinct functionalities:

- environment: This package manages environment variables, leveraging a global Config object initialized from a .env file (see environment.go).
- database: Responsible for handling database connections and interactions (see connection.go).
- server: Initiates and launches the server, serving as the entry point for incoming requests (see server.go).
- authentication: This package governs user authentication using JWT (JSON Web Tokens) for secure access control (see jwt).
- user: Manages operations related to users, such as creation, modification, and retrieval.

### Modular Components
Each module is structured into sub-modules, featuring the following directories:

- handlers: Hosts API handlers responsible for processing incoming requests and generating appropriate responses.
- domain: Contains data models defining the structure of entities within the module, along with exception handling mechanisms.
- repository: Houses database operations specific to the module, encapsulating CRUD (Create, Read, Update, Delete) functionalities.
- service: Encapsulates the business logic pertaining to the module's operations, promoting separation of concerns and modularity.
- routes: Defines API routes associated with the module, establishing endpoints for client-server communication.

## Benefits of Modular Architecture
- Scalability: Modular design facilitates the addition or removal of features/modules without significant impact on other parts of the system, enabling seamless scalability.
- Maintainability: Clear module boundaries enhance code maintainability by isolating changes within specific components, minimizing the risk of unintended side effects.
- Extensibility: Well-defined interfaces between modules enable easy extension of functionalities by plugging in new components or replacing existing ones.
- Comprehension: Modular architecture promotes better understanding of system components and their interactions, simplifying troubleshooting, debugging, and onboarding of new developers.
- Microservice Readiness: The modular structure lays the groundwork for future refactoring into microservices, as each module encapsulates a coherent set of functionalities with clear boundaries.

By adhering to this modular architecture, the project ensures flexibility, maintainability, and readiness for future expansion or migration to a microservices-based architecture.

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
