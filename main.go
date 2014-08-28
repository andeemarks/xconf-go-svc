package main

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/op/go-logging"
	"net/http"
	"os"
	// "io"
	"os/signal"
	"syscall"
)

var log = logging.MustGetLogger("UserService.main")

func getPort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
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

func configureExitHandler(port string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		log.Notice("stop listening on localhost:" + port)
		os.Exit(1)
	}()
}

func configureLogging() {
	stdErrorLogger := logging.NewLogBackend(os.Stderr, "", 3)
	logFile, err := os.OpenFile("xconf-go-svc.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	fileLogger := logging.NewLogBackend(logFile, "", 3)

	logging.SetBackend(fileLogger, stdErrorLogger)
	logging.SetFormatter(logging.MustStringFormatter("%{color}[%{module}] %{level}%{color:reset} - %{message}"))
	logging.SetLevel(logging.DEBUG, "UserService.main")
}

func main() {
	configureLogging()

	port := getPort()

	configureExitHandler(port)

	u := UserService{map[string]User{}}
	u.Register()

	configureSwagger(port)

	log.Notice("start listening on localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
