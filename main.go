package main

import (
	"CTF-backend/internal/api"
	"CTF-backend/internal/mysql"
	"CTF-backend/internal/utils"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
func main() {
	if level, err := logger.ParseLevel(utils.GetStringEnv("LOG_LEVEL", "debug")); err != nil {
		logger.Panic(err)
	} else {
		logger.SetLevel(level)
		logger.Info("Log Level Set")
	}
	// loading other components
	if err := mysql.InitDB(); err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(Cors())

	question := r.Group("/api/question")
	{
		question.GET("/list", api.GetQuestionList)
		question.GET("/recentList/:id", api.GetUserRecentRankData)
		//question.GET("/listByCategory", api.ProductionListByCategory)
		//question.GET("/info/:id", ProductHandler.ProductInfoHandler)
		//question.POST("/add", ProductHandler.AddProductHandler)
		//question.POST("/edit", ProductHandler.EditProductHandler)
		//question.POST("/delete/:id", ProductHandler.DeleteProductHandler)
	}
	user := r.Group("/api/user")
	{
		user.POST("/login", api.UserLogin)
		user.POST("register", api.UserRegister)
		//user.GET("/list", UserHandler.UserListHandler)
		//user.GET("/info/:id", api.UserInfoHandler)
		//user.POST("/add", UserHandler.AddUserHandler)
		user.POST("/edit", api.UserEdit)
		//user.POST("/delete/:id", UserHandler.DeleteUserHandler)
	}
	//banner := r.Group("/api/banner")
	//{
	//	banner.GET("/list", BannerHandler.BannerListHandler)
	//	banner.GET("/info/:id", BannerHandler.BannerInfoHandler)
	//	banner.POST("/add", BannerHandler.AddBannerHandler)
	//	banner.POST("/edit", BannerHandler.EditBannerHandler)
	//	banner.POST("/delete/:id", BannerHandler.DeleteBannerHandler)
	//}
	//
	//category := r.Group("/api/category")
	//{
	//	category.GET("/list", CategoryHandler.CategoryListHandler)
	//	category.GET("/list4backend", CategoryHandler.CategoryList4BackendHandler)
	//	category.GET("/info/:id", CategoryHandler.CategoryInfoHandler)
	//	category.POST("/add", CategoryHandler.AddCategoryHandler)
	//	category.POST("/edit", CategoryHandler.EditCategoryHandler)
	//	category.POST("/delete/:id", CategoryHandler.DeleteCategoryHandler)
	//}

	//order := r.Group("/api/order")
	//{
	//	order.GET("/list", OrderHandler.OrderListHandler)
	//	order.GET("/info/:id", OrderHandler.OrderInfoHandler)
	//	order.POST("/add", OrderHandler.AddOrderHandler)
	//	order.POST("/edit", OrderHandler.EditOrderHandler)
	//	order.POST("/delete/:id", OrderHandler.DeleteOrderHandler)
	//}

	port := utils.GetStringEnv("PORT", ":8080")
	if err := r.Run(port); err != nil {
		panic("run failed!")
	}
}
