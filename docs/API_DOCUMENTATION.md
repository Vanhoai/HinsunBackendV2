# API Documentation

## Overview
This document provides comprehensive API documentation for the Hinsun Backend V2 API endpoints including Blog, Project, Account, and Experience management.

## Base URL
```
/api/v1
```

## Authentication
Most endpoints require JWT authentication. Include the token in the Authorization header:
```
Authorization: Bearer <your_jwt_token>
```

---

## Experience API

### Get All Experiences
**GET** `/api/v1/experiences`

Retrieves all experiences.

**Response:**
```json
{
  "success": true,
  "message": "Experiences retrieved successfully",
  "data": [
    {
      "id": "uuid",
      "orderIdx": 0,
      "position": "Software Engineer",
      "company": "Tech Company",
      "location": "San Francisco, CA",
      "technologies": ["Go", "React", "PostgreSQL"],
      "responsibilities": ["Develop features", "Code review"],
      "period": "2020 - Present",
      "extra": {},
      "createdAt": 1234567890,
      "updatedAt": 1234567890
    }
  ]
}
```

### Get Experience by ID
**GET** `/api/v1/experiences/{id}`

Retrieves a specific experience by ID.

**Response:** Same as single experience object above.

### Create Experience
**POST** `/api/v1/experiences`

Creates a new experience entry.

**Request Body:**
```json
{
  "orderIdx": 0,
  "position": "Software Engineer",
  "company": "Tech Company",
  "location": "San Francisco, CA",
  "technologies": ["Go", "React", "PostgreSQL"],
  "responsibilities": ["Develop features", "Code review"],
  "period": "2020 - Present"
}
```

**Validation Rules:**
- `orderIdx`: Required, min: 0, max: 100
- `position`: Required, min: 2, max: 100
- `company`: Required, min: 2, max: 100
- `location`: Required, min: 2, max: 100
- `period`: Required, min: 2, max: 100
- `technologies`: Required array, each item min: 1, max: 50
- `responsibilities`: Required array, each item min: 5, max: 500

**Response:** Returns created experience object with status 201.

### Update Experience
**PUT** `/api/v1/experiences/{id}`

Updates an existing experience.

**Request Body:** Same as Create Experience

**Response:** Returns updated experience object.

### Delete Experience
**DELETE** `/api/v1/experiences/{id}`

Deletes a specific experience.

**Response:**
```json
{
  "success": true,
  "message": "Experience deleted successfully",
  "data": {
    "rowsAffected": 1,
    "payload": "experience-id"
  }
}
```

### Delete Multiple Experiences
**DELETE** `/api/v1/experiences?ids=id1,id2,id3`

Deletes multiple experiences by IDs.

**Query Parameters:**
- `ids`: Comma-separated list of experience IDs

**Response:**
```json
{
  "success": true,
  "message": "Experiences deleted successfully",
  "data": {
    "rowsAffected": 3,
    "payload": ["id1", "id2", "id3"]
  }
}
```

---

## Blog API

### Get All Blogs
**GET** `/api/v1/blogs`

Retrieves all blog posts.

**Response:**
```json
{
  "success": true,
  "message": "Blogs retrieved successfully",
  "data": [
    {
      "ID": "uuid",
      "AuthorID": "uuid",
      "Categories": ["Technology", "Programming"],
      "Name": "Introduction to Go",
      "Description": "A comprehensive guide to Go programming",
      "IsPublished": true,
      "Markdown": "# Blog content...",
      "Favorites": 42,
      "Views": 1000,
      "EstimatedReadTimeSeconds": 300,
      "CreatedAt": 1234567890,
      "UpdatedAt": 1234567890
    }
  ]
}
```

### Get Blog by ID
**GET** `/api/v1/blogs/{id}`

Retrieves a specific blog post by ID.

**Response:** Same as single blog object above.

### Get Blogs by Author
**GET** `/api/v1/blogs/author/{authorId}`

Retrieves all blog posts by a specific author.

**Response:** Array of blog objects.

### Create Blog
**POST** `/api/v1/blogs`

Creates a new blog post.

