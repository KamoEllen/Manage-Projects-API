# Project Management API

## Overview

This API allows you to manage projects with CRUD operations. You can create, read, update, and delete projects. Each project has a name, description, due date, and status.

## Database Schema

### Projects Table

The `projects` table is used to store project information. Below is the schema for the `projects` table:

```sql
CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,           -- Unique identifier for the project
    name TEXT NOT NULL,              -- Name of the project
    description TEXT,                -- Detailed description of the project
    due_date TIMESTAMP,              -- Due date for the project
    status TEXT                      -- Status of the project (e.g., "Not Started", "In Progress", "Completed")
);
```

## Table Details

- **id**: A unique identifier for each project, automatically incremented.
- **name**: The name of the project. This field is required (`NOT NULL`).
- **description**: A textual description of the project. This field is optional.
- **due_date**: The due date of the project. This field is optional and should be in a timestamp format.
- **status**: The status of the project, which can be values like "Not Started", "In Progress", or "Completed". This field is optional.

## API Endpoints

- **GET /projects**: Retrieve a list of all projects.
- **GET /projects/{id}**: Retrieve details of a specific project by ID.
- **POST /projects**: Create a new project.
- **PUT /projects/{id}**: Update an existing project by ID.
- **DELETE /projects/{id}**: Delete a project by ID.

## Example Usage

### Creating a Project

```bash
curl -X POST http://localhost:8000/projects \
  -H "Content-Type: application/json" \
  -d '{"name": "New Project", "description": "A description of the new project", "due_date": "2024-12-31T23:59:59Z", "status": "Not Started"}'
```

### Getting All Projects

```bash
curl http://localhost:8000/projects
```

### Updating a Project

```bash
curl -X PUT http://localhost:8000/projects/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Updated Project Name", "description": "Updated description", "due_date": "2024-12-31T23:59:59Z", "status": "In Progress"}'
```

### Deleting a Project

```bash
curl -X DELETE http://localhost:8000/projects/1
```

## Running the API

To run the API, use Docker to build and start the services:

```bash
docker-compose up --build
```

This will start the API server and PostgreSQL database. The API will be available at `http://localhost:8000`.

## Environment Variables

- `DATABASE_URL`: Connection string for the PostgreSQL database (e.g., `host=go_db user=postgres password=postgres dbname=postgres sslmode=disable`).


