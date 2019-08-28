package main

import (
	"log"
	"net"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

type UserData struct {
	user_id int
	name    string
}

func main() {

	users := makeUsersData()

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/lookup/#host", func(w rest.ResponseWriter, req *rest.Request) {
			ip, err := net.LookupIP(req.PathParam("host"))
			if err != nil {
				rest.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteJson(&ip)
		}),
		rest.Get("/seiya/#id", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(users)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8383", api.MakeHandler()))
}

func makeUsersData() map[int]UserData {
	users := map[int]UserData{
		0: UserData{1, "AA"},
		1: UserData{2, "BB"},
	}
	return users
}
