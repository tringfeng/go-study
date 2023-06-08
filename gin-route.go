package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "hello gin")
    })
    r.POST("/post", postFn)
    r.PUT("/put", func(c *gin.Context) {
        
    })
   
    r.Run(":8787")  // 监听端口默认为8080
}



func postFn(c *gin.Context){
    types := c.DefaultPostForm("type", "post")
    username := c.PostForm("username")
    password := c.PostForm("userpassword")
    c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))

}
