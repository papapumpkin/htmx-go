# Go + htmx Fullstack CRUD Application

This is a basic fullstack web application built with Go and htmx. The application demonstrates a CRUD (Create, Read, Update, Delete) form where users can submit fields via an interface, and the data is stored in a PostgreSQL database. It is designed to be a simple starting point that can be extended to a more complex application in the future.

## Stack Overview

- [Go](https://golang.org/doc/install): Backend logic and routing.
- [PostgreSQL](https://www.postgresql.org/download/): Relational database used to store data.
- [Air](https://github.com/cosmtrek/air): Live reloading for Go applications when developing locally.
- [GORM](https://gorm.io/index.html): An ORM library for Go, used to interact with the PostgreSQL database and model database objects.
- [Templ](https://github.com/a-h/templ): A templating engine to render HTML on the server side.
- [htmx](https://htmx.org/): For making HTTP requests from HTML without JavaScript.
- [Atlas](https://atlasgo.io/): Managing and applying database migrations.
- [Goose](https://github.com/pressly/goose): Managing and applying database migrations.

## Features

- **CRUD Operations**: Basic operations to create, read, update, and delete records via a form interface.
- **Live-reloading Local Development**: Start up production replica with Air, allowing you to see code changes reflected instantly in your running application on your local machine.
- **Server-side Templating**: Dynamic HTML rendering with Templ keeps the backend and frontend tightly coupled. This whole application could be packaged as a single Container and deployed to Kubernetes or similar.
- **Minimal JavaScript**: Uses htmx to simplify frontend interaction without writing custom JavaScript. Not the right choice for all web applications but in this simple CRUD app it reduces cost of maintainance.

### Why Use a Repository Pattern?

The **Repository Pattern** is a way of separating concerns, in this case, the data access logic and the business logic. This improves:

- **Abstraction**: Specifically, the underlying database implementation. Changes to one layer of my application do not affect the other, making it simpler to update over time.
- **Testability**: The pattern makes it easy to mock the repository layer during unit tests, eliminating the need for a live database connection or docker containers in CI/CD or local testing.

Building applications in layers like this (e.g., separating business logic, data access, and presentation) is considered a best practice because it:

- **Increases Flexibility**: Makes it easier to change one part of an application in isolation (e.g., swapping GORM for Gin) without affecting other parts.
- **Improves Scalability**: As the application grows, maintaining and adding new features becomes easier, as each layer can evolve independently.
- **Enhances Maintainability**: By isolating concerns, it becomes easier to understand, debug, and maintain the codebase.

### Why Use Atlas and Goose for Database Migrations?

**Atlas** is a tool for managing database migrations, often described as "Terraform for databases." It allows you to define the desired state of your database schema, and it plans and applies the necessary changes (in SQL) to migrate your database to that state. Atlas supports two primary approaches to migration:

- **Declarative Approach**: You specify the desired state of the database schema, and Atlas plans and applies changes to bring the database to that state automatically.
- **Versioned Approach**: Changes to the database schema are stored as versioned SQL files, where each version contains the SQL statements required to move the database from its current state to the desired state. Atlas automatically generates these versioned migration files, making it easy to track and review schema changes over time. This approach integrates well with version control systems, allowing teams to review database migrations before applying them.

In this project, I've chosen the **versioned approach** because each time a schema change is made, Atlas generates a new SQL migration file, which can be checked into version control and reviewed. This helps maintain a clear history of all changes applied to the database over time, enhancing compliance adherance.

**Goose** is another migration tool that supports versioned migrations. Goose allows for two types of migration files: SQL files for schema changes and Go files for more complex data migrations. While Goose doesn’t support automatic migrations like Atlas, it offers flexibility by allowing both schema and data migrations to be performed in an ordered, controlled manner.

For this project, we use both tools:
- **Atlas** is used to detect schema changes, plan the required operations, and generate versioned SQL migration files.
- **Goose** is used to handle both schema migrations (using the SQL files generated by Atlas) and data migrations (using Go migration files). This ensures a clear and organized process for managing both schema and data changes in the correct order.

By leveraging both **Atlas** and **Goose**, we ensure that the database evolves predictably and systematically, while allowing the team to review, version, and control all changes to the schema and data. This combination gives us the best of both worlds: automated planning for schema changes and flexible, manually-controlled data migrations.

In my opinion, downgrading database schemas in production can lead to data corruption and loss. Thus, I usually disable this feature to avoid accidents. You should always manually validate auto-generated SQL to make sure that it doesn't do silly things, like delete a columns before recreating it with a new datatype that was specified in your ORM. Atlas outputs SQL into `internals/db/migrations` which can be manually adjusted before running this in production. Also, this migration feature can be decoupled from the main application and run as its own docker container.

## Getting Started

### Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/papapumpkin/go-htmx.git
   cd go-htmx-crud
   ```

2. Install dependencies:

  ```bash
  go mod tidy
  ```

3. Start Postgres (docker-compose):

  ```bash
  docker compose up -d
  ```

4. Start the application with Air (note: this runs a database migration):

  ```bash
  air -c .air.toml
  ```

5. Open `localhost:8081` in browser
