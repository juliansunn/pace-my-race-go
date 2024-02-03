# Go API

## Description

This is a skeleton API built in Golang using SQLC and PostgreSQL. It provides a starting point for building your own API.

## Prerequisites

- ### Golang - [Install Go](https://go.dev/doc/install)
- ### Docker Desktop - [Docker Desktop Installation Instructions](https://docs.docker.com/desktop/)
- ### Golang Migrate (optional)

```
curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz
```

## Getting Started

1. Copy the `.env.sample` file and rename it to `.env`.
2. Update the environment variables in the `.env` file to match your preferences.
3. You will need to add the `.env` file to the root of your project using:
   ```touch .env```
5. Run the following command to copy the `.env.sample` file to populate your .env with the necessary vars:
   ```
   make copyenv
   ```

## Database Setup

1. Start the PostgreSQL database in a Docker container by running the following command:
   ```
   make postgres
   ```

## Running the Server

1. Start the server by running the following command:
   ```
   make server
   ```
   This will automatically apply any unapplied migrations.

## Additional Commands

1. ### Applying Migrations

   - To apply all migrations, run the following command:

   ```
   make migrateup
   ```

   - To apply one migration file up or down, run the following commands:

   ```
   make migrateup1
   make migratedown1
   ```

   - To unapply all migrations, run the following command:

   ```
   make migratedown
   ```

2. ### Adding Database Functions

   - To add functions that interact with the database, follow these steps:

   1. Add a new file or query in the `/query` subfolder.
   2. Run the following command to generate the necessary code to interact with the database:
      ```
      make sqlc
      ```
   3. Run the following command to create the necessary mocked code for tests:
      ```
      make mock
      ```

3. ### Testing

   - To run the tests, use the following command:

   ```
   make test
   ```

4. ### Building the app

- To build the application, use the following command:
  ```
  make build
  ```

## Contributing

Contributions are welcome! Please follow the [contribution guidelines](CONTRIBUTING.md).

## License

This project is licensed under the [MIT License](LICENSE).
