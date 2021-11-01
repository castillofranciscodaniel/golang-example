package main

import (
	"bitbucket.org/chattigodev/chattigo-golang-library/spring-cloud-config"
	"github.com/castillofranciscodaniel/golang-example/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)

var routes = chi.NewRouter()

func init() {
	spring_cloud_config.LoadConfiguration(os.Getenv("SPRING_CLOUD_CONFIG_URI"), os.Getenv("APP_NAME"), os.Getenv("SPRING_PROFILES_ACTIVE"))
	//container := InitializeServer()
	//route(container)
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", routes))
}

func route(container config.ContainerServiceImp) {
	routes.Use(middleware.AllowContentType("application/json"))
	routes.Use(middleware.RequestID)
	routes.Use(middleware.RealIP)
	routes.Use(middleware.Logger)
	routes.Use(middleware.Recoverer)

	//TODO: notas antes de pasar a qa
	// TokenAuth --> sacar el doble true

	routes.Get("/health", container.HealthHandler.Health)
	routes.Get("/metrics", promhttp.Handler().ServeHTTP)


	routes.Post("/modifyProductById", container.ProductHandler.HandlerProductByID)

}

