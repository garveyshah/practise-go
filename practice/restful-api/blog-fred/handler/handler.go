package handler

import (
	"api1/model"
	"encoding/json"
	"fmt"
	"net/http"
)

// an in-built memory database to store and retrieve our posts.
var allPosts = make(map[string]model.Post)

// Create Post(POST)
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	// read content from body into a new decoder
	decoder := json.NewDecoder(r.Body)
	// decode content into our Post struct
	err := decoder.Decode(&post)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusInternalServerError)
		return
	}

	// check for unique post title
	_, ok := allPosts[post.Title]
	if ok {
		http.Error(w, "Post title already exists", http.StatusBadRequest)
		return
	}

	// append post to our memory
	allPosts[post.Title] = post

	// prints out structs with field names
	fmt.Fprintf(w, "%+v", post)
}

// Get Post (GET)
func ListPost(w http.ResponseWriter, r *http.Request) {
	titles := []string{}
	for _, post := range allPosts {
		titles = append(titles, post.Title)
	}

	if len(titles) == 0 {
		http.Error(w, "no posts found", http.StatusNotFound)
}

	json.NewEncoder(w).Encode(titles)
	// uncomment to print out structs with field names
	// fmt.Fprintf(w, "%+v", titles)
}

// Update Post(PUT)
func GetPostByTitle(w, http.ResponseWriter, r *http.Request) {
	// retrieve title of post
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	//checck title is present
	post, ok := allPosts[title]
	if !ok {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	// encode the contents
	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Post not found", http.StatusInternalServerError)
		return
	}
	
}
