package main

import (
  "github.com/gin-gonic/gin"
  "./db"
  _ "github.com/mattn/go-sqlite3"
  _ "net/http"
)


func main() {
  r := gin.Default()
  r.LoadHTMLGlob("views/*")

  db.Init()

  r.GET("/", func(c *gin.Context) {
    links := db.GetAll()
    c.HTML(200, "index.tmpl", gin.H{
      "links": links,
    })
  })

  r.POST("/new", func(c *gin.Context) {
    name := c.PostForm("name")
    path := c.PostForm("path")
    db.Create(name, path)
    c.Redirect(302, "/")
  })

  r.Run()
}
