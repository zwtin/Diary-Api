package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
)

type Person struct {
	gorm.Model
	Name string
	Age  int
}

func db_init() {
	DBMS := "mysql"
	USER := "root"
	PASS := "hoge"
	PROTOCOL := ""
	DBNAME := "bookshelf"
	OPTION := "charset=utf8&parseTime=True&loc=Local"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	db.LogMode(true)
	db.AutoMigrate(&Person{})
}

func create(name string, age int) {
	DBMS := "mysql"
	USER := "root"
	PASS := "hoge"
	PROTOCOL := ""
	DBNAME := "bookshelf"
	OPTION := "charset=utf8&parseTime=True&loc=Local"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic("failed to connect database\n")
	}
	db.Create(&Person{Name: name, Age: age})
}

func get_all() []Person {
	DBMS := "mysql"
	USER := "root"
	PASS := "hoge"
	PROTOCOL := ""
	DBNAME := "bookshelf"
	OPTION := "charset=utf8&parseTime=True&loc=Local"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic("failed to connect database\n")
	}
	var people []Person
	db.Find(&people)
	return people
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	db_init()
	r.GET("/", func(c *gin.Context) {
		people := get_all()
		c.HTML(200, "index.tmpl", gin.H{
			"people": people,
		})
	})
	r.POST("/new", func(c *gin.Context) {
		name := c.PostForm("name")
		age, _ := strconv.Atoi(c.PostForm("age"))
		create(name, age)
		c.Redirect(302, "/")
	})
	r.Run()
}
