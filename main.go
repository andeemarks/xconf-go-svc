package main

import (
	"net/http"
	"os"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("UserService.main")
var format = logging.MustStringFormatter("[%{module}] %{level} - %{message}")

func main() {
	u := UserService{map[string]User{}}
	u.Register()

	port := os.Getenv("PORT")
	if (port == "") {
		port = "8080"
	}

	config := swagger.Config{
		WebServices:    restful.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: "http://localhost:" + port,
		ApiPath:        "/apidocs.json"}

	swagger.InstallSwaggerService(config)

	logging.SetFormatter(format)
	logging.SetLevel(logging.INFO, "UserService.main")

	log.Notice("start listening on localhost:" + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
