package main

import(
	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	
	router.GET("/", func(c *gin.Context) {

		c.HTML(200, "index.tmpl", gin.H{})
	})

	router.Run(":5000")
}
