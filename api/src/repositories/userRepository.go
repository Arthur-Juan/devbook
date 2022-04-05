package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

//UserRepository Como se estivesse populando uma classe em um construct
type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (userRepository UserRepository) CreateUser(user models.User) (uint64, error) {

	query := "INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)"
	stmt, err := userRepository.db.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, nil
	}

	lastIDInsert, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastIDInsert), nil
}

func (userRepository UserRepository) SearchUsers(nameOrNick string) ([]models.User, error) {

	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) //%%%s == %s
	query := "SELECT id, name, nick, email, createdAt FROM users WHERE name LIKE ? or nick LIKE ?"

	rows, err := userRepository.db.Query(query, nameOrNick, nameOrNick)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (userRepository UserRepository) SearchUserById(userId uint64) (models.User, error) {

	query := "SELECT id, name, nick, email, createdAt FROM users WHERE id = ?"

	rows, err := userRepository.db.Query(query, userId)

	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	var user models.User
	if rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (userRepository UserRepository) UpdateUser(ID uint64, user models.User) error {

	query := "UPDATE users SET name = ?, nick = ?,  email = ? WHERE id = ?"
	stmt, err := userRepository.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (userRepository UserRepository) Delete(id uint64) error {

	query := "DELETE FROM users WHERE id = ?"

	stmt, err := userRepository.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
