package main

import (
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/melkecelioglu/e-commerce/controllers"
	"github.com/melkecelioglu/e-commerce/database"
	"github.com/melkecelioglu/e-commerce/middleware"
	"github.com/melkecelioglu/e-commerce/routes"
)



func main(){



	port:= os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client,"Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	routes.UserRoutes(router)


	router.Use(gin.Logger())
	router.Use(middleware.Authentication())
	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/listcart", controllers.GetItemFromCart())
	router.POST("/addaddress", controllers.AddAddress())
	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.PUT("/editworkaddress", controllers.EditWorkAddress())
	router.GET("/deleteaddresses", controllers.DeleteAddress())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())
	log.Fatal(router.Run(":" + port))


}