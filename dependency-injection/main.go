package main

import (
	"fmt"
	"net/http"
	"restful-api/app"
	"restful-api/controller"
	"restful-api/exception"
	"restful-api/helper"
	"restful-api/middleware"
	"restful-api/repository"
	"restful-api/service"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

const categoryPath = "/api/categories/:categoryId"

func main() {

	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	NewcategoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(NewcategoryService)

	router := app.NewRouter(categoryController)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("Server sedang berjalan di http://localhost:3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
