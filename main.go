package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {

	Logger := zerolog.New(os.Stdout).With().Caller().Logger()

	router := gin.Default()

	router.POST("/answer", func(context *gin.Context) {

		playMessage(context, &Logger, messageStrings["greeting"])

	})

	router.Run(":443")
}
