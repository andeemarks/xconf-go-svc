package main

import (
	"log"
	"log/syslog"
	"net/http"

	"bitbucket.org/kardianos/osext"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
)

func main() {
	u := UserService{map[string]User{}}
	u.Register()

	homeFolder, _ := osext.ExecutableFolder()
	config := swagger.Config {
		WebServices:    restful.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: "http://localhost:8080",
		ApiPath:        "/apidocs.json",

		// Optionally, specifiy where the UI is located
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: homeFolder + "swagger-ui"}

	swagger.InstallSwaggerService(config)

	logwriter, e := syslog.New(syslog.LOG_NOTICE, "xconf-go-svc")
    if e == nil {
        log.SetOutput(logwriter)
    }

	log.Printf("start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
