package controllers

import (
	"github.com/erabxes/golang-jwt-project/database"
	"github.com/go-playground/validator/v10"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()

func VerifyPassword()

func SignUp()

func Login()

func GetUsers()

func GetUser()
