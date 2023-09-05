package repository

import (
	"belajar-go-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

// db *sql.DB adalah dependency yang dibutuhkan oleh repository
type commentRepositoryImpl struct {
	DB *sql.DB
}

// NewCommentRepository adalah constructor untuk membuat CommentRepository baru (dependency injection) -> ini akan dipanggil di service/comment_service.go (komen repository akan mengimplementasikan interface ini) -> ini akan mengembalikan object yang mengimplementasikan interface CommentRepository
func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

// insert into comments(email, comment) values (?, ?) -> ini adalah script untuk insert data ke database (script ini akan dijalankan oleh repository)
func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

// select id, email, comment from comments where id = ? limit 1 -> ini adalah script untuk mencari data di database (script ini akan dijalankan oleh repository)
func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// tidak ada
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

// select id, email, comment from comments -> ini adalah script untuk mencari semua data di database (script ini akan dijalankan oleh repository)
func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
