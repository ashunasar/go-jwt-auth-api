# Go JWT Auth with PostgreSQL – Project Plan (Clean Architecture)

---

## 📁 Folder Structure

```
go-auth-jwt-postgres/
│
├── cmd/
│   └── main.go                  # Entry point
│
├── config/
│   └── config.go                # Load .env and app config
│
├── controllers/
│   └── auth_controller.go       # Signup/Login
│   └── user_controller.go       # Protected route
│
├── middleware/
│   └── auth_middleware.go       # JWT middleware
│
├── models/
│   └── user.go                  # User model definition
│
├── routes/
│   └── routes.go                # Route initialization
│
├── services/
│   └── auth_service.go          # Token generation/validation logic
│
├── utils/
│   └── hash.go                  # Password hashing
│   └── response.go              # JSON responses
│
├── database/
│   └── postgres.go              # DB connection and migration
│
├── go.mod
├── .env
└── README.md

```

---

## ⚙️ Tech Stack

- Golang
- PostgreSQL
- JWT
- bcrypt for hashing
- go-playground/validator for input validation

---

## 🧑‍💻 Features

### ➕ Signup

- Validate email format
- Validate password strength
- Check if email exists
- Create user with hashed password and UUID
- Return JWT access & refresh tokens

### 🔑 Login

- Validate credentials
- Return new access & refresh tokens

### 🔁 Refresh Token

- Verify stored refresh token
- Return new tokens
- Update DB with latest refresh token

### 🔐 Protected Route (/home)

- Access only with valid access token
- Return user's info

---

## 🧹 Database Schema (PostgreSQL)

```sql
CREATE TABLE users (
  id UUID PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  refresh_token TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## 📦 Packages to Use

| Purpose           | Package                                |
| ----------------- | -------------------------------------- |
| JWT               | github.com/golang-jwt/jwt/v5           |
| UUID              | github.com/google/uuid                 |
| Validation        | github.com/go-playground/validator/v10 |
| Hashing           | golang.org/x/crypto/bcrypt             |
| Env Config        | github.com/joho/godotenv               |
| PostgreSQL Driver | github.com/lib/pq                      |

---

## ⏳ Token Expiry

- Access Token: 1 hour
- Refresh Token: 60 days

---

## 🚀 Next Steps

1. Set up Go module and environment config
2. Build PostgreSQL connection & user migration
3. Implement Signup API
4. Add Login API
5. Add Refresh Token endpoint
6. Add JWT Middleware and Protected route
7. Clean up and document

---

> This project is perfect for learning authentication, middleware, token handling, and clean architecture in Go!
