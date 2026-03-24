package repository

import (
	"context"
	"fmt"
	mysqldb "go-mysql-database"
	"go-mysql-database/entity"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(mysqldb.GetConnection())
	ctx := context.Background()

	comment := entity.Comment{
		Email:   "test@gmail.com",
		Comment: "Ini komentar dari unit test",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id:", result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(mysqldb.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 35)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(mysqldb.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}