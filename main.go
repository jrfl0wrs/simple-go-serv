package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
    "sync"
)

type Post struct {
    ID   int    `json:"id"`
    Body string `json:"body"`
}

var (
    posts   = make(map[int]Post)
    nextID  = 1
    postsMu sync.Mutex
)

func main() {
    http.HandleFunc("/posts", postsHandler)
    http.HandleFunc("/posts/", postHandler)

    fmt.Println("Server is running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

http.HandleFunc("/posts", postsHandler)
http.HandleFunc("/posts/", postHandler)
log.Fatal(http.ListenAndServe(":8080", nil))

func postsHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        handleGetPosts(w, r)
    case "POST":
        handlePostPosts(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func postHandler(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Path[len("/posts/"):])
    if err != nil {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case "GET":
        handleGetPost(w, r, id)
    case "DELETE":
        handleDeletePost(w, r, id)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
id, err := strconv.Atoi(r.URL.Path[len("/posts/"):])
if err != nil {
    http.Error(w, "Invalid post ID", http.StatusBadRequest)
    return
}