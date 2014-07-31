package main

import (
	"log"
	"log/syslog"
	"net/http"
	"os"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
)


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

	logwriter, e := syslog.New(syslog.LOG_NOTICE, "xconf-go-svc")
	if e == nil {
		log.SetOutput(logwriter)
	}

	log.Printf("start listening on localhost:" + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
