package controllers

import(
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/melkecelioglu/e-commerce/database"
	"github.com/melkecelioglu/e-commerce/models"
	generate "github.com/melkecelioglu/e-commerce/tokens"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"


)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")
var ProductCollection *mongo.Collection = database.ProductData(database.Client, "Products")
var Validate = validator.New()

func HashPassword(password string) string{

}


func VerifyPassword(userpassword string, givenpassword string)(bool, string){

}

func SignUp() gin.HandlerFunc{

}

func Login() gin.HandlerFunc{

}

func ProductViewerAdmin() gin.HandlerFunc{

}

func SearchProduct() gin.HandlerFunc{

}

func SearchProductByQuery() gin.HandlerFunc{
	
}