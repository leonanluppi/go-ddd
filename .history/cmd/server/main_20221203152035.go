package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devfullcycle/catalogo-fc/config"
	database "github.com/devfullcycle/catalogo-fc/internal/infra/mysql"
	infra "github.com/devfullcycle/catalogo-fc/internal/infra/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	conf, err := config.LoadConfig("./cmd/server")
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DBUser,
		conf.DBPassword,
		conf.DBHost,
		conf.DBPort,
		conf.DBName,
	)

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
	// 	conf.DBHost,
	// 	conf.DBUser,
	// 	conf.DBPassword,
	// 	conf.DBName,
	// 	conf.DBPort,
	// )

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Mount("/categories", infra.CategoriesResource{}.NewRoutes(
		database.NewCategoryRepository(db),
	))

	// http.ListenAndServe(":8080", router)
	http.ListenAndServe(conf.WebServerPort, router)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
