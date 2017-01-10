package ants

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"strconv"

	"google.golang.org/appengine"
)

var (
	DB *Persistance
)

func init() {
	// Connect to DB
	DB = &Persistance{}

	r := mux.NewRouter()
	r.HandleFunc("/api/pictures/last", getLastPicture)
	r.HandleFunc("/api/pictures/previous/{DateTime}", getPreviousPicture)
	r.HandleFunc("/api/pictures/next/{DateTime}", getNextPicture)
	r.HandleFunc("/api/gif/last", getLastGIF)
	//r.HandleFunc("/api/test/generate", makeTest)

	http.Handle("/", r)
}

func getLastPicture(w http.ResponseWriter, r *http.Request) {
	pic, err := DB.GetLastPicture(appengine.NewContext(r), 1)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	pic.Next = pic.DateTime + 60000
	fmt.Fprint(w, structToJSON(pic))
}

func getPreviousPicture(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dateTime, err := strconv.ParseInt(vars["DateTime"], 10, 64)
	if err != nil {
		fmt.Fprint(w, "Bad timestamp format")
	}
	pic, err := DB.GetPreviousPicture(appengine.NewContext(r), 1, dateTime)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	fmt.Fprint(w, structToJSON(pic))
}

func getNextPicture(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dateTime, err := strconv.ParseInt(vars["DateTime"], 10, 64)
	if err != nil {
		fmt.Fprint(w, "Bad timestamp format")
	}
	pic, err := DB.GetNextPicture(appengine.NewContext(r), 1, dateTime)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	fmt.Fprint(w, structToJSON(pic))
}

func getLastGIF(w http.ResponseWriter, r *http.Request) {
	pics, _ := DB.GetAllPicture(appengine.NewContext(r), 1)
	files := make([]string, len(pics))
	for index, pic := range pics {
		files[index] = pic.Link
	}
	gif := makeGif(files, appengine.NewContext(r))
	w.Header().Set("Content-Type", "image/gif")
	w.Write(gif)
}

func makeTest(w http.ResponseWriter, r *http.Request) {
	DB.PutDataTest(appengine.NewContext(r))

	fmt.Fprint(w, "OK")
}

func structToJSON(structure interface{}) string {
	b, err := json.Marshal(structure)
	if err != nil {
		return ""
	}
	return string(b)
}
