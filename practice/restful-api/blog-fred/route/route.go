// We define all the routes for the CRUD operations
package route

import (
	"project01/practice/restful-api/blog-fred/handler"
	"project01/practice/restful-api/blog-fred/middleware"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// create a post
	r.HandleFunc("/create", handler.CreatePost).Methods("POST")
	// List the titles of all posts
	r.HandleFunc("/list", handler.ListPosts).Methods("GET")
	// Get a post by the post title
	r.HandleFunc("/post", handler.GetPostByTitle).Methods("GET")
	r.HandleFunc("/edit", handler.UpdatePost).Methods("PUT")
	// delete a blog post
	r.HandleFunc("/delete", middleware.RequireAuth(handler.DeletePost)).Methods("DELETE")
	return r
}
