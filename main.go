package main
import (
      "code.google.com/p/gorest"
      "github.com/jagregory/halgo"
      "database/sql"
      "log" 
      "net/http"
      "strconv"
)

func main() {
    // GoREST usage: http://localhost:8181/tutorial/hello
    gorest.RegisterService(new(Tutorial)) //Register our service
    http.Handle("/",gorest.Handle())    
    http.ListenAndServe(":8181",nil)
}

//Service Definition
type Tutorial struct {
    halgo.Links
    gorest.RestService `root:"/tutorial/" consumes:"application/json" produces:"application/json"`
    hello gorest.EndPoint `method:"GET" path:"/hello/" output:"string"`
    insert gorest.EndPoint `method:"POST" path:"/insert/" postdata:"int"`
}

func(serv Tutorial) Hello() string {
    return "Hello World"
}

func(serv Tutorial) Insert(number int) {
    db, err := sql.Open("mysql", "root:password@/dbname?charset=utf8")
    if err != nil {
      db.Close()      
      log.Fatal(err)
      serv.ResponseBuilder().SetResponseCode(500)
    } else {
      db.Exec("INSERT INTO table (number) VALUES(" + strconv.Itoa(number) + ");") 
      db.Close()      
      serv.ResponseBuilder().SetResponseCode(200)
    }
}