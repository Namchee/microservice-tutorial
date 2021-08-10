package repository

import (
	"context"
	"database/sql"

	"github.com/Namchee/microservice-tutorial/user/pb"
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

func (repo *pgUserRepository) GetUsers(ctx context.Context, _ *pb.GetUsersRequest) ([]*pb.User, error) {
	rows, err := repo.db.QueryContext(
		ctx,
		getAllQuery,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*pb.User
	for rows.Next() {
		var user *pb.User

		if err = rows.Scan(&user.Id, &user.Username, &user.Name, &user.Bio); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo *pgUserRepository) GetUserById(ctx context.Context, id int32) (*pb.User, error) {
	row := repo.db.QueryRowContext(
		ctx,
		getByIdQuery,
		id,
	)

	var user *pb.User

	if row != nil {
		row.Scan(&user.Id, &user.Username, &user.Name, &user.Bio)

		return user, nil
	}

	return nil, nil
}

func (repo *pgUserRepository) CreateUser(ctx context.Context, data *pb.CreateUserRequest) (*pb.User, error) {
	var id int32

	err := repo.db.QueryRowContext(
		ctx,
		createQuery,
		data.Username,
		data.Name,
		data.Bio,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:       id,
		Username: data.Username,
		Name:     data.Name,
		Bio:      data.Bio,
	}, nil
}
