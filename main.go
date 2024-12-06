package main

import (
	"net/http"

	"github.com/agunfir98/gobroker/config"
	"github.com/agunfir98/gobroker/lib"
	"github.com/agunfir98/gobroker/rabbitmq"
	"github.com/agunfir98/gobroker/server"
)

type ehek struct {
	Message string `json:"message"`
}

func main() {
	cfg := config.LoadConfig()

	mb := rabbitmq.NewRabbitMq(cfg.RabbitMQURL)

	defer mb.Close()

	s := server.NewApiService(cfg.ApiPort)

	s.Use(func(ah *server.ApiHandler) {
		ah.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			msg := ehek{
				Message: "halo sekai",
			}
			w.WriteHeader(http.StatusOK)
			lib.Json(w, msg)
		})
	})

	s.Run()
}
