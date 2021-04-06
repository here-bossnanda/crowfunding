package main

import (
	"crowfunding/auth"
	"crowfunding/handler"
	"crowfunding/helper"
	"crowfunding/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checker", userHandler.CheckEmailAvailability)
	api.POST("/avatar",authMiddleware(authService, userService), userHandler.UploadAvatar)


	router.Run(":8090")
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
			return
		}
	
		var stringToken string
	
		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) == 2 {
			stringToken = splitToken[1]
		}
	
		token, err := authService.ValidateToken(stringToken)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
			return
		}

		payload, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
			return
		}

		userId := int(payload["userId"].(float64))

		user, err := userService.GetUserById(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
			return
		}

		c.Set("currentUser", user)
	}
}

