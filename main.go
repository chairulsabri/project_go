package main

import (
	"fmt"
	"log"
	"net/http"
	"startup/auth"
	"startup/campaign"
	"startup/handler"
	"startup/helper"
	"startup/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)

	campaignRepository := campaign.NewRepository(db)

	campaigns, err := campaignRepository.FindByUserID(1)

	fmt.Println("debug")
	fmt.Println("debug")
	fmt.Println("debug")
	fmt.Println(len(campaigns))
	for _, campaign := range campaigns {
		fmt.Println(campaign.Name)
		if len(campaign.CampaignImages) > 0 {
			println("jumlah gambar")
			println(len(campaign.CampaignImages))
			fmt.Println(campaign.CampaignImages[0].FileName)
		}
	}

	userService := user.NewService(userRepository)
	authService := auth.NewService()

	// token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyfQ.wmCsHDmiH1Q5WjoaEODpsaJ8bzzelXL366HD-L8mzYw")
	// if err != nil {
	// 	fmt.Println("ERROR")
	// 	fmt.Println("ERROR")
	// 	fmt.Println("ERROR")
	// }
	// if token.Valid {
	// 	fmt.Println("VALID")
	// 	fmt.Println("VALID")
	// 	fmt.Println("VALID")
	// } else {
	// 	fmt.Println("INVALID")
	// 	fmt.Println("INVALID")
	// 	fmt.Println("INVALID")
	// }
	// fmt.Println(authService.GenerateToken(1001))

	// userService.SaveAvatar(4, "images/1-profile.png")

	// input := user.LoginInput{
	// 	Email:    "chairulsabri@gmail.com",
	// 	Password: "rahasia",
	// }
	// user, err := userService.Login(input)
	// if err != nil {
	// 	fmt.Println("Terjadi kesalahan")
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(user.Email)
	// fmt.Println(user.Name)

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()

}

//input dari user
//handler mapping input dari usr -> struct input
//service melakukan MAPING DARI STRUCT
//repository CEKLIS
//db CEKLIS

//Midleware
// Ambil nilai header Authorization: Bearer tokentoken
// dari header Authorizaton, kita ambil nilai tikennya saja
// kita validasi token
// user ambil user_id
//ambil user dari db berdasarkan user_id lewat service
// kita set context isinya user

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//Bearer token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}
