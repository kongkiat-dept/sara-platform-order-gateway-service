package protocol

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sara-platform-order-gateway-service/configs"
	"sara-platform-order-gateway-service/internal/core/services"
	"sara-platform-order-gateway-service/internal/handlers"
	repository "sara-platform-order-gateway-service/internal/repositories"
	"sara-platform-order-gateway-service/pkg/apprequest"

	"github.com/gofiber/fiber/v2"
)

type config struct {
	Env string
}

func ServeHTTP() error {
	app := fiber.New()
	var cfg config
	flag.StringVar(&cfg.Env, "env", "", "the environment to use")
	flag.Parse()
	configs.InitViper("./configs", cfg.Env)
	// Graceful shutdown ...
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("Gracefull shut down ...")
			app.Shutdown()
		}
	}()

	var orderHandler *handlers.HTTPHandler
	{
		portOrderRepo := repository.NewPlatformOrderHTTP(
			configs.GetViper().Platform.Order.Services.Http,
			apprequest.NewRequester(),
		)
		svc := services.New(portOrderRepo)
		orderHandler = handlers.NewHTTPHandler(svc, configs.GetViper().Order.Services.Http)
	}

	app.Get("/healthz", orderHandler.HealthCheck)
	orderSearchApi := app.Group("/v1/api")
	{
		orderSearchApi.Get("/search/customer", orderHandler.SearchOrderByCustomer)
		orderSearchApi.Get("/search/customer/:id", orderHandler.SearchOrderByCustomer)

		orderSearchApi.Get("/search/vendor", orderHandler.SearchOrderByVendor)
		orderSearchApi.Get("/search/vendor/:id", orderHandler.SearchOrderByVendor)

		orderSearchApi.Put("/update/status", orderHandler.UpdateStatus)
	}

	err := app.Listen(":" + configs.GetViper().App.Port)
	if err != nil {
		return err
	}

	fmt.Println("Listerning on port: ", configs.GetViper().App.Port)
	return nil
}
