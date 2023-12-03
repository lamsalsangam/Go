# ToDo List Web API

**Implement a simple CRUD API for managing ToDo items.**

**cmd/todo-api**: This is where main application code will reside. It includes the **_main.go_** file, which will initialize your HTTP server and set up routes.

**internal/database**: This directory contains the database-related code. The **_database.go_** file will have functions to interact with the POSTGRES database, such as insert, retrieve, update, and delete operations.

**internal/handlers**: Here, define HTTP request handlers. **_todo_handler.go_** will contain functions that handle CRUD operations for ToDo items.

**pkg/middleware**: This directory is for middleware functions. Have **_logging.go_** file for logging requests, and optionally, we can add another file for error handling middleware.

**tests**: This directory is for our unit tests. **_database_test.go_** will contain tests for your database functions, and **_api_test.go_** will include tests for our API's CRUD operations.
