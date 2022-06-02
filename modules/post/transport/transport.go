package gintransport

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-module/components"
	postrepo "go-module/modules/post/repo"
	postService "go-module/modules/post/service"
	"net/http"
	"os"
)

func GetPostById(ctx *components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		repo := postrepo.NewPostRepo(ctx.FakeDB)
		service := postService.NewPostService(repo)
		post, _ := service.GetPostById(id)

		if os.Getenv("PANIC_ERROR_TRANSPORT") == "true" {
			panic(errors.New("there are something wrong with transport layer"))
		}

		c.JSON(http.StatusOK, post)
	}
}
