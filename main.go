package main

import (
	"fmt"
	"net/http"
	"pustaka-api/controllers"
	"pustaka-api/models"
	"path/filepath"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//connect database
	db := models.SetupModels()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.POST("/upload", uploadFile)
	router.POST("/uploads", uploadFiles)
	router.GET("/mahasiswa", controllers.MahasiswaTampil)
	router.POST("/mahasiswa", controllers.MahasiswaTambah)
	router.PUT("/mahasiswa/:nim", controllers.MahasiswaUbah)
	router.DELETE("/mahasiswa/:nim", controllers.MahasiswaHapus)

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/hello/:id/:title", helloDetailHandler)
	router.GET("/query", queryHandler)

	router.POST("/", bookHandler)

	router.Run(":8888")
}

func uploadFile(c *gin.Context){
	file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		//filename := filepath.Base(file.Filename)
		filename := filepath.Join("images", file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully", file.Filename))
}
func uploadFiles(c *gin.Context){

		form, _ := c.MultipartForm()
		files := form.File["files[]"]
		filename := "." 
		for _, file := range files {
			
			filename = filepath.Join("images", file.Filename)
			// Upload the file to specific dst.
			c.SaveUploadedFile(file, filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
// func mahasiswaHandler(c *gin.Context) {
// 	rows, err := mdb.Query("select * from mahasiswa;")
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	nid := 0
// 	nnim := ' '
// 	nname := ' '
// 	for rows.Next() {
// 		err = rows.Scan(nid, nnim, nname)
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"name": nname,
// 		"bio":  nnim,
// 	})
// }

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Sarifudin Hidayat",
		"bio":  "A Software Engenering",
	})
}
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "Software engenring",
		"subtitle": "Golang",
	})
}

func helloDetailHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"Name": title,
		"age":  "30",
	})
}
func queryHandler(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"Name": name,
		"age":  "30",
	})
}

type BookInput struct {
	Title    string `json:"title" binding:"required"`
	Price    int    `json:"price" binding:"required,number"`
	SubTitle string `json:"sub_title"`
}

func bookHandler(c *gin.Context) {
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		errormessages := []string{}
		//for _, n := range err.(validator.ValidationErrors) {
		//errormessage := fmt.Sprintf("Error on Field %s, Condition %s", n.Filed(), n.ActualTag())
		//errormessages = append(errormessages errormessage)

		//}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errormessages,
		})
		//fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title":     bookInput.Title,
		"price":     bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})
}
