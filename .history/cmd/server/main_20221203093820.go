package main

import (
	"log"
	"net/http"

	"github.com/devfullcycle/catalogo-fc/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	conf, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }
	// db.AutoMigrate(&entity.Product{}, &entity.User{})
	// productDB := database.NewProduct(db)
	// productHandler := handlers.NewProductHandler(productDB)

	// userDB := database.NewUser(db)
	// userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

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
