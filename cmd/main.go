package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lesta-battleship/matchmaking/internal/api/websocket"
	"github.com/lesta-battleship/matchmaking/internal/app/multiplayer"
	"github.com/lesta-battleship/matchmaking/internal/app/multiplayer/actors/matchmakers"
	"github.com/lesta-battleship/matchmaking/internal/infra/registries"

	"github.com/gin-gonic/gin"
)

const backendPortEnv = "BACKEND_PORT"

func main() {
	matchmakerRegistry := registries.NewMatchmakerRegistry()
	roomRegistry := registries.NewRoomRegistry()
	playerRegistry := registries.NewPlayerRegistry()

	engine := multiplayer.NewEngine(matchmakerRegistry, roomRegistry, playerRegistry)

	engine.CreateHub()

	engine.CreateMatchmaker(matchmakers.RandomMatch)
	engine.CreateMatchmaker(matchmakers.RankedMatch)
	engine.CreateMatchmaker(matchmakers.CustomMatch)

	websocketServer := websocket.NewWebsocketServer()

	router := gin.Default()

	websocket.SetupRouter(router, websocketServer, engine)

	port, ok := os.LookupEnv(backendPortEnv)
	if !ok {
		log.Printf("Main: ENV %q is not defined", backendPortEnv)

		return
	}
	router.Run(fmt.Sprintf(":%s", port))
}
