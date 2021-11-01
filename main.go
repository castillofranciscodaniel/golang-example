package main

import (
	"bitbucket.org/chattigodev/chattigo-golang-library/spring-cloud-config"
	"github.com/castillofranciscodaniel/golang-example/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)

var routes = chi.NewRouter()

func init() {
	spring_cloud_config.LoadConfiguration(os.Getenv("SPRING_CLOUD_CONFIG_URI"), os.Getenv("APP_NAME"), os.Getenv("SPRING_PROFILES_ACTIVE"))
	container := InitializeServer()
	route(container)
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

	routes.Get("/health", container.HealthHandler.Health)
	routes.Get("/metrics", promhttp.Handler().ServeHTTP)
	routes.Post("/test", container.ProductHandler.TestErrorDto)

	routes.Post("/modifyProductById", container.ProductHandler.HandlerProductByID)

}
