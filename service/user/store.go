package user

import (
	"database/sql"
	"fmt"

	"github.com/lucaastorres7/ecom-api/types"
)

// Dependency Injection Here
type Store struct {
	db *sql.DB
}

// New Function for Data Layer
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// Data Methods
func (s *Store) GetUsers() ([]types.User, error) {
	var users []types.User

	rows, err := s.db.Query(`
		SELECT "id", "firstName", "lastName", "email", "createdAt"
		FROM users
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u types.User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	u := new(types.User)
	err := s.db.QueryRow(`
		SELECT "id", "firstName", "lastName", "email", "password", "createdAt"
		FROM users
		WHERE email = $1
	`, email).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	u := new(types.User)
	err := s.db.QueryRow(`
		SELECT "id", "firstName", "lastName", "email", "password", "createdAt"
		FROM users
		WHERE id = $1
	`, id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec(`
		INSERT INTO USERS (
			"firstName",
			"lastName",
			"email",
			"password"
		) VALUES (
			$1, $2, $3, $4
		)
	`, user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}

// Data Helper
func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}