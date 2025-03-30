# AI Prompt Gallery API Documentation

Version: 1.0  
Base URL: `http://localhost:8080/api`

## Overview

The AI Prompt Gallery API allows developers to interact with the platform's resources, including prompts, categories, users, and statistics. This document provides details on all available endpoints, their parameters, and response formats.

## Authentication

Most endpoints require authentication through a JWT token. To authenticate, include the token in the Authorization header:

```
Authorization: Bearer <jwt_token>
```

You can obtain a token by calling the `/auth/login` endpoint.

## Response Format

All responses are in JSON format and follow a consistent structure:

- For successful responses:
  - Status codes: 200 OK, 201 Created
  - Body contains requested data

- For error responses:
  - Status codes: 400 Bad Request, 401 Unauthorized, 404 Not Found, 500 Internal Server Error
  - Body contains an error message

Example error response:
```json
{
  "error": "Prompt not found"
}
```

## Endpoints

### Health Check

#### GET /health

Checks the API server status and configuration.

**Response:**

```json
{
  "status": "ok",
  "message": "AI Prompt Gallery API is running",
  "supabase_url": "https://example.supabase.co",
  "supabase_key_set": true,
  "version": "1.0"
}
```

### Prompts

#### GET /prompts

Returns a list of all prompts.

**Query Parameters:**

| Parameter | Type   | Description                            |
|-----------|--------|----------------------------------------|
| limit     | int    | Maximum number of results (default: 20) |
| offset    | int    | Offset for pagination (default: 0)      |
| sort      | string | Sort field (created_at, likes_count, views_count) |
| order     | string | Sort order (asc or desc)               |

**Response:**

```json
[
  {
    "id": "prompt-uuid",
    "title": "Creative Story Prompt",
    "content": "Write a story about a world where dreams become physical objects when people wake up.",
    "category_id": "category-uuid",
    "category_name": "Creative Writing",
    "user_id": "user-uuid",
    "user_name": "DreamWeaver",
    "created_at": "2025-03-01T12:00:00Z",
    "updated_at": "2025-03-01T12:00:00Z",
    "likes_count": 42,
    "views_count": 156,
    "image_url": "https://example.com/image.jpg"
  },
  // More prompts...
]
```

#### GET /prompts/:id

Returns details of a specific prompt.

**Path Parameters:**

| Parameter | Type   | Description |
|-----------|--------|-------------|
| id        | string | Prompt ID   |

**Response:**

```json
{
  "id": "prompt-uuid",
  "title": "Creative Story Prompt",
  "content": "Write a story about a world where dreams become physical objects when people wake up.",
  "category_id": "category-uuid",
  "category_name": "Creative Writing",
  "user_id": "user-uuid",
  "user_name": "DreamWeaver",
  "created_at": "2025-03-01T12:00:00Z",
  "updated_at": "2025-03-01T12:00:00Z",
  "likes_count": 42,
  "views_count": 156,
  "image_url": "https://example.com/image.jpg"
}
```

#### POST /prompts

Creates a new prompt. Requires authentication. 

**Note**: Users are limited to creating a maximum of 5 prompts per week. Attempting to exceed this limit will result in a 429 Too Many Requests response.

**Request Body:**

```json
{
  "title": "My Prompt Title",
  "content": "The prompt content goes here",
  "category_id": "category-uuid",
  "image_url": "https://example.com/image.jpg" // optional
}
```

**Response:**

```json
{
  "id": "new-prompt-uuid",
  "title": "My Prompt Title",
  "content": "The prompt content goes here",
  "category_id": "category-uuid",
  "category_name": "Category Name",
  "user_id": "user-uuid",
  "user_name": "Username",
  "created_at": "2025-03-29T12:00:00Z",
  "updated_at": "2025-03-29T12:00:00Z",
  "likes_count": 0,
  "views_count": 0,
  "image_url": "https://example.com/image.jpg"
}
```

#### PUT /prompts/:id

Updates an existing prompt. Requires authentication and ownership.

**Path Parameters:**

| Parameter | Type   | Description |
|-----------|--------|-------------|
| id        | string | Prompt ID   |

**Request Body:**

```json
{
  "title": "Updated Prompt Title",
  "content": "Updated prompt content",
  "category_id": "category-uuid",
  "image_url": "https://example.com/new-image.jpg" // optional
}
```

