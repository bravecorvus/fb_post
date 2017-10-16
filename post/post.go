package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	// "net"
	// "path/filepath"

	fb "github.com/huandu/facebook"
)


var SharedVolumePath string

type request_struct struct {
	Id          string `json:"id"`
	Type          string `json:"type"`
	AccessToken string `json:"access_token"`
}

func main() {
	fb.Version = "v2.10"
	// os.Getenv("")
	// hostPort := net.JoinHostPort("0.0.0.0", os.Getenv("PORT"))
	// SharedVolumePath, _ = filepath.Abs(os.Getenv("MOUNTPATH"))
	SharedVolumePath = "/exports/assets"
	http.HandleFunc("/", Post)
	// http.ListenAndServe(hostPort, nil)
	http.ListenAndServe(":8080", nil)
}

func Post(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var query request_struct
	err := decoder.Decode(&query)
	if err != nil {
		panic(err)
	}
	var posttype, postextension string
	if query.Type == "image" {
		posttype = "photo"	
		postextension = ".jpg"
	} else if query.Type == "video" {
		posttype = "video"
		postextension = ".mp4"
	}
	fmt.Println(SharedVolumePath + "/" + query.Id + postextension)
	uploadfile, err := os.Open(SharedVolumePath + "/" + query.Id + postextension)
	res, e := fb.Post("/924875817589943/videos", fb.Params{
		"type": posttype,
        "name": "",
        "caption": "",
        // "picture": fb.File(SharedVolumePath + "/" + query.Id + postextension),
		// "picture": fb.FileAliasWithContentType(SharedVolumePath, query.Id + postextension),
		// "link": fb.File(SharedVolumePath + "/" + query.Id + postextension),
		"video": fb.Data(query.Id + postextension, uploadfile),
		// "link": fb.Data(query.Id + postextension, uploadfile),
		"link": "",
		"description":"",
        "access_token": query.AccessToken,
		"published": "false",
    })
    fmt.Println(e)
    fmt.Println(res)
}

