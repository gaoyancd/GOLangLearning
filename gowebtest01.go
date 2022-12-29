package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func _json(c *gin.Context) {
	type user struct {
		Username string `json:"username1"`
		Userid   string `json:"userid"`
		Age      int    `json:"age"`
	}
	user1 := user{Username: "gaoyan", Userid: "1001", Age: 23}
	fmt.Println(user1)
	c.JSON(200, user1)
}

func _jsondirect(c *gin.Context) {

	c.JSON(200, gin.H{"username": "gaoyan1"})
}

func _jsonmap(c *gin.Context) {
	type usermap = map[string]string
	user2 := usermap{
		"username": "gaoyan",
		"userid":   "1001",
	}
	c.JSON(http.StatusOK, user2)
}

func _string(c *gin.Context) {
	c.String(http.StatusOK, "Hi GO WEB!")
}

func _redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
}

func main() {
	router := gin.Default()
	//golang中没有相对文件的路径，只有相对项目的路径
	//router.StaticFile("/static/food.png", "./static/food.png")
	router.StaticFS("/static", http.Dir("./static"))
	router.StaticFile("/food.png", "./static/food.png")
	router.GET("/json", _json)
	router.GET("/string", _string)
	router.GET("/jsonmap", _jsonmap)
	router.GET("/jsondirect", _jsondirect)
	router.GET("/baidu", _redirect)
	router.Run(":8080")
}
