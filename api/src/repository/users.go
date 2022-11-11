package repository

import (
	"api/src/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewRepositoryUsers(db *sql.DB) *Users {
	return &Users{db}
}

// Insert user to DB
func (repository Users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, nil
	}
	defer statement.Close()
	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, nil
	}

	lastIdInsert, erro := result.LastInsertId()
	if erro != nil {
		return 0, nil
	}
	return uint64(lastIdInsert), nil
}