**Request Body:**
```json
{
  "authorId": "uuid",
  "categories": ["Technology", "Programming"],
  "name": "Introduction to Go",
  "description": "A comprehensive guide to Go programming",
  "markdown": "# Blog content...",
  "isPublished": true,
  "estimatedReadTimeSeconds": 300
}
```

**Validation Rules:**
- `authorId`: Required UUID
- `categories`: Required, min: 1, max: 5 items, each item min: 2, max: 50
- `name`: Required, min: 2, max: 100
- `description`: Required, min: 10, max: 300
- `markdown`: Required, min: 10
- `estimatedReadTimeSeconds`: Required, min: 0

**Response:** Returns created blog object with status 201.

### Update Blog
**PUT** `/api/v1/blogs/{id}`

Updates an existing blog post.

**Request Body:**
```json
{
  "categories": ["Technology", "Programming"],
  "name": "Introduction to Go",
  "description": "A comprehensive guide to Go programming",
  "markdown": "# Blog content...",
  "isPublished": true,
  "estimatedReadTimeSeconds": 300
}
```

**Response:** Returns updated blog object.

### Delete Blog
**DELETE** `/api/v1/blogs/{id}`

Deletes a specific blog post.

**Response:** Returns deletion result with rowsAffected.

### Delete Multiple Blogs
**DELETE** `/api/v1/blogs?ids=id1,id2,id3`

Deletes multiple blog posts by IDs.

**Query Parameters:**
- `ids`: Comma-separated list of blog IDs

**Response:** Returns deletion result with rowsAffected and array of IDs.

### Increment Blog Views
**POST** `/api/v1/blogs/{id}/views`

Increments the view count for a blog post.

**Response:**
```json
{
  "success": true,
  "message": "Blog views incremented successfully",
  "data": null
}
```

### Increment Blog Favorites
**POST** `/api/v1/blogs/{id}/favorites`

Increments the favorite count for a blog post.

**Response:**
```json
{
  "success": true,
  "message": "Blog favorites incremented successfully",
  "data": null
}
```

### Decrement Blog Favorites
**DELETE** `/api/v1/blogs/{id}/favorites`

Decrements the favorite count for a blog post.

**Response:**
```json
{
  "success": true,
  "message": "Blog favorites decremented successfully",
  "data": null
}
```

---

## Project API

### Get All Projects
**GET** `/api/v1/projects`

Retrieves all projects.

**Response:**
```json
{
  "success": true,
  "message": "Projects retrieved successfully",
  "data": [
    {
      "ID": "uuid",
      "Name": "Awesome Project",
      "Description": "A description of the awesome project",
      "Github": "https://github.com/user/repo",
      "Tags": ["Go", "React", "Docker"],
      "Markdown": "# Project Details...",
      "CreatedAt": 1234567890,
      "UpdatedAt": 1234567890
    }
  ]
}
```

### Get Project by ID
**GET** `/api/v1/projects/{id}`

Retrieves a specific project by ID.

**Response:** Same as single project object above.

### Create Project
**POST** `/api/v1/projects`

Creates a new project.

**Request Body:**
```json
{
  "name": "Awesome Project",
  "description": "A description of the awesome project",
  "github": "https://github.com/user/repo",
  "tags": ["Go", "React", "Docker"],
  "markdown": "# Project Details..."
}
```

**Validation Rules:**
- `name`: Required, min: 2, max: 100
- `description`: Required, min: 10, max: 500
- `github`: Required, must be a valid URL
- `tags`: Required, min: 1, max: 5 items, each item min: 2, max: 50
- `markdown`: Required, min: 10

**Response:** Returns created project object with status 201.

### Update Project
**PUT** `/api/v1/projects/{id}`

Updates an existing project.

**Request Body:** Same as Create Project

**Response:** Returns updated project object.

### Delete Project
**DELETE** `/api/v1/projects/{id}`

Deletes a specific project.

**Response:** Returns deletion result with rowsAffected.

### Delete Multiple Projects
**DELETE** `/api/v1/projects?ids=id1,id2,id3`

Deletes multiple projects by IDs.

**Query Parameters:**
- `ids`: Comma-separated list of project IDs

**Response:** Returns deletion result with rowsAffected and array of IDs.

---

## Account API

### Get Account by ID
**GET** `/api/v1/accounts/{id}`

