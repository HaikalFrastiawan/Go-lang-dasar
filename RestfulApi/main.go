package main

import (
	"fmt"
	"net/http"
	"restful-api/app"
	"restful-api/controller"
	"restful-api/exception"
	"restful-api/helper"
	"restful-api/repository"
	"restful-api/service"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

const categoryPath = "/api/categories/:categoryId"

func main() {

	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	NewcategoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(NewcategoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindByID)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	fmt.Println("Server sedang berjalan di http://localhost:3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}