**Response:**

```json
{
  "id": "prompt-uuid",
  "title": "Updated Prompt Title",
  "content": "Updated prompt content",
  "category_id": "category-uuid",
  "category_name": "Category Name",
  "user_id": "user-uuid",
  "user_name": "Username",
  "created_at": "2025-03-01T12:00:00Z",
  "updated_at": "2025-03-29T12:30:00Z",
  "likes_count": 42,
  "views_count": 156,
  "image_url": "https://example.com/new-image.jpg"
}
```

#### DELETE /prompts/:id

Deletes a prompt. Requires authentication and ownership.

**Path Parameters:**

| Parameter | Type   | Description |
|-----------|--------|-------------|
| id        | string | Prompt ID   |

**Response:**

```json
{
  "message": "Prompt deleted successfully"
}
```

### Categories

#### GET /categories

Returns a list of all categories.

**Response:**

```json
[
  {
    "id": "category-uuid",
    "name": "Creative Writing",
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-01-01T00:00:00Z",
    "prompt_count": 42
  },
  // More categories...
]
```

#### GET /categories/:category/prompts

Returns all prompts in a specific category.

**Path Parameters:**

| Parameter | Type   | Description   |
|-----------|--------|---------------|
| category  | string | Category ID   |

**Query Parameters:**

| Parameter | Type   | Description                            |
|-----------|--------|----------------------------------------|
| limit     | int    | Maximum number of results (default: 20) |
| offset    | int    | Offset for pagination (default: 0)      |
| sort      | string | Sort field (created_at, likes_count, views_count) |
| order     | string | Sort order (asc or desc)               |

**Response:**

```json
[
  {
    "id": "prompt-uuid",
    "title": "Creative Story Prompt",
    "content": "Write a story about a world where dreams become physical objects when people wake up.",
    "category_id": "category-uuid",
    "category_name": "Creative Writing",
    "user_id": "user-uuid",
    "user_name": "DreamWeaver",
    "created_at": "2025-03-01T12:00:00Z",
    "updated_at": "2025-03-01T12:00:00Z",
    "likes_count": 42,
    "views_count": 156,
    "image_url": "https://example.com/image.jpg"
  },
  // More prompts in this category...
]
```

### Authentication

#### POST /auth/register

Registers a new user.

**Request Body:**

```json
{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "securepassword"
}
```

**Response:**

```json
{
  "id": "user-uuid",
  "username": "johndoe",
  "email": "john@example.com",
  "created_at": "2025-03-29T12:00:00Z"
}
```

#### POST /auth/login

Logs in a user and returns a JWT token.

**Request Body:**

```json
{
  "email": "john@example.com",
  "password": "securepassword"
}
```

**Response:**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "user-uuid",
    "username": "johndoe",
    "email": "john@example.com",
    "created_at": "2025-03-29T12:00:00Z"
  }
}
```

### Statistics

#### GET /auth/stats

Returns user and prompt statistics.

**Response:**

```json
{
  "total_users": 105,
  "active_users": 78,
  "new_users_today": 3,
  "new_users_week": 12,
  "new_users_month": 35,
  "prompts_created": 287,
  "avg_prompts_per_user": 2.7
}
```

## Error Codes

| Status Code | Description          | Example Situations                                 |
|-------------|----------------------|---------------------------------------------------|
| 400         | Bad Request          | Invalid input data, missing required fields        |
| 401         | Unauthorized         | Missing or invalid authentication token            |
| 403         | Forbidden            | Insufficient permissions (e.g., modifying others' prompts) |
| 404         | Not Found            | Resource not found (e.g., prompt, category)        |
| 429         | Too Many Requests    | Rate limit exceeded                                |
| 500         | Internal Server Error | Server-side error                                 |

## Rate Limiting

The API enforces rate limiting to prevent abuse. Limits are as follows:

- General API requests:
  - Authenticated requests: 100 requests per minute
  - Unauthenticated requests: 20 requests per minute

- Prompt Creation:
  - Maximum of 5 prompts per user per week

When rate limited, the API will respond with a 429 status code and a message explaining the limit.

## Versioning

The API version is included in the base URL. The current version is v1.

## Support

For questions or support, contact support@aipromptgallery.com.