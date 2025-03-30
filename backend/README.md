# AI Prompt Gallery Backend

This is the backend API for the AI Prompt Gallery platform, built with Go and Supabase.

## Features

- User authentication (signup/login) via JWT
- CRUD operations for AI prompts
- Categories management
- User tracking
- Integration with Supabase for database operations

## Project Structure

```
backend/
├── database/       # Database interaction code
├── handlers/       # HTTP request handlers
├── middleware/     # Middleware functions
├── models/         # Data models
├── .env.example    # Example environment variables
├── .gitignore      # Git ignore file
├── go.mod          # Go module file
├── go.sum          # Go module checksums
├── main.go         # Main application file
└── README.md       # This file
```

## API Endpoints

### Authentication

- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Login a user
- `GET /api/auth/validate` - Validate a JWT token

### Users

- `GET /api/user` - Get current user information
- `PUT /api/user` - Update user information

### Prompts

- `GET /api/prompts` - List prompts with optional filtering
- `GET /api/prompts/:id` - Get a specific prompt
- `POST /api/prompts` - Create a new prompt
- `PUT /api/prompts/:id` - Update a prompt
- `DELETE /api/prompts/:id` - Delete a prompt

### Categories

- `GET /api/categories` - List all categories
- `GET /api/categories/:id` - Get a specific category
- `POST /api/categories` - Create a new category
- `PUT /api/categories/:id` - Update a category
- `DELETE /api/categories/:id` - Delete a category

## Setup

1. Clone the repository
2. Copy `.env.example` to `.env` and fill in your configuration
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Run the server:
   ```bash
   go run main.go
   ```
   
## Environment Variables

- `PORT` - Port to run the server on (default: 8080)
- `JWT_SECRET` - Secret key for JWT token generation
- `SUPABASE_URL` - Your Supabase project URL
- `SUPABASE_KEY` - Your Supabase API key

## Database Schema (Supabase)

### users
- id (uuid, primary key)
- email (text, unique)
- name (text)
- password (text, hashed)
- created_at (timestamp)
- updated_at (timestamp)

### prompts
- id (uuid, primary key)
- title (text)
- content (text)
- category_id (uuid, foreign key)
- user_id (uuid, foreign key)
- image_url (text)
- tags (text array)
- created_at (timestamp)
- updated_at (timestamp)
- likes_count (integer)
- views_count (integer)
- is_published (boolean)

### categories
- id (uuid, primary key)
- name (text)
- description (text)
- created_at (timestamp)
- updated_at (timestamp)