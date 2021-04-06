package main

import (
	"crowfunding/handler"
	"crowfunding/user"
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
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run(":8090")

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Test dari user input"
	// userInput.Occupation = "Pebisnis"
	// userInput.Email = "test1@gmail.com"
	// userInput.Password = "qwerty123"

	// userService.RegisterUser(userInput)
	
	// user := user.User{
	// 	Name: "Ahmad",
	// 	Occupation: "Singer",
	// 	Email: "ahmad@gmail.com",
	// 	Password: "qwerty123",
	// 	Role: "user",
	// }

	// userRepository.Save(user)

}