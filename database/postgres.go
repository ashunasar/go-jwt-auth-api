package database

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/ashunasar/go-jwt-auth-api/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func Connect(dbPath string) error {
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  name TEXT  NOT NULL,
  password TEXT NOT NULL,
  refresh_token TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`)

	if err != nil {
		return err
	}
	Db = db

	slog.Info("Connected to Database")

	return err

}
func CreateUser(user models.User) (uuid.UUID, error) {

	if user.Id == uuid.Nil {
		user.Id = uuid.New()
	}
	now := time.Now()

	query := `
		INSERT INTO users (id, email, name, password, refresh_token, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := Db.Exec(
		query,
		user.Id,
		user.Email,
		user.Name,
		user.Password,
		user.RefreshToken,
		now,
		now,
	)

	if err != nil {

		slog.Error(err.Error())
		return uuid.Nil, err

	}

	fmt.Println("id is ", user.Id)

	return user.Id, err

}

func GetUseByEmail(email string) (uuid.UUID, error) {

	var id uuid.UUID

	query := `Select Id from users where email = $1`

	err := Db.QueryRow(query, email).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with that email")
			return uuid.Nil, nil
		}
		log.Printf("Query error: %v\n", err)
		return uuid.Nil, err
	}

	return id, nil

}
