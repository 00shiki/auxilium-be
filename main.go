package main

import (
	POSTS_HANDLER "auxilium-be/api/handler/posts"
	USERS_HANDLER "auxilium-be/api/handler/users"
	"auxilium-be/infrastructure/database"
	POSTS_REPO "auxilium-be/infrastructure/repository/posts"
	USERS_REPO "auxilium-be/infrastructure/repository/users"
	"auxilium-be/pkg/utils"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	godotenv.Load(".env")

	// Database
	postgres, err := database.NewDatabase()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// Repository
	ur := USERS_REPO.NewUsersRepository(postgres)
	pr := POSTS_REPO.NewPostsRepository(postgres)

	// Controller
	usersHandler := USERS_HANDLER.ControllerHandler(ur)
	postsHandler := POSTS_HANDLER.ControllerHandler(pr, ur)

	// JWT
	utils.InitJWT()

	// Router
	port := os.Getenv("PORT")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
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
				r.Post("/", postsHandler.CreatePost)
				r.Get("/", postsHandler.ListPosts)
			})
		})
	})
	http.ListenAndServe(":"+port, r)
}
