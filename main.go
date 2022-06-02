package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-module/components"
	fakedb "go-module/db/fake"
	gintransport "go-module/modules/post/transport"
	"os"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	db, err := fakedb.OpenDB(os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}

	runGinService(db)
}

func runGinService(db *fakedb.FakeDB) error {

	router := gin.Default()

	ctx := components.NewAppContext(db)

	//router.Use(middleware.Recover())

	router.GET("/api/v1/post/find/:id", gintransport.GetPostById(ctx))

	return router.Run(":" + os.Getenv("SERVER_PORT"))
}
