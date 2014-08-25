package main

import (
	"github.com/emicklei/go-restful"
	// "github.com/op/go-logging"
	"net/http"
)

type UserService struct {
	Users map[string]User
}

type User struct {
	Id, Name string
}

func (u UserService) Register() {
	log.Notice("Service registration started")
	restful.Filter(logRequests)
	restful.SetCacheReadEntity(false)

	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/{user-id}").To(u.FindUser).
		Doc("get a user").
		Operation("findUser").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Writes(User{})) // on the response

	ws.Route(ws.PUT("/{user-id}").To(u.UpdateUser).
		Doc("update a user").
		Operation("updateUser").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Reads(User{})) // from the request

	ws.Route(ws.PUT("").To(u.CreateUser).
		Doc("create a user").
		Operation("createUser").
		Reads(User{})) // from the request

	ws.Route(ws.DELETE("/{user-id}").To(u.RemoveUser).
		Doc("delete a user").
		Operation("removeUser").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")))

	restful.Add(ws)
	log.Notice("Service registration finished")
}

func logRequests(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	log.Info("Request received: %s %s", request.Request.Method, request.Request.URL)
	chain.ProcessFilter(request, response)
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
	log.Info("Creating user with id: %s", userId)
	usr := User{Id: userId}
	err := request.ReadEntity(&usr)
	if err == nil {
		u.Users[usr.Id] = usr
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
	delete(u.Users, userId)
}
