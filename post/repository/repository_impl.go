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
		var post *entity.Post

		rows.Scan(&post.Id, &post.Text, &post.User)

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository *pgPostRepository) GetPostById(ctx context.Context, id int) (*entity.Post, error) {
	row := repository.db.QueryRowContext(
		ctx,
		getByIdQuery,
		id,
	)

	if row != nil {
		var post *entity.Post

		row.Scan(&post.Id, &post.Text, &post.User)

		return post, nil
	}

	return nil, nil
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
