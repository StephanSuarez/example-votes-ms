package http

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, deps *AppDependencies) {
	routesVotes := router.Group("/votes")
	routesVotes.GET("/", GetVotes)
	routesVotes.POST("/", deps.VoteHandler.CreateVote)
}