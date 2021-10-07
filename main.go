package main

import (
	_ "embed"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/on_publish", VerifyPublish)
	router.POST("/streamers", POSTStreamer)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func VerifyPublish(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		println("Can't parse request form on publish: " + fmt.Sprintf("%v", err))
		return
	}
	tcurl := request.Form.Get("tcurl")
	parts := strings.Split(tcurl, "?key=")
	if len(parts) != 2 {
		writer.WriteHeader(http.StatusForbidden)
		return
	}
	streamer, err := getStreamerByKey(parts[1])
	if err != nil || streamer.Key != parts[1] {
		writer.WriteHeader(http.StatusForbidden)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func POSTStreamer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	name := r.PostForm.Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key := randomString(15)
	err = addStreamer(Streamer{Name: name, Key: key})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

//go:embed index.gohtml
var index string

func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	parsed, err := template.New("index").Parse(index)
	if err != nil {
		return
	}

	streamers, err := getAllStreamers()
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = parsed.Execute(w, IndexD{streamers})
	if err != nil {
		log.Printf("%v", err)
	}
}

type IndexD struct {
	Streamers []Streamer
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
