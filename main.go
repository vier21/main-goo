package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Booklist struct {
	judulBuku string `json: "judul"`
	penulis   string `json: "penulis"`
}

var ListBuku []Booklist

func main() {
	engine := gin.Default()

	bookRoutes := engine.Group("/buku")
	{
		bookRoutes.GET("/", GetBuku)
		bookRoutes.POST("/", CreateBuku)
		bookRoutes.PUT("/", UpdateBuku)
		bookRoutes.PUT("/:judul", DeleteBuku)
	}
	if err := engine.Run(); err != nil {
		log.Fatal(err.Error())
	}

}

func GetBuku(c *gin.Context) {
	c.JSON(200, ListBuku)

}

func CreateBuku(c *gin.Context) {
	var body Booklist
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"massage": "Request gagal",
		})
		return
	}
	ListBuku = append(ListBuku, body)

	c.JSON(200, gin.H{
		"error": false,
	})

}

func UpdateBuku(c *gin.Context) {
	var body Booklist
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"massage": "Request gagal",
		})
		return
	}
	for i, _ := range ListBuku {
		ListBuku[i].judulBuku = body.judulBuku
		ListBuku[i].penulis = body.penulis

		c.JSON(200, gin.H{
			"error": false,
		})
		return
	}
}

func DeleteBuku(c *gin.Context) {
	judul := c.Param("judul")
	for i, u := range ListBuku {
		if u.judulBuku == judul {
			ListBuku = append(ListBuku[:i], ListBuku[i+1:]...)

			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}
}
