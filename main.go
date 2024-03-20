package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

var urlMap map[string]string

func init() {
	urlMap = make(map[string]string)
}

func shorturl(c *gin.Context) {
	longUrl := c.Query("url")
	if longUrl == "" {
		c.String(http.StatusBadRequest, "URL is necessary")
		return
	}
	shortUrl := makeShortUrl()
	urlMap[shortUrl] = longUrl
	c.Header("Location", "http://localhost:8080/"+shortUrl)
	c.String(http.StatusCreated, "http://localhost:8080/%s\n", shortUrl)
}

func makeShortUrl() string {
	const symbol = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
	shortUrl := make([]byte, 4)
	for i := range shortUrl {
		shortUrl[i] = symbol[rand.Intn(len(symbol))]
	}
	return string(shortUrl)
}

func redirectHandler(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	longUrl, ok := urlMap[shortUrl]
	if !ok {
		c.String(http.StatusNotFound, "URL is missing")
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, longUrl)
}

func main() {
	server := gin.Default()
	server.GET("/shorten", shorturl)
	server.GET("/:shortUrl", redirectHandler)
	server.Run(":8080")
}
