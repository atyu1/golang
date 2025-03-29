package main

import (
	"net/http"
	"time"

	"github.com/atyu1/permetezunk/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	//r.Static("./static")

	r.LoadHTMLGlob("./templates/*")

	models.ConnectDatabase()

	r.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	server := &http.Server{
		Addr:           ":8888",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
