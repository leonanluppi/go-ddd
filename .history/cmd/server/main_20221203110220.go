package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/devfullcycle/catalogo-fc/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	conf, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&entity.Product{}, &entity.User{})
	// productDB := database.NewProduct(db)
	// productHandler := handlers.NewProductHandler(productDB)

	// userDB := database.NewUser(db)
	// userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/categories", postsResource{}.Routes())

	// r.Route("/products", func(r chi.Router) {
	// 	r.Use(jwtauth.Verifier(configs.TokenAuth))
	// 	r.Use(jwtauth.Authenticator)
	// 	r.Post("/", productHandler.CreateProduct)
	// 	r.Get("/", productHandler.GetProducts)
	// 	r.Get("/{id}", productHandler.GetProduct)
	// 	r.Put("/{id}", productHandler.UpdateProduct)
	// 	r.Delete("/{id}", productHandler.DeleteProduct)
	// })
	// r.Post("/users", userHandler.Create)
	// r.Post("/users/generate_token", userHandler.GetJWT)

	// r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
