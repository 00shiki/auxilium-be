package main

import (
	POSTS_HANDLER "auxilium-be/api/handler/posts"
	USERS_HANDLER "auxilium-be/api/handler/users"
	"auxilium-be/entity/responses"
	"auxilium-be/infrastructure/database"
	POSTS_REPO "auxilium-be/infrastructure/repository/posts"
	USERS_REPO "auxilium-be/infrastructure/repository/users"
	"auxilium-be/pkg/storage"
	"auxilium-be/pkg/utils"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"
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
	usersHandler := USERS_HANDLER.ControllerHandler(ur, pr)
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
			r.With(jwtauth.Verifier(utils.TokenAuth)).With(jwtauth.Authenticator).Post("/upload", func(w http.ResponseWriter, r *http.Request) {
				err := r.ParseMultipartForm(10 << 20)
				if err != nil {
					render.Render(w, r, &responses.Response{
						Code:    http.StatusBadRequest,
						Message: fmt.Sprintf("parse: %v", err.Error()),
					})
					return
				}

				_, claims, errClaims := jwtauth.FromContext(r.Context())
				if errClaims != nil {
					render.Render(w, r, &responses.Response{
						Code:    http.StatusUnauthorized,
						Message: fmt.Sprintf("claims: %v", errClaims.Error()),
					})
					return
				}

				now := time.Now()
				exp := claims["exp"].(time.Time)
				if exp.Unix() < now.Unix() {
					render.Render(w, r, &responses.Response{
						Code:    http.StatusUnauthorized,
						Message: "token expired",
					})
					return
				}

				file, header, errImage := r.FormFile("image")
				if errImage != nil {
					render.Render(w, r, &responses.Response{
						Code:    http.StatusBadRequest,
						Message: fmt.Sprintf("parse: %v", errImage.Error()),
					})
					return
				}

				imageURL, errUpload := storage.ClientInit().UploadToBucket(file, header.Filename)
				if errUpload != nil {
					render.Render(w, r, &responses.Response{
						Code:    http.StatusInternalServerError,
						Message: fmt.Sprintf("upload: %v", errUpload.Error()),
					})
					return
				}

				render.Render(w, r, &responses.Response{
					Code: http.StatusOK,
					Data: imageURL,
				})
			})
			r.Route("/users", func(r chi.Router) {
				r.Post("/register", usersHandler.CreateUsers)
				r.Post("/", usersHandler.Login)
				r.With(jwtauth.Verifier(utils.TokenAuth)).With(jwtauth.Authenticator).Get("/detail", usersHandler.DetailUser)
				r.With(jwtauth.Verifier(utils.TokenAuth)).With(jwtauth.Authenticator).Post("/update", usersHandler.UpdateUser)
			})
			r.Route("/posts", func(r chi.Router) {
				r.Use(jwtauth.Verifier(utils.TokenAuth))
				r.Use(jwtauth.Authenticator)
				r.Post("/", postsHandler.CreatePost)
				r.Get("/", postsHandler.ListPosts)
				r.Route("/{postID}", func(r chi.Router) {
					r.Get("/", postsHandler.DetailPost)
					r.Post("/comment", postsHandler.CreateComment)
				})
			})
		})
	})
	http.ListenAndServe(":"+port, r)
}
