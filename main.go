package main

import (
	USERS_HANDLER "auxilium-be/api/handler/users"
	"auxilium-be/infrastructure/database"
	USERS_REPO "auxilium-be/infrastructure/repository/users"
	"auxilium-be/pkg/utils"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database
	postgres, err := database.NewDatabase()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// Repository
	ur := USERS_REPO.NewUsersRepository(postgres)
	//pr := POSTS_REPO.NewPostsRepository(postgres)

	// Controller
	usersHandler := USERS_HANDLER.ControllerHandler(ur)
	//postsHandler := POSTS_HANDLER.ControllerHandler(pr)

	// JWT
	utils.InitJWT()

	// Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("welcome"))
			})
			r.Route("/users", func(r chi.Router) {
				r.Post("/register", usersHandler.CreateUsers)
				r.Post("/", usersHandler.Login)
			})
			r.Route("/posts", func(r chi.Router) {
				r.Use(jwtauth.Verifier(utils.TokenAuth))
				r.Use(jwtauth.Authenticator)
				//r.Post("/", postsHandler.CreatePosts)
			})
		})
	})
	http.ListenAndServe(":3000", r)
}
