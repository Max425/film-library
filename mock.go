package film_library

//go:generate mockgen -source=internal/http-server/handler/actor.go -destination=mocks/service/mock_actor.go
//go:generate mockgen -source=internal/http-server/handler/film.go -destination=mocks/service/mock_film.go
//go:generate mockgen -source=internal/http-server/handler/auth.go -destination=mocks/service/mock_auth.go
//go:generate mockgen -source=internal/service/actor.go -destination=mocks/db/mock_actor.go
//go:generate mockgen -source=internal/service/film.go -destination=mocks/db/mock_film.go
//go:generate mockgen -source=internal/service/auth.go -destination=mocks/db/mock_auth.go
