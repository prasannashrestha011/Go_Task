package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
	IsActive bool `json:"isActive"`
}
func main() {
	_=godotenv.Load()

	r := gin.Default()
	
	r.POST("/user",func(ctx *gin.Context){
		var user User 
		if err:=ctx.BindJSON(&user);err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"details":"Invalid input field",
			})
		}
		ctx.JSON(http.StatusCreated,gin.H{
			"details":"User created",
			"user":user,
		}) 
	})

	PORT:=os.Getenv("PORT")
	fmt.Println("PORT: ",PORT)
	r.Run(":"+PORT)

}