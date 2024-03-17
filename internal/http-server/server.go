package http_server

import (
	"fmt"
	_ "github.com/Max425/film-library.git/docs"
	"github.com/Max425/film-library.git/internal/comfig"
	"github.com/Max425/film-library.git/internal/constants"
	"github.com/Max425/film-library.git/internal/http-server/handler"
	"github.com/Max425/film-library.git/internal/repository"
	"github.com/Max425/film-library.git/internal/service"
	"github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"net/http"
)

type Service interface {
	handler.FilmService
}

func NewHttpServer(log *zap.Logger, postgres config.PostgresConfig, listenAddr string) (*http.Server, error) {
	// connect to db
	dbConnect, err := repository.NewPostgresDB(postgres)
	if err != nil {
		return nil, err
	}

	// create all repositories
	repositories := repository.NewRepository(dbConnect, log)

	// create all services
	services := service.NewService(repositories, log)

	h := handler.NewHandler(services, log)

	mux := http.NewServeMux()

	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", constants.Host)),
	))

	// Actors endpoints
	mux.HandleFunc("/api/actors", h.Use(h.CreateActor))
	mux.HandleFunc("/api/actors/{id}", h.Use(h.UpdateActor))
	mux.HandleFunc("/api/actors/{id}", h.Use(h.DeleteActor))
	mux.HandleFunc("/api/actors", h.Use(h.GetAllActors))

	// Films endpoints
	mux.HandleFunc("/api/films", h.Use(h.CreateFilm))
	mux.HandleFunc("/api/films/{id}", h.Use(h.UpdateFilm))
	mux.HandleFunc("/api/films/{id}", h.Use(h.DeleteFilm))
	mux.HandleFunc("/api/films", h.Use(h.GetAllFilms))

	return &http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}, nil
}
