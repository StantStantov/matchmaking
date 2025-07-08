package websocket

import (
	"github.com/lesta-battleship/matchmaking/internal/app/multiplayer"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router gin.IRouter, websocketServer *WebsocketServer, engine *multiplayer.Engine) {
	randomHandler := HandleGetJoinRandom(websocketServer, engine)
	rankedHandler := HandleGetJoinRanked(websocketServer, engine)
	customHandler := HandleGetJoinCustom(websocketServer, engine)

	router.GET("/matchmaking/random", randomHandler)
	router.GET("/matchmaking/ranked", rankedHandler)
	router.GET("/matchmaking/custom", customHandler)
}
