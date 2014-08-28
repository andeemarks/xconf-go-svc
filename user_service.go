package main

import (
	"bytes"
	"github.com/emicklei/go-restful"
	"github.com/rcrowley/go-metrics"
	"net/http"
)

type UserService struct {
	Users map[string]User
}

type User struct {
	Id, Name string
}

var addRequests = metrics.NewCounter()
var updateRequests = metrics.NewCounter()
var deleteRequests = metrics.NewCounter()
var defaultRegistry metrics.Registry = metrics.NewRegistry()

func initMetrics() {
	metrics.Register("number-of-users-added", addRequests)
	metrics.Register("number-of-users-updated", updateRequests)
	metrics.Register("number-of-users-deleted", deleteRequests)
	// metrics.Log(defaultRegistry, 60e9, log.New(os.Stdout, "metrics: ", log.Lmicroseconds))
}

func (u UserService) setupRoutes(ws *restful.WebService) {
	ws.Route(ws.GET("/status").
		To(u.status).
		Doc("show service stats").
		Operation("serviceStats"))

	ws.Route(ws.GET("/{user-id}").
		To(u.FindUser).
		Doc("get a user").
		Operation("findUser").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Writes(User{})) // on the response

	ws.Route(ws.PUT("/{user-id}").
		To(u.UpdateUser).
		Doc("update a user").
		Operation("updateUser").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Reads(User{})) // from the request

	ws.Route(ws.PUT("").
		To(u.CreateUser).
		Doc("create a user").
		Operation("createUser").
		Reads(User{})) // from the request

	ws.Route(ws.DELETE("/{user-id}").
		To(u.RemoveUser).
		Doc("delete a user").
		Operation("removeUser").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")))
}

func (u UserService) Register() {
	log.Notice("Service registration started")
	restful.SetCacheReadEntity(false)
	restful.Filter(logReceivedRequests)

	initMetrics()

	ws := new(restful.WebService)
	ws.Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	u.setupRoutes(ws)

	ws.Filter(logSupportedRoutes)

	restful.Add(ws)
	cors := restful.CrossOriginResourceSharing{ExposeHeaders: []string{"X-My-Header"}, AllowedHeaders: []string{"content-type"}, CookiesAllowed: true, Container: restful.DefaultContainer}
	restful.Filter(cors.Filter)

	log.Notice("Service registration finished")
}

func logReceivedRequests(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	log.Info("Request received: %s %s", request.Request.Method, request.Request.URL)
	chain.ProcessFilter(request, response)
	log.Info("Response status: %d", response.StatusCode())
}

func logSupportedRoutes(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	log.Info("Route supported: %s %s", request.Request.Method, request.Request.URL)
	chain.ProcessFilter(request, response)
}

func (u UserService) status(request *restful.Request, response *restful.Response) {
	var b bytes.Buffer
	metrics.WriteJSONOnce(metrics.DefaultRegistry, &b)
	response.WriteHeader(http.StatusOK)
	response.WriteEntity(b.String())
}

// GET http://localhost:8080/users/1
func (u UserService) FindUser(request *restful.Request, response *restful.Response) {
	u.findUser(request.PathParameter("user-id"), response)
}

func (u *UserService) findUser(userId string, response *restful.Response) {
	log.Info("Looking for user with id: %s", userId)
	usr := u.Users[userId]
	if len(usr.Id) == 0 {
		log.Info("User with id: %s not found", userId)
		response.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		response.WriteEntity(usr)
	}
}

// PUT http://localhost:8080/users/1
func (u *UserService) UpdateUser(request *restful.Request, response *restful.Response) {
	usr := new(User)
	err := request.ReadEntity(&usr)
	if err == nil {
		u.Users[usr.Id] = *usr
		log.Info("Updating user with id: %s", usr.Id)
		updateRequests.Inc(1)
		response.WriteEntity(usr)
	} else {
		log.Info("Error updating user with id: %s - %s", usr.Id, err)
		response.WriteError(http.StatusInternalServerError, err)
	}
}

// PUT http://localhost:8080/users
func (u *UserService) CreateUser(request *restful.Request, response *restful.Response) {
	u.createUser(request.PathParameter("user-id"), request, response)
}

func (u *UserService) createUser(userId string, request *restful.Request, response *restful.Response) {
	usr := User{Id: userId}
	err := request.ReadEntity(&usr)
	log.Info("Creating user: %s", usr)
	if err == nil {
		u.Users[usr.Id] = usr
		addRequests.Inc(1)
		response.WriteHeader(http.StatusCreated)
		response.WriteEntity(usr)
	} else {
		log.Info("Error creating user with id: %s - %s", userId, err)
		response.WriteError(http.StatusInternalServerError, err)
	}
}

// DELETE http://localhost:8080/users/1
func (u *UserService) RemoveUser(request *restful.Request, response *restful.Response) {
	u.removeUser(request.PathParameter("user-id"), request, response)
}

func (u *UserService) removeUser(userId string, request *restful.Request, response *restful.Response) {
	log.Info("Removing user with id: %s", userId)
	deleteRequests.Inc(1)
	delete(u.Users, userId)
}
