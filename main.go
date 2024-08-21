package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/rizky201008/belajar-golang-restful-api/app"
	"github.com/rizky201008/belajar-golang-restful-api/controller"
	"github.com/rizky201008/belajar-golang-restful-api/exception"
	"github.com/rizky201008/belajar-golang-restful-api/helper"
	"github.com/rizky201008/belajar-golang-restful-api/repository"
	"github.com/rizky201008/belajar-golang-restful-api/service"
	"net/http"
)

func main() {
	validate := validator.New()
	db := app.NewDb()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    ":3333",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfErr(err)
}
