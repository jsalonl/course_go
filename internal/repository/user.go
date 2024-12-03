package repository

import (
	"database/sql"
	"sample-api/internal/model"
	"sample-api/internal/service"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) service.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Add(user *model.User) (*model.User, error) {
	_, err := u.db.Exec("INSERT INTO users(id, name) values($1, $2)", user.ID, user.Name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) List() ([]*model.User, error) {
	var userList []*model.User
	rows, err := u.db.Query("SELECT id, name FROM users")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.ID,
			&user.Name)
		if err != nil {
			return nil, err
		}

		userList = append(userList, &user)
	}

	return userList, nil
}

func (u *userRepository) ExistByName(name string) (bool, error) {
	rows, err := u.db.Query("SELECT * FROM users WHERE name =$1", name)
	if err != nil {
		return false, err
	}
	
	if rows.Next() {
		return true, nil
	}

	return false, nil
}
