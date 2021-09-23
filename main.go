package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// set mode
	gin.SetMode(gin.ReleaseMode)
	// menginitialisasi untuk menghasilkan sebuah object router dan didalamnya ada juga middleware (logger)
	r := gin.Default()
	// menginitialisasi tanpa ada middleware apapun
	// r := gin.New()
	// menggunakan middleware logger
	// r.Use(gin.Logger())
	// mengembalikan sebuah response string "hello"
	// menggunakan http method apa ?
	// call function handler
	r.GET("/hello/:name", helloWithParam)
	// r.GET("/hello", hello)
	// with parameter in path
	r.GET("/hello/:name/:age", helloWithParam)
	// Query string in path
	r.GET("/hello", hello)
	r.GET("/list", listBook)
	r.GET("/login", login)

	// path parameter "hello" -> / :name
	// run
	// port 8080
	r.Run(":8089")
}

type Login struct {
	Username string `binding:"required" json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func login(c *gin.Context) {
	// username := c.PostForm("username")
	login := Login{}

	// Parse dari json atau form ke sebuah struct
	// form
	err := c.ShouldBindJSON(&login)

	if err != nil {
		c.AbortWithError(400, err).SetType(gin.ErrorTypeBind)
	}

	if login.Username == "root" && login.Password == "123" {
		c.JSON(200, gin.H{"code": 200, "message": "success"})
	}
	// shouldBindJSON()
}

type Book struct {
	Name string
	Page int
}

// For custom Response
type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func listBook(c *gin.Context) {
	var books []Book
	book1 := Book{Name: "bobo", Page: 10}
	books = append(books, book1)
	c.JSON(http.StatusOK, ResponseMessage(http.StatusOK, "Success", books))
	// page := c.DefaultQuery("page", "1")
	// order := c.DefaultQuery("order", "ASC")
	// mapQuery := c.Request
	// fmt.Println("page " + page + " order " + order)
}

// function untuk handle custom response
func ResponseMessage(code int, msg string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

func hello(c *gin.Context) {
	queryFirstname := c.DefaultQuery("firstname", "Guest")
	queryAge := c.Query("age")
	// mapQuery := c.Request.URL.Query()
	// fmt.Println(mapQuery)

	// parse data from string to int
	// age, _ := strconv.Atoi(queryAge)
	// fmt.Println(reflect.TypeOf(age))

	// fmt.Println(queryFirstname, queryAge)
	c.String(200, "My name is :"+queryFirstname+" my age : "+queryAge)
}

func helloWithParam(c *gin.Context) {
	name := c.Param("name")
	fmt.Println(name)
	c.String(201, "hello my name is : "+name)
}
