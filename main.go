package main

import (
	"net/http"
	"os"
	"os/signal"
    "syscall"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("UserService.main")
var format = logging.MustStringFormatter("[%{module}] %{level} - %{message}")

func getPort() (port string) {
	port = os.Getenv("PORT")
	if (port == "") {
		port = "8080"
	}

	return port
}

func configureSwagger(port string) {
	config := swagger.Config{
		WebServices:    restful.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: "http://localhost:" + port,
		ApiPath:        "/apidocs.json"}

	swagger.InstallSwaggerService(config)
}

func handleExit(port string) {
	c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    signal.Notify(c, syscall.SIGTERM)
    go func() {
        <-c
		log.Notice("stop listening on localhost:" + port)
        os.Exit(1)
    }()
}

func main() {
	port := getPort()

	handleExit(port)

	u := UserService{map[string]User{}}
	u.Register()

	configureSwagger(port)

	logging.SetFormatter(format)
	logging.SetLevel(logging.INFO, "UserService.main")

	log.Notice("start listening on localhost:" + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
