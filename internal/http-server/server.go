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
	handler.CompanyService
}

func NewHttpServer(log *zap.Logger, postgres config.PostgresConfig, listenAddr string) (*http.Server, error) {
	// connect to db
	dbConnect, err := repository.NewPostgresDB(postgres)
	if err != nil {
		return nil, err
	}

	// create all repositories
	managerRepo := repository.NewRepository(dbConnect, log)

	// create all services
	managerService := service.NewService(managerRepo, log)

	r := http.NewServeMux()

	r.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", constants.Host)),
	))

	r.HandleFunc("/create_event", h.Use(h.createEvent))
	r.HandleFunc("/update_event", h.Use(h.updateEvent))
	r.HandleFunc("/delete_event", h.Use(h.deleteEvent))
	r.HandleFunc("/events_for_day", h.Use(h.getEventsForDay))
	r.HandleFunc("/events_for_week", h.Use(h.getEventsForWeek))
	r.HandleFunc("/events_for_month", h.Use(h.getEventsForMonth))

	return &http.Server{
		Addr:    listenAddr,
		Handler: router,
	}, nil
}

func (h *Handler) Use(next http.HandlerFunc) http.HandlerFunc {
	return h.panicRecoveryMiddleware(
		h.loggingMiddleware(next),
	)
}
