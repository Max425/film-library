package http_server

import (
	"fmt"
	_ "github.com/Max425/film-library.git/docs"
	"github.com/Max425/film-library.git/internal/comfig"
	"github.com/Max425/film-library.git/internal/common/constants"
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

func NewHttpServer(log *zap.Logger, postgres config.PostgresConfig, redis config.RedisConfig, listenAddr string) (*http.Server, error) {
	// connect to db
	dbConnect, err := repository.NewPostgresDB(postgres)
	if err != nil {
		return nil, err
	}

	// connect to redis
	redisClient, err := repository.NewRedisClient(redis)
	if err != nil {
		return nil, err
	}

	// create all repositories
	repositories := repository.NewRepository(dbConnect, log, redisClient)

	// create all services
	services := service.NewService(repositories, log)

	h := handler.NewHandler(services, log)

	mux := http.NewServeMux()

	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", constants.Host)),
	))
	// Auth
	mux.HandleFunc("/api/auth/login", h.UseRecoveryLogging(h.SignIn))
	mux.HandleFunc("/api/auth/logout", h.UseRecoveryLogging(h.Logout))
	mux.HandleFunc("/api/auth/sign-up", h.UseRecoveryLogging(h.SignUp))

	// Actors endpoints
	mux.HandleFunc("/api/create_actors", h.UseRecoveryLoggingAuth(h.CreateActor))
	mux.HandleFunc("/api/update_actors", h.UseRecoveryLoggingAuth(h.UpdateActor))
	mux.HandleFunc("/api/actors/", h.UseRecoveryLoggingAuth(h.DeleteActor))
	mux.HandleFunc("/api/actors", h.UseRecoveryLoggingAuth(h.GetAllActors))

	// Films endpoints
	mux.HandleFunc("/api/create_films", h.UseRecoveryLoggingAuth(h.CreateFilm))
	mux.HandleFunc("/api/update_films", h.UseRecoveryLoggingAuth(h.UpdateFilm))
	mux.HandleFunc("/api/update_films_actors/", h.UseRecoveryLoggingAuth(h.UpdateFilmActors))
	mux.HandleFunc("/api/films/", h.UseRecoveryLoggingAuth(h.DeleteFilm))
	mux.HandleFunc("/api/films", h.UseRecoveryLoggingAuth(h.GetAllFilms))

	return &http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}, nil
}
