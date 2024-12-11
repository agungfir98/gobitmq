package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/agunfir98/gobroker/config"
	"github.com/agunfir98/gobroker/lib"
	"github.com/agunfir98/gobroker/rabbitmq"
	"github.com/agunfir98/gobroker/server"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	cfg := config.LoadConfig()

	mb := rabbitmq.NewRabbitMq(cfg.RabbitMQURL)

	defer mb.Close()

	s := server.NewApiService(cfg.ApiPort)

	s.Use(func(api *server.ApiHandler) {

		api.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			lib.Json(w, "halo sekai")
		})

		api.Post(
			"/notification/email",
			func(w http.ResponseWriter, r *http.Request) {

				msg, err := io.ReadAll(r.Body)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					w.Header().Add("Content-Type", "application/json")
					lib.Json(w, err)
				}
				// msg := fmt.Sprintf("halo sekai %d", rand.New(rand.NewSource(99)))

				err = mb.Publish("email", []byte(msg))
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					w.Header().Add("Content-Type", "application/json")
					lib.Json(w, err)
				}

				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				lib.Json(w, "success")
			},
		)

	})

	msgs, err := mb.Consume("email")
	if err != nil {
		log.Fatalf("failed to start consumer\n %v", err)
	}

	go func() {
		for msg := range msgs {
			fmt.Printf("dapat: %v", string(msg.Body))
		}
	}()

	s.Run()
}
