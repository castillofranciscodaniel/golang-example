package main

import (
	"bitbucket.org/chattigodev/chattigo-golang-library/spring-cloud-config"
	"encoding/json"
	"fmt"
	"github.com/castillofranciscodaniel/golang-example/config"
	"github.com/castillofranciscodaniel/golang-example/pkg/dto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/reactivex/rxgo/v2"
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
	test()
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
	Routes.Mount("/debug", middleware.Profiler())

	Routes.Get("/health", container.HealthHandler.Health)
	Routes.Get("/metrics", promhttp.Handler().ServeHTTP)
	Routes.Post("/message", container.MessageChannelHandler.SaveMessageChannel)

	Routes.Get("/message", container.MessageChannelHandler.GetMessageChannel)

	return Routes

}

func test() {
	//messageMarshall()
	//messageRx()
	nativeError().FlatMap(func(item rxgo.Item) rxgo.Observable {
		return rxgo.Empty()
	}).Observe()
}

func messageMarshall() {
	messageChannel := dto.MessageChannel{}
	bytes := []byte(`{"id":2,"name":"Agua","price":50}`)
	err := json.Unmarshal(bytes, &messageChannel)
	if err != nil {

		return
	}

}

func messageRx() {
	messageChannel := dto.MessageChannel{}
	rxgo.Just(messageChannel)().Observe()
}

type NativeError struct {
	StatusCode  int                    `json:"statusCode,omitempty"`
	Msg         string                 `json:"msg,omitempty"`
	TraceDetail map[string]interface{} `json:"traceDetail,omitempty"`
}

// Error - implemented from error interface
func (e *NativeError) Error() string {
	return fmt.Sprintf("Msg: %v. StatusCode: %v. TraceDetail: %v", e.Msg, e.StatusCode, e.TraceDetail)
}

func nativeError() rxgo.Observable {
	err := NativeError{
		StatusCode: 404,
		Msg:        "Status Not Found",
	}
	return rxgo.Just(&err)().FlatMap(func(item rxgo.Item) rxgo.Observable {
		if item.Error() {
			return rxgo.Just(item.E)()
		}

		return rxgo.Empty()
	})
}
