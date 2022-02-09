package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	// "gitlab.com/pragmaticreview/golang-mux-api/entity"
	"gitlab.com/pragmaticreview/golang-mux-api/errors"
	"gitlab.com/pragmaticreview/golang-mux-api/service"
)

type controller struct{}

var (
	postService service.PostService
)

type PostController interface {
	GetPostByID(response http.ResponseWriter, request *http.Request)
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError{Message: "Error getting the posts"})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*controller) GetPostsById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	postID := strings.Split(request.URL.Path, "/")[2]
	posts, err := postService.FindByID(postID)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Non posts found!"})
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(posts)
	}
}
