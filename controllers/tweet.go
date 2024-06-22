package controllers

import (
	rabbitmq "TweetDelivery/messagequeue/RabbitMQ"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TweetController struct{}

func (tc *TweetController) Get(c *gin.Context) {
	currentUser := c.Param("user")

	tweetsHome := rabbitmq.ConsumeTweet(currentUser)

	c.IndentedJSON(http.StatusOK, tweetsHome)
}

func (tc *TweetController) Post(c *gin.Context) {
	currentUser := c.Param("user")

	var tweet string

	if err := c.BindJSON(&tweet); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := rabbitmq.PublishTweet(currentUser, tweet)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tweet posted successfully",
	})
}
