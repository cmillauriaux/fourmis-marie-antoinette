package ants

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

/*func main() {
	persistance := &persistance.Persistance{}
	err := persistance.Init("prototype-149014")

	if err != nil {
		log.Fatal(err)
	}

	pic, err := persistance.GetLastPicture(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pic)
}*/

var (
	DB *Persistance
)

func init() {
	// Connect to DB
	DB = &Persistance{}

	http.HandleFunc("/", handler)
	http.HandleFunc("/makeTest", makeTest)
}

func handler(w http.ResponseWriter, r *http.Request) {
	pic, _ := DB.GetLastPicture(appengine.NewContext(r), 1)
	fmt.Fprint(w, pic)
}

func makeTest(w http.ResponseWriter, r *http.Request) {
	DB.PutDataTest(appengine.NewContext(r))
	fmt.Fprint(w, "OK")
}
