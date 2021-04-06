package main

import (
	"crowfunding/handler"
	"crowfunding/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/dbcrowfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})	

	if err != nil{ 
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	input := user.LoginUserInput{
		Email: "testsa1@gmail.com",
		Password: "qwerty123s",
	}
	user, err := userService.Login(input)

	if err != nil {
		fmt.Println("terjadi kesalahan")
		fmt.Println(err.Error())
		return
	}

	fmt.Println(user.Email)
	fmt.Println(user.Name)
	
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run(":8090")
}