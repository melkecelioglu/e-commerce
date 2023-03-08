package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
	"github.com/melkecelioglu/e-commerce/database"
	"github.com/melkecelioglu/e-commerce/models"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-conic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}


func NewApplication(prodCollection, userCollection *mongo.Collection) *Application{
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}


func (app *Application) AddToCart() gin.HandlerFunc{

}


func(app *Application) RemoveItem() gin.HandlerFunc{

}

func GetItemFromCart() gin.HandlerFunc{

}


func (app *Application) BuyFromCart() gin.HandlerFunc{


}
func (app *Application) InstantBuy() gin.HandlerFunc{
	
}