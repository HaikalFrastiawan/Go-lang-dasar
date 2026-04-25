package main

import (
	"restful-api/app"
	"restful-api/controller"
	"restful-api/repository"
	"restful-api/service"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

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

}