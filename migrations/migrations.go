package main

import (
	"thelist/inits"
	"thelist/models"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	inits.DB.AutoMigrate(&models.Entry{})
	inits.DB.AutoMigrate(&models.User{})
}

