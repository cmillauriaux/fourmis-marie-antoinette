package ants

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

var (
	DB *Persistance
)

func init() {
	// Connect to DB
	DB = &Persistance{}

	http.HandleFunc("/api/pictures/last", getLastPicture)
	http.HandleFunc("/api/pictures/previous", getLastPicture)
	http.HandleFunc("/api/pictures/next", getLastPicture)
	http.HandleFunc("/api/pictures/from", getLastPicture)
	http.HandleFunc("/api/gif/last", getLastGIF)
	http.HandleFunc("/api/test/generate", makeTest)
}

func getLastPicture(w http.ResponseWriter, r *http.Request) {
	pic, _ := DB.GetLastPicture(appengine.NewContext(r), 1)
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
	//fmt.Fprint(w, gif)
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
