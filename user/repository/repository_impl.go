package repository

import (
	"context"
	"database/sql"

	"github.com/Namchee/microservice-tutorial/user/entity"
)

const (
	getAllQuery  = "SELECT * FROM user"
	getByIdQuery = "SELECT * FROM user WHERE user.id = $1"
	createQuery  = "INSERT INTO user (username, name, bio) VALUES ($1, $2, $3) RETURNING id"
)

type pgUserRepository struct {
	db *sql.DB
}

func NewPgUserRepository(db *sql.DB) *pgUserRepository {
	return &pgUserRepository{
		db: db,
	}
}

func (repo *pgUserRepository) GetUsers(ctx context.Context, _ *entity.Pagination) ([]*entity.User, error) {
	rows, err := repo.db.QueryContext(
		ctx,
		getAllQuery,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user *entity.User

		if err = rows.Scan(&user.Id, &user.Username, &user.Name, &user.Bio); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo *pgUserRepository) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	row := repo.db.QueryRowContext(
		ctx,
		getByIdQuery,
		id,
	)

	var user *entity.User

	if row != nil {
		row.Scan(&user.Id, &user.Username, &user.Name, &user.Bio)

		return user, nil
	}

	return nil, nil
}

func (repo *pgUserRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	var id int

	err := repo.db.QueryRowContext(
		ctx,
		createQuery,
		user.Username,
		user.Name,
		user.Bio,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &entity.User{
		Id:       id,
		Username: user.Username,
		Name:     user.Name,
		Bio:      user.Bio,
	}, nil
}
