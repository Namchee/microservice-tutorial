package repository

import (
	"context"
	"database/sql"

	"github.com/Namchee/microservice-tutorial/post/entity"
)

const (
	getAllQuery  = "SELECT * FROM post;"
	getByIdQuery = "SELECT * FROM post WHERE post.id = $1;"
	createQuery  = "INSERT INTO post (text, user) VALUES ($1, $2) RETURNING id;"
)

type pgPostRepository struct {
	db *sql.DB
}

func NewPgPostRepository(db *sql.DB) PostRepository {
	return &pgPostRepository{
		db: db,
	}
}

func (repository *pgPostRepository) GetPosts(ctx context.Context, pagination *entity.Pagination) ([]*entity.Post, error) {
	rows, err := repository.db.QueryContext(
		ctx,
		getAllQuery,
	)

	if err != nil {
		return nil, err
	}

	var posts []*entity.Post

	for rows.Next() {
		var id int
		var text string
		var user int

		rows.Scan(&id, &text, &user)

		posts = append(posts, &entity.Post{
			Id:   id,
			Text: text,
			User: user,
		})
	}

	return posts, nil
}

func (repository *pgPostRepository) GetPostById(ctx context.Context, queryId int) (*entity.Post, error) {
	row := repository.db.QueryRowContext(
		ctx,
		getByIdQuery,
		queryId,
	)

	var id int
	var text string
	var user int

	switch err := row.Scan(&id, &text, &user); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &entity.Post{
			Id:   id,
			Text: text,
			User: user,
		}, nil
	default:
		panic(err)
	}
}

func (repository *pgPostRepository) CreatePost(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	var id int

	err := repository.db.QueryRowContext(
		ctx,
		createQuery,
		post.Text,
		post.User,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &entity.Post{
		Id:   id,
		Text: post.Text,
		User: post.User,
	}, nil
}
