package controllers

import (
	"os"
	"fmt"
	"time"
	"net/http"

	"thelist/inits"
	"thelist/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var body struct {
		Username 	string `json:"username"`
		Email 		string `json:"email"`
		Password 	string `json:"password"`
	}

	if ctx.BindJSON(&body) != nil {
		fmt.Println(body)
		ctx.JSON(400, gin.H{"error": "bad request"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	user := models.User{Username: body.Username, Email: body.Email, Password: string(hash)}

	result := inits.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result .Error})
		return
	}

	ctx.JSON(200, gin.H{"data": user})
}

func Login(ctx *gin.Context) {
	var body struct {
		Email 		string
		Password 	string
	}

	if ctx.BindJSON(&body) != nil {
		ctx.JSON(400, gin.H{"error": "bad request"})
		return
	}

	var user models.User

	result := inits.DB.Where("email = $1", body.Email).First(&user)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30). Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "error signing token"})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "localhost", false, true)
}

func Validate(ctx *gin.Context) {
	user, err := ctx.Get("user")
	if err != false {
		ctx.JSON(500, gin.H{"error": err})
		return
	}
	fmt.Println(user)

	ctx.JSON(200, gin.H{"data": "You are logged in!", "user": user})
}

func GetUsers(ctx *gin.Context) {
	var users []models.User

	err := inits.DB.Model(&models.User{}).Preload("Entries").Find(&users).Error
	if err != nil {
		fmt.Println(err)
		ctx.JSON(500, gin.H{"error": "error getting users"})
		return
	}

	ctx.JSON(200, gin.H{"data": users})
}

func Logout(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", "", -1, "", "loclhost", false, true)
	ctx.JSON(200, gin.H{"data": "You have logged out successfully."})
}