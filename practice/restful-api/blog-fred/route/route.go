// We define all the routes for the CRUD operations
package route

import (
	"github.com/gorilla/mux"

	"api1/handler"
	"api1/middleware"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// create a post
	r.Handler("/create", handler.CreatePost).Methods("POST")
	// List the titles of all posts
	r.HandlerFunc("/list", handler.ListPosts).Methods("GET")
	// Get a post by the post title
	r.HandlerFunc("/post", handler.GetPostByTitle).Methods("GET")
	r.HandlerFunc("/edit", handler.UpdatePost).Methods("PUT")
	// delete a blog post
	r.HandlerFunc("/delete", middleware.RequireAuth(handler.DeletePost)).Methods("DELETE")
	return r
}