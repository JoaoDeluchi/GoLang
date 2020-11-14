package main 

import (
	"github.com/gin-gonic/gin" 
)

func GetProduct(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Get Run well",
	})
}

func CreateProduct(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Post Run ok",
	})
}

func UpdateProduct(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Put ok",
	})
}

func DeleteProduct(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Delete OK !!!"
	})
}

func main() {
	r := gin.Default()
	r.GET("/product", GetProduct)
	r.POST("/product", CreateProduct)
	r.PUT("/product", UpdateProduct)
	r.DELETE("/product", DeleteProduct)
	r.Run()
}
