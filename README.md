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
