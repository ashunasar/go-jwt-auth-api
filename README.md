# Go JWT Auth API with PostgreSQL – Clean Architecture

---

## 📁 Folder Structure

```
go-jwt-auth-api/
│
├── cmd/
│   └── main.go                  # Entry point
│
├── config/
│   └── config.go                # Load environment and app config
│
├── controllers/
│   ├── auth_controller.go       # Signup, Login, Refresh Token logic
│   └── home_controller.go       # Protected route logic
│
├── database/
│   └── postgres.go              # PostgreSQL DB connection and queries
│
├── middleware/
│   └── validation.go            # Input validation and auth middleware
│
├── models/
│   ├── auth_body.go             # Signup/Login/Refresh request structs
│   └── users.go                 # User model
│
├── routes/
│   └── routes.go                # Route setup and initialization
│
├── utils/
│   ├── hash.go                  # Password hashing & comparison
│   ├── jwt.go                   # Access and refresh token logic
│   └── response.go              # Standard JSON response formatting
│
├── tmp/                         # Temp build artifacts/logs
│   ├── build-errors.log
│   └── main
│
├── go.mod
├── go.sum
├── local.yaml
├── .env                         # Env variables
├── go-jwt.postman_collection.json
└── README.md
```

---

## ⚙️ Tech Stack

- Golang
- PostgreSQL
- JWT (Access & Refresh Tokens)
- bcrypt for password hashing
- UUIDs for user IDs
- go-playground/validator for input validation

---

## 🧑‍💻 Features

### ➕ Signup

- Validates email format and password rules
- Checks if the email already exists
- Creates user with hashed password and UUID
- Returns Access & Refresh JWT tokens

### 🔑 Login

- Validates credentials
- Returns new JWT Access & Refresh tokens

### 🔁 Refresh Token

- Verifies existing stored refresh token
- Generates and returns new access & refresh tokens
- Updates DB with new refresh token

### 🔐 Protected Route (/api/home)

- Requires valid access token
- Returns user info

---

## 🔄 API Routes

| Method | Path                      | Description                   |
| ------ | ------------------------- | ----------------------------- |
| POST   | `/api/auth/signup`        | Register a new user           |
| POST   | `/api/auth/login`         | Log in with credentials       |
| POST   | `/api/auth/refresh-token` | Get new access/refresh tokens |
| GET    | `/api/home`               | Protected route (JWT req.)    |
| GET    | `/`                       | Hello world route             |

---

## 🧹 PostgreSQL User Table Schema

```sql
CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  refresh_token TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## 📦 Required Packages

| Purpose           | Package                                |
| ----------------- | -------------------------------------- |
| JWT               | github.com/golang-jwt/jwt/v5           |
| UUID              | github.com/google/uuid                 |
| Validation        | github.com/go-playground/validator/v10 |
| Hashing           | golang.org/x/crypto/bcrypt             |
| Env Loader        | github.com/joho/godotenv               |
| PostgreSQL Driver | github.com/lib/pq                      |

---

## ⏳ Token Expiry

- Access Token: **1 hour**
- Refresh Token: **60 days**

---

## 🛠 Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/ashunasar/go-jwt-auth-api.git
cd go-jwt-auth-api
```

### 2. Create `local.yaml`

In the root of your project, create a `local.yaml` file with the following structure:

```yaml
env: '<environment>'
db_path: '<your_postgres_connection_string>'
http_server:
  address: '<host>:<port>'
access_token_secret: '<your_access_token_secret>'
refresh_token_secret: '<your_refresh_token_secret>'
```

> Example for `db_path`: `postgres://postgres:yourpassword@localhost:5432/student_db?sslmode=disable`

---

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Run the Application

```bash
go run cmd/main.go
```

Your API should now be running on: `http://localhost:8082`

---

## 📬 Contact

Created by [Ashu Nasar](https://github.com/ashunasar) — feel free to reach out if you have questions or want to contribute!

---

> This project is perfect for learning clean architecture, authentication with JWT, Go middleware patterns, and working with PostgreSQL!
