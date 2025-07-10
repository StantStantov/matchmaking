package websocket

import (
	"github.com/lesta-battleship/matchmaking/internal/api/websocket/middlewares"
	"github.com/lesta-battleship/matchmaking/internal/app/multiplayer"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router gin.IRouter, websocketServer *WebsocketServer, engine *multiplayer.Engine) {
	handleError := middlewares.HandleErrorMiddleware()
	requireToken := middlewares.CheckTokenMiddleware()

	randomHandler := HandleGetJoinRandom(websocketServer, engine)
	rankedHandler := HandleGetJoinRanked(websocketServer, engine)
	customHandler := HandleGetJoinCustom(websocketServer, engine)

	router.Use(handleError, requireToken)
	router.GET("/matchmaking/random", randomHandler)
	router.GET("/matchmaking/ranked", rankedHandler)
	router.GET("/matchmaking/custom", customHandler)
}
