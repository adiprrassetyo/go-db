package repository

import (
	"belajar-go-database/entity"
	"context"
)

type CommentRepository interface {
	// komen repository akan mengimplementasikan interface ini untuk melakukan operasi CRUD ke database (insert, findbyid, findall) dan akan diimplementasikan di file comment_repository_impl.go (komen repository akan mengimplementasikan interface ini)
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}
