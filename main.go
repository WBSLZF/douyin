package main

import (
	message "github.com/RaymondCode/simple-demo/messsage"
	"github.com/gin-gonic/gin"
)

func main() {

	go message.RunMessageServer()

	r := gin.Default()
	initRouter(r)
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
