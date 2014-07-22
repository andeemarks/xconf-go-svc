package xconf_go_svc

import (
	"log"
	"net/http"

	"bitbucket.org/kardianos/osext"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
)

// This example is functionally the same as the example in restful-user-resource.go
// with the only difference that is served using the restful.DefaultContainer


func main() {
	u := UserService{map[string]User{}}
	u.Register()

	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs and enter http://localhost:8080/apidocs.json in the api input field.
	homeFolder, _ := osext.ExecutableFolder()
	config := swagger.Config {
		WebServices:    restful.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: "http://localhost:8080",
		ApiPath:        "/apidocs.json",

		// Optionally, specifiy where the UI is located
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: homeFolder + "swagger-ui"}

	swagger.InstallSwaggerService(config)

	log.Printf("start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