Retrieves a specific account by ID.

**Response:**
```json
{
  "success": true,
  "message": "Account retrieved successfully",
  "data": {
    "id": "uuid",
    "name": "John Doe",
    "email": {
      "value": "john@example.com"
    },
    "emailVerified": false,
    "isActive": true,
    "avatar": "https://example.com/avatar.jpg",
    "bio": "Software Engineer",
    "createdAt": 1234567890,
    "updatedAt": 1234567890
  }
}
```

---

## Error Responses

All endpoints may return the following error responses:

### 400 Bad Request
```json
{
  "success": false,
  "message": "Validation failed",
  "error": {
    "code": "VALIDATION_FAILURE",
    "details": {}
  }
}
```

### 404 Not Found
```json
{
  "success": false,
  "message": "Resource not found",
  "error": {
    "code": "NOT_FOUND",
    "details": {}
  }
}
```

### 409 Conflict
```json
{
  "success": false,
  "message": "Resource already exists",
  "error": {
    "code": "CONFLICT",
    "details": {}
  }
}
```

### 500 Internal Server Error
```json
{
  "success": false,
  "message": "Internal server error",
  "error": {
    "code": "INTERNAL_FAILURE",
    "details": {}
  }
}
```

---

## Domain Constraints

### Experience
- Maximum 20 technologies
- Maximum 5 responsibilities
- Position, Company, Location: max 100 characters
- OrderIdx must be unique

### Blog
- Maximum 5 categories
- Name: max 100 characters (must be unique)
- Description: max 300 characters
- At least 1 category required

### Project
- Maximum 5 tags
- Name: max 100 characters (must be unique)
- Description: max 500 characters
- At least 1 tag required
- Github must be a valid URL

### Account
- Name: max 50 characters
- Bio: max 160 characters
- Email must be unique

---

## Notes

1. All timestamps are in Unix nanoseconds format.
2. UUIDs use version 7 (time-ordered) for better database performance.
3. Soft deletes are implemented - deleted records have a non-null `DeletedAt` field.
4. String arrays in PostgreSQL use the `text[]` type.
5. Blog content and Project details support full Markdown formatting.
6. Authentication is handled separately via the `/api/v1/auth` endpoints.

---

## Database Schema

### experiences
- `id` (UUID, primary key)
- `order_idx` (int, indexed)
- `position` (varchar(100))
- `company` (varchar(100))
- `location` (varchar(100))
- `technologies` (text[])
- `responsibilities` (text[])
- `period` (varchar(100))
- `extra` (jsonb)
- `created_at` (bigint)
- `updated_at` (bigint)
- `deleted_at` (bigint, nullable, indexed)

### blogs
- `id` (UUID, primary key)
- `author_id` (UUID, indexed)
- `categories` (text[])
- `name` (varchar(100), unique indexed)
- `description` (varchar(300))
- `is_published` (boolean)
- `markdown` (text)
- `favorites` (bigint)
- `views` (bigint)
- `estimated_read_time_seconds` (bigint)
- `created_at` (bigint)
- `updated_at` (bigint)
- `deleted_at` (bigint, nullable, indexed)

### projects
- `id` (UUID, primary key)
- `name` (varchar(100), unique indexed)
- `description` (varchar(500))
- `github` (varchar(255))
- `tags` (text[])
- `markdown` (text)
- `created_at` (bigint)
- `updated_at` (bigint)
- `deleted_at` (bigint, nullable, indexed)

### accounts
- `id` (UUID, primary key)
- `name` (varchar(100))
- `email` (varchar(100), unique indexed)
- `email_verified` (boolean)
- `is_active` (boolean)
- `password` (varchar(255))
- `avatar` (varchar(255))
- `bio` (text)
- `created_at` (bigint)
- `updated_at` (bigint)
- `deleted_at` (bigint, nullable, indexed)

---

## Migration

To enable auto-migration, uncomment the line in `adapters/shared/databases/postgres_database.go`:

```go
gormDB.AutoMigrate(&models.ExperienceModel{}, &models.AccountModel{}, &models.BlogModel{}, &models.ProjectModel{})
```

This will automatically create or update the database schema based on the model definitions.