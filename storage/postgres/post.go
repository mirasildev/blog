package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mirasildev/blog/storage/repo"
)

type postRepo struct {
	db *sqlx.DB
}

func NewPost(db *sqlx.DB) repo.PostStorageI {
	return &postRepo{
		db: db,
	}
}

func (pr *postRepo) Create(post *repo.Post) (*repo.Post, error) {
	query := `
		INSERT INTO posts(
			title,
			description,
			image_url,
			user_id,
			category_id
		) VALUES($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`

	row := pr.db.QueryRow(
		query,
		post.Title,
		post.Description,
		post.ImageUrl,
		post.UserID,
		post.CategoryID,
	)

	err := row.Scan(
		&post.ID,
		&post.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (pr *postRepo) Get(id int64) (*repo.Post, error) {
	var result repo.Post

	query := `
		SELECT
			id,
			title,
			description,
			image_url,
			user_id,
			category_id,
			created_at,
			updated_at,
			views_count
		FROM posts
		WHERE id=$1
	`

	row := pr.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.Title,
		&result.Description,
		&result.ImageUrl,
		&result.UserID,
		&result.CategoryID,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.ViewsCount,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (pr *postRepo) GetAll(params *repo.GetAllPostsParams) (*repo.GetAllPostsResult, error) {
	result := repo.GetAllPostsResult{
		Posts: make([]*repo.Post, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := "WHERE true"
	if params.Search != "" {
		filter += " AND title ilike '%" + params.Search + "%' "
	}

	if params.UserID != 0 {
		filter += fmt.Sprintf(" AND user_id=%d ", params.UserID)
	}

	if params.CategoryID != 0 {
		filter += fmt.Sprintf(" AND category_id=%d ", params.CategoryID)
	}

	orderBy := " ORDER BY created_at desc "
	if params.SortByDate != "" {
		orderBy = fmt.Sprintf(" ORDER BY created_at %s ", params.SortByDate)
	}

	query := `
		SELECT
			id,
			title,
			description,
			image_url,
			user_id,
			category_id,
			created_at,
			updated_at,
			views_count
		FROM posts
		` + filter + orderBy + limit

	rows, err := pr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var p repo.Post

		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Description,
			&p.ImageUrl,
			&p.UserID,
			&p.CategoryID,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.ViewsCount,
		)
		if err != nil {
			return nil, err
		}

		result.Posts = append(result.Posts, &p)
	}

	queryCount := `SELECT count(1) FROM posts ` + filter
	err = pr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (pr *postRepo) UpdatePost(post *repo.Post) (*repo.Post, error) {
	query := `
		UPDATE posts SET
			title=$1,
			description=$2,
			image_url=$3,
			user_id=$4,
			category_id=$5,
			updated_at=$6
		WHERE id=$7
		RETURNING id, title, description, image_url, user_id, category_id, created_at, updated_at, views_count
	`
	var result repo.Post
	err := pr.db.QueryRow(
		query,
		post.Title,
		post.Description,
		post.ImageUrl,
		post.UserID,
		post.CategoryID,
		post.UpdatedAt,
		post.ID,
	).Scan(
		&result.ID,
		&result.Title,
		&result.Description,
		&result.ImageUrl,
		&result.UserID,
		&result.CategoryID,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.ViewsCount,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (pr *postRepo) DeletePost(id int64) error {

	query := "DELETE FROM posts WHERE id=$1"
	result, err := pr.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsCount == 0 {
		return sql.ErrNoRows
	}

	return nil
}
