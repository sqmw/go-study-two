package user

import "database/sql"

type User struct {
	ID   int
	Name string
	Age  int
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetUserByID(id int) (*User, error) {
	user := &User{}
	err := r.db.QueryRow("SELECT id, name, age FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return nil, err
	}
	return user, nil
}
