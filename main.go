package main

import (
	"log"
	"net"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

type Result struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total    int     `json:"total"`
		MaxScore float32 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			Id     string  `json:"_id"`
			Score  float32 `json:"_score"`
			Source struct {
				User    string `json:"user"`
				Message string `json:"message"`
				Date    string `json:"date"`
				Likes   int    `json:"likes"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
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

func makeUsersData() Result {
	var result Result

	result.Took = 12

	result.TimedOut = true

	result.Shards.Total = 5
	result.Shards.Successful = 1
	result.Shards.Failed = 0

	result.Hits.Total = 5
	result.Hits.MaxScore = 300

	return result
}
