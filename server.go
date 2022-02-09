package main

import (
	"orign/cache"
	"os"

	"gitlab.com/pragmaticreviews/golang-mux-api/cache"
	"gitlab.com/pragmaticreviews/golang-mux-api/controller"
	router "gitlab.com/pragmaticreviews/golang-mux-api/http"
	"gitlab.com/pragmaticreviews/golang-mux-api/repository"
	"gitlab.com/pragmaticreviews/golang-mux-api/service"
)

var (
	postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.postService       = service.NewPostService(postRepository)
	postCache      cache.PostCache           = cache.NewRedisCache("localhost:6379", 1, 10)
	postController controller.postController = controller.NewPostController(postService, postCache)
	httpRouter     router.router             = router.NewMuxRouter()
)

func main() {
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.GET("/posts/{id}", postController.GetPostByID)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(os.Getenv("PORT"))
}
