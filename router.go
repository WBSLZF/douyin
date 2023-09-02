package main

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/docs"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func initRouter(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", middleware.JWTMiddleWare, controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", middleware.JWTMiddleWare, controller.Publish)
	apiRouter.GET("/publish/list/", middleware.JWTMiddleWare, controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", middleware.JWTMiddleWare, controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", middleware.JWTMiddleWare, controller.FavoriteList)

	apiRouter.POST("/comment/action/", middleware.JWTMiddleWare, controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", middleware.JWTMiddleWare, controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", middleware.JWTMiddleWare, controller.FollowList)
	apiRouter.GET("/relation/follower/list/", middleware.JWTMiddleWare, controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", middleware.JWTMiddleWare, controller.FriendList)
	apiRouter.GET("/message/chat/", middleware.JWTMiddleWare, controller.MessageChat)
	apiRouter.POST("/message/action/", middleware.JWTMiddleWare, controller.MessageAction)
}
