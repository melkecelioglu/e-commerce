package routes


import (
	"github.com/melkecelioglu/e-commerce/controllers"
	"github.com/gin-gonic/gin"
)


func UserRoutes(incomingRroutes *gin.Engine){
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login",controllers.Login())
	incomingRoutes.POST("/admin/addproduct",controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())

}