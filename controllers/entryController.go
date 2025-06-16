package controllers

import (
	"thelist/inits"
	"thelist/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// TODO: update all crud endpoints with appropriate error handling and response codes

func CreateEntry(ctx *gin.Context) {
	var body struct {
		Fname 	string
		Lname 	string
		State 	string
		Phone 	string
		UserId  uint `json:"user_id"`
	}

	ctx.BindJSON(&body)

	user, exists := ctx.Get("user")


	if !exists {
		ctx.JSON(50, gin.H{"error": "user not found"})
		return
	}

	body.UserId = user.(models.User).ID

	entry := models.Entry{
		Fname: 	body.Fname,
		Lname: 	body.Lname,
		State: 	body.State,
		Phone: 	body.Phone,
	}

	fmt.Println(entry)
	result := inits.DB.Create(&entry)
	fmt.Printf("result.Error: %v, %T", result.Error, result.Error)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": entry})
}


func GetEntries(ctx *gin.Context) {
	var entries []models.Entry
	
	result := inits.DB.Find(&entries)
	fmt.Printf("result.Error: %v, %T", result.Error, result.Error)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": entries})
}

func GetEntry(ctx *gin.Context) {
	id := ctx.Param("entryId")
	var entry models.Entry

	result := inits.DB.First(&entry, id)
	fmt.Printf("result.Error: %v, %T", result.Error, result.Error)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": entry})
}

func UpdateEntry(ctx *gin.Context) {
	var body struct {
		Fname string
		Lname string
		State string
		Phone string
	}

	ctx.BindJSON(&body)

	var entry models.Entry

	result := inits.DB.First(&entry, ctx.Param("entryId"))
	fmt.Printf("result.Error: %v, %T", result.Error, result.Error)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	inits.DB.Model(&entry).Updates(models.Entry{Fname: body.Fname, Lname: body.Lname, State: body.State, Phone: body.Phone})

	ctx.JSON(200, gin.H{"data": entry})
}

func DeleteEntry(ctx *gin.Context) {
	id := ctx.Param("entryId")

	result := inits.DB.Delete(&models.Entry{}, id)
	fmt.Printf("result.Error: %v, %T", result.Error, result.Error)

	ctx.JSON(200, gin.H{"data": "entry has been deleted successfully"})
}
