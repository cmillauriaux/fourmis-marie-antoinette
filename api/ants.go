package ants

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"strconv"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
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

	// Admin routes
	r.HandleFunc("/api/blog/isAuthorized", isAuthorized)
	r.HandleFunc("/api/blog/articles", getArticlesList)
	r.HandleFunc("/api/blog/articles/add", addArticle).Methods("POST")
	r.HandleFunc("/api/blog/article/{articleID}", getArticle).Methods("GET")
	r.HandleFunc("/api/blog/article/{articleID}", addArticle).Methods("PUT")
	r.HandleFunc("/api/blog/article/{articleID}", deleteArticle).Methods("DELETE")

	http.Handle("/", r)
}

func getLastPicture(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	pic, err := DB.GetLastPicture(appengine.NewContext(r), 1)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	pic.Next = pic.DateTime + 60000
	fmt.Fprint(w, structToJSON(pic))
}

func getPreviousPicture(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
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
	w.Header().Add("Access-Control-Allow-Origin", "*")
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
	w.Header().Add("Access-Control-Allow-Origin", "*")
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

func getArticlesList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	message := checkAdminAuthorization(r)
	published := false

	// If user is an administrator, he can see unpublished articles
	if message.IsAdmin {
		published = true
	}

	// Get articles from DB
	articles, err := DB.GetAllArticles(appengine.NewContext(r), published)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	// Return response
	log.Println(structToJSON(articles))
	fmt.Fprint(w, structToJSON(articles))
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	article, err := DB.GetArticle(appengine.NewContext(r), params["articleID"])
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, structToJSON(article))
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	message := checkAdminAuthorization(r)
	if !message.IsAdmin || !message.IsLogin {
		fmt.Fprint(w, "Authorisation not found")
		return
	}

	err := DB.DeleteArticle(appengine.NewContext(r), params["articleID"])
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, "OK")
}

func addArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	message := checkAdminAuthorization(r)
	if !message.IsAdmin || !message.IsLogin {
		fmt.Fprint(w, "Authorisation not found")
		return
	}

	// Decode article from body request
	decoder := json.NewDecoder(r.Body)
	var article Article
	err := decoder.Decode(&article)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	defer r.Body.Close()

	_, err = DB.AddArticle(appengine.NewContext(r), &article)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, structToJSON(article))
}

func isAuthorized(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	message := checkAdminAuthorization(r)
	fmt.Fprint(w, structToJSON(message))
}

func checkAdminAuthorization(r *http.Request) Message {
	message := Message{NeedAdminAuthorization: true}
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)

	// If user is not logged, get signin URL
	if u == nil {
		url, _ := user.LoginURL(ctx, "/")
		message.SignInURL = url
		return message
	}

	// If user is logged
	message.IsLogin = true

	// Get Signout URL
	url, _ := user.LogoutURL(ctx, "/")
	message.SignOutURL = url

	if u.Admin {
		message.IsAdmin = true
	}

	return message
}
