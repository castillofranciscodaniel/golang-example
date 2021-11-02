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

var Routes = chi.NewRouter()

func init() {

	os.Setenv("SPRING_CLOUD_CONFIG_URI", "https://dev-gke.chattigo.com/leones/config")
	os.Setenv("SPRING_PROFILES_ACTIVE", "development")
	os.Setenv("APP_NAME", "bff-chattigo-webchat")
	spring_cloud_config.LoadConfiguration(os.Getenv("SPRING_CLOUD_CONFIG_URI"), os.Getenv("APP_NAME"), os.Getenv("SPRING_PROFILES_ACTIVE"))
	container := InitializeServer()
	Route(container)
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", Routes))
}

func Route(container config.ContainerServiceImp) *chi.Mux {
	Routes.Use(middleware.AllowContentType("application/json"))
	Routes.Use(middleware.RequestID)
	Routes.Use(middleware.RealIP)
	Routes.Use(middleware.Logger)
	Routes.Use(middleware.Recoverer)

	Routes.Get("/health", container.HealthHandler.Health)
	Routes.Get("/metrics", promhttp.Handler().ServeHTTP)
	Routes.Post("/modifyProductById/pointer", container.ProductHandler.HandlerProductByIDPointer)

	Routes.Post("/modifyProductById", container.ProductHandler.HandlerProductByID)

	return Routes

}
