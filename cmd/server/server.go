package server

import (
	"fmt"
	"log"
	"net/http"
	"tfg/medical-center-service/cmd/config"
	routeHttp "tfg/medical-center-service/pkg/ports/http"
)

func StartServer(cfg *config.Config, deps *config.Dependencies) error {
	router := routeHttp.NewRouter(deps)
	log.Println("Servidor escuchando en el puerto", cfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), router)
}
