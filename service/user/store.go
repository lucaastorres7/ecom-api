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
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query(`
		SELECT * 
		FROM users 
		WHERE email = $1
	`, email)

	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (s *Store) CreateUser(user types.User) error {
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