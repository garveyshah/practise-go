package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"project01/practice/restful-api/blog-fred/model"
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
func ListPosts(w http.ResponseWriter, r *http.Request) {
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

// Get Post by Title
func GetPostByTitle(w http.ResponseWriter, r *http.Request) {
	// retrieve title of post
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	// checck title is present
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

// Update Post(PUT)
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// get title
	// check title is provided
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	// check if such post exists
	post, ok := allPosts[title]
	if !ok {
		http.Error(w, "BlogPost not found", http.StatusNotFound)
		return
	}

	var updatedPost model.Post

	// read request body
	if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// update post
	allPosts[title] = updatedPost
	post = updatedPost

	// return ok status
	w.WriteHeader(http.StatusOK)

	// return updated content
	json.NewEncoder(w).Encode(post)
}

// Delete Post(Delete)
func DeletePost(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")

	// retrieve post
	_, ok := allPosts[title]
	if !ok {
		http.Error(w, "No post with such title", http.StatusNotFound)
	}

	// deletes the post from the map.
	delete(allPosts, title)

	w.WriteHeader(http.StatusOK)
}
