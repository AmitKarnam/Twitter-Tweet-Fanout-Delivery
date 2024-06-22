package server

import (
	"log"

	"TweetDelivery/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() (*gin.Engine, error) {

	gin.SetMode(gin.ReleaseMode)

	// Initialise New gin engine
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("%v %v %v %v", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	healthController := controllers.HealthController{}
	router.GET("/health", healthController.Status)

	service := router.Group("twitterdelivery")
	{
		apiGroup := service.Group("api")
		{
			versionGroup := apiGroup.Group("v1")
			{
				usersGroup := versionGroup.Group("users")
				userController := controllers.UserController{}
				usersGroup.GET("", userController.Get)
				usersGroup.POST("", userController.Post)

				log.Println("User Cotroller Initialized")
			}

			{
				followGroup := versionGroup.Group("follow")
				followController := controllers.FollowController{}
				followGroup.POST(":user", followController.Post)
			}

			{
				tweetGroup := versionGroup.Group("tweet")
				tweetController := controllers.TweetController{}
				tweetGroup.GET(":user", tweetController.Get)
				tweetGroup.POST(":user", tweetController.Post)
			}
		}
	}

	return router, nil
}
