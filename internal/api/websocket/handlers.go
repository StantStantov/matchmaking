package websocket

import (
	"github.com/lesta-battleship/matchmaking/internal/app/multiplayer"
	"github.com/lesta-battleship/matchmaking/internal/app/multiplayer/actors/matchmakers"

	"github.com/gin-gonic/gin"
)

func HandleGetJoinRandom(websocketServer *WebsocketServer, engine *multiplayer.Engine) gin.HandlerFunc {
	return gin.HandlerFunc(
		func(c *gin.Context) {
			interfacer, err := websocketServer.Connect(c)
			if err != nil {
				c.Error(err)

				return
			}
			go interfacer.ReadPump()
			go interfacer.WritePump()

			player := engine.CreatePlayer(interfacer)
			engine.SendToMatchmaking(player, matchmakers.RandomMatch)
		},
	)
}

func HandleGetJoinRanked(websocketServer *WebsocketServer, engine *multiplayer.Engine) gin.HandlerFunc {
	return gin.HandlerFunc(
		func(c *gin.Context) {
			interfacer, err := websocketServer.Connect(c)
			if err != nil {
				c.Error(err)

				return
			}
			go interfacer.ReadPump()
			go interfacer.WritePump()

			player := engine.CreatePlayer(interfacer)
			engine.SendToMatchmaking(player, matchmakers.RankedMatch)
		},
	)
}

func HandleGetJoinCustom(websocketServer *WebsocketServer, engine *multiplayer.Engine) gin.HandlerFunc {
	return gin.HandlerFunc(
		func(c *gin.Context) {
			interfacer, err := websocketServer.Connect(c)
			if err != nil {
				c.Error(err)

				return
			}
			go interfacer.ReadPump()
			go interfacer.WritePump()

			player := engine.CreatePlayer(interfacer)
			engine.SendToMatchmaking(player, matchmakers.CustomMatch)
		},
	)
}
