package main

import (
	"go-web-frame/gee/origin"
	"net/http"
)

func main() {
	r := origin.New()
	r.GET("/", func(c *origin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *origin.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *origin.Context) {
		c.JSON(http.StatusOK, origin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
