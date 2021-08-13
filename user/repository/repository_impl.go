package repository

import (
	"context"
	"database/sql"

	"github.com/Namchee/microservice-tutorial/user/entity"
)

const (
	getAllQuery  = "SELECT * FROM \"user\";"
	getByIdQuery = "SELECT * FROM \"user\" WHERE id = $1;"
	createQuery  = "INSERT INTO \"user\" (username, name, bio) VALUES ($1, $2, $3) RETURNING id;"
)

type pgUserRepository struct {
	db *sql.DB
}

func NewPgUserRepository(db *sql.DB) UserRepository {
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
		var id int
		var username string
		var name string
		var bio string

		err = rows.Scan(&id, &username, &name, &bio)

		if err != nil {
			return nil, err
		}

		users = append(users, &entity.User{
			Id:       id,
			Username: username,
			Name:     name,
			Bio:      bio,
		})
	}

	return users, nil
}

func (repo *pgUserRepository) GetUserById(ctx context.Context, queryId int) (*entity.User, error) {
	row := repo.db.QueryRowContext(
		ctx,
		getByIdQuery,
		queryId,
	)

	var id int
	var username string
	var name string
	var bio string

	switch err := row.Scan(&id, &username, &name, &bio); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &entity.User{
			Id:       id,
			Username: username,
			Name:     name,
			Bio:      bio,
		}, nil
	default:
		panic(err)
	}
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
