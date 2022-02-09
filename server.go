package main

import (
	"os"

	"gitlab.com/pragmaticreviews/golang-mux-api/controller"
	router "gitlab.com/pragmaticreviews/golang-mux-api/http"
	"gitlab.com/pragmaticreviews/golang-mux-api/repository"
	"gitlab.com/pragmaticreviews/golang-mux-api/service"
)

var (
	postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.postService       = service.NewPostService(postRepository)
	postController controller.postController = controller.NewPostController(postService)
	httpRouter     router.router             = router.NewMuxRouter()
)

func main() {
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.GET("/posts/{id}", postController.GetPostByID)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(os.Getenv("PORT"))
}
