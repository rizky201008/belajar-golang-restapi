package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rizky201008/belajar-golang-restapi/app"
	"github.com/rizky201008/belajar-golang-restapi/controller"
	"github.com/rizky201008/belajar-golang-restapi/helper"
	"github.com/rizky201008/belajar-golang-restapi/middleware"
	"github.com/rizky201008/belajar-golang-restapi/repository"
	"github.com/rizky201008/belajar-golang-restapi/service"
	"net/http"
)

func main() {
	validate := validator.New()
	db := app.NewDb()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    ":3333",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
