package server

import (
	"fmt"
	"net/http"
	"urlShortener/internal/link"
	"urlShortener/internal/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Start(port int) {
	engine := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	engine.Use(cors.New(config))

	engine.GET("/", serveFrontend)
	engine.GET("/link", showLinks)
	engine.POST("/link", createLink)

	engine.Static("/static", "./frontend/static")

	engine.Use(customMiddleware)

	engine.Run(fmt.Sprintf("localhost:%d", port))
}

func serveFrontend(c *gin.Context) {
	c.File("./frontend/index.html")
}

func customMiddleware(c *gin.Context) {

	shortPath := c.Request.URL.Path[1:]
	fullUrl, err := storage.GetFullUrl(shortPath)
	if err != nil {
		c.File("./frontend/404.html")
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, fullUrl)

	c.Next()
}

func checkIfReachable(url string) bool {
	_, error := http.Get(url)
	return error == nil
}

func createLink(c *gin.Context) {
	var newLink link.Link

	if error := c.ShouldBindWith(&newLink, binding.JSON); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad JSON format",
		})
		return
	}

	if checkIfReachable(newLink.Url) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Url is not reachable",
		})
		return
	}

	err := storage.Put(&newLink)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, newLink)
}

func showLinks(c *gin.Context) {
	links, err := storage.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on database selection",
		})
		return
	}
	c.JSON(http.StatusOK, links)
}
