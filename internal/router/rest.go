package router

import (
	"fmt"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
	"order-service/internal/config"
	"order-service/internal/rest"
)

const (
	healthRestPath = "/health"
	orderRestPath  = "/order"
)

func RunRest() {
	restConfig := config.GetConfig().Rest
	basePath := fmt.Sprintf(`/%v%v`, restConfig.APIVersion, restConfig.PathPrefix)
	router := chi.NewRouter()
	router.Route(basePath, func(r chi.Router) {
		r.Get(healthRestPath, rest.HealthHandler())
		r.Post(orderRestPath, rest.CreateOrderEntry)
	})
	log.Info("order service listening at http://localhost:" + restConfig.Port + basePath)
	log.Fatal(http.ListenAndServe(":"+restConfig.Port, router))
}
