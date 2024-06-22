package controllers

import (
	rabbitmq "TweetDelivery/messagequeue/RabbitMQ"
	"TweetDelivery/users"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FollowController struct{}

func (f *FollowController) Post(c *gin.Context) {
	currentUser := c.Param("user")

	var followUser users.User

	//TODO: Need to check is the user sent in the body is present or not

	if err := c.BindJSON(&followUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := rabbitmq.CreateBinding(currentUser, string(followUser))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s followed user %s successfully", currentUser, followUser)})
}
