# Installing Swagger for Go API Documentation

This document explains how to install and use Swagger to generate API documentation for our Go backend.

## Installing swag CLI

### Option 1: Using Go Install (Recommended)

If you have Go installed on your system (which you should for a Go project), you can install `swag` using the following command:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Make sure your Go bin directory is in your PATH:
- On Windows, add `%USERPROFILE%\go\bin` to your PATH
- On Linux/Mac, add `$HOME/go/bin` to your PATH

### Option 2: Using pre-built binaries

Alternatively, you can download pre-built binaries from the GitHub releases page:
https://github.com/swaggo/swag/releases

## Verifying Installation

After installation, verify that the `swag` command is available:

```bash
swag --version
```

You should see the version of swag printed to the console.

## Generating Swagger Documentation

To generate Swagger documentation, run:

```bash
swag init -g main.go -o ./docs
```

This will:
- Parse all the Swagger annotations in your code
- Create a `docs` directory with Swagger documentation
- Generate an `index.html` file that you can use to view the documentation

## Setting Up Swagger UI in Your Application

### Step 1: Add Required Packages

```bash
go get -u github.com/swaggo/http-swagger
go get -u github.com/swaggo/files
```

### Step 2: Update your main.go file

Add the following imports:

```go
import (
    // ... other imports
    _ "your-module/docs" // This is important! Replace "your-module" with your actual module name
    httpSwagger "github.com/swaggo/http-swagger"
)
```

### Step 3: Add Swagger route

Add this to your route setup:

```go
// Swagger documentation
router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
```

### Step 4: Add General API Info in main.go

Add these annotations to your main.go file:

```go
// @title AI Prompt Gallery API
// @version 1.0
// @description API for managing AI prompts, categories, and users
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
```

## Accessing Swagger UI

After starting your application, you can access the Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

The UI will show all your API endpoints with descriptions, parameters, and response formats.
