package routes

import(
	"github.com/Parva-Parmar/GO-ecom/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("user/signup", controllers.Signup())
	incomingRoutes.POST("user/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct",controllers.ProductViewerAdmin())
	incomingRoutes.POST("/users/productview",controllers.SearchProduct())
	incomingRoutes.POST("/user/search", controllers.SearchProductByQuery())
}