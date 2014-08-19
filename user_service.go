package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

type UserService struct {
	Users map[string]User
}

type User struct {
	Id, Name string
}

func (u UserService) Register() {
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
}

func logRequests(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	log.Printf("%s %s", request.Request.Method, request.Request.URL)
	chain.ProcessFilter(request, response)
}

// GET http://localhost:8080/users/1
func (u UserService) FindUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	log.Printf("looking for user with id %s", id)
	usr := u.Users[id]
	if len(usr.Id) == 0 {
		response.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		response.WriteEntity(usr)
	}
}

// PUT http://localhost:8080/users/1
// <User><Id>1</Id><Name>Melissa Raspberry</Name></User>
func (u *UserService) UpdateUser(request *restful.Request, response *restful.Response) {
	usr := new(User)
	log.Printf("updating user: %s", usr)
	err := request.ReadEntity(&usr)
	if err == nil {
		u.Users[usr.Id] = *usr
		response.WriteEntity(usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

// PUT http://localhost:8080/users
// <User><Id>1</Id><Name>Melissa</Name></User>
func (u *UserService) CreateUser(request *restful.Request, response *restful.Response) {
	u.createUser(request.PathParameter("user-id"), request, response)
}

func (u *UserService) createUser(userId string, request *restful.Request, response *restful.Response) {
	usr := User{Id: userId}
	// log.Printf("%s", request.Request.Body)
	err := request.ReadEntity(&usr)
	// log.Printf("%s %s %s %s %s", err, usr, u.Users, usr.Id, u.Users[usr.Id])
	if err == nil {
		u.Users[usr.Id] = usr
		response.WriteHeader(http.StatusCreated)
		response.WriteEntity(usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

// DELETE http://localhost:8080/users/1
func (u *UserService) RemoveUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	log.Printf("removing user with id %s", id)
	delete(u.Users, id)
}
