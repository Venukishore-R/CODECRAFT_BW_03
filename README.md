# JWT-Based Authentication & Authorization

## Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Tech Stack](#tech-stack)
4. [Setup and Installation](#setup-and-installation)
5. [Usage](#usage)
6. [API Endpoints](#api-endpoints)
7. [License](#license)

## Introduction

This project implements a secure authentication and authorization system using JSON Web Tokens (JWT). It provides user registration, login functionality, and ensures role-based access control for endpoints.

## Features

- User registration and login endpoints.
- Password hashing using libraries like bcrypt.
- Generation of JWT tokens upon successful login.
- Protection of specific routes for authenticated users.
- Role-based access control (e.g., Admin, User, Owner).

## Tech Stack

- **Backend**: Golang
- **Database**: PostgreSQL, MongoDB
- **Authentication**: JWT
- **Password Hashing**: bcrypt

## Setup and Installation

1. Clone the repository:

```bash
git clone https://github.com/your-repo/jwt-auth-system.git
cd jwt-auth-system
```

2. Install dependencies:

```bash
go mod tidy
```

3. Set up the environment variables: Create a `.env` file with the following:

```env
DB_HOST=your-database-host
DB_USER=your-database-user
DB_PASSWORD=your-database-password
DB_NAME=your-database-name
JWT_SECRET=your-secret-key
```

4. Run the application:

```bash
make run
```

## Usage

1. **Register a user** by making a POST request to `/api/signup` with:

```json
{
  "name": "exampleUser",
  "email": "user@gmail.com",
  "password": "examplePassword",
  "age": 14,
  "role":"user"
}
```

2. **Login** to obtain a JWT token:

```json
{
  "username": "exampleUser",
  "password": "examplePassword"
}
```

3. Use the token to access protected routes by adding it to the Authorization header:

```http
Authorization: Bearer <your-token>
```

## API Endpoints

| **Endpoint**                          | **Method** | **Description**                          |
|---------------------------------------|------------|------------------------------------------|
| `/api/signup`                         | POST       | Register a new user                      |
| `/api/login`                          | POST       | Authenticate and get a JWT token         |
| `/api/user/profile`                   | GET        | Access profile of the authenticated user |
| `/api/admin/users`                    | GET        | Retrieve a list of all users (Admin only)|
| `/api/admin/user`                     | POST       | Create a new user (Admin only)           |
| `/api/admin/user`                     | PUT        | Update user details (Admin only)         |
| `/api/admin/user`                     | DELETE     | Delete a user (Admin only)               |

## License

This project is licensed under the MIT License.
