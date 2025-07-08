package websocket

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lesta-battleship/matchmaking/internal/api/websocket/middlewares"
	"github.com/lesta-battleship/matchmaking/internal/app/multiplayer/actors"
	"github.com/lesta-battleship/matchmaking/internal/infra"

	"github.com/gorilla/websocket"
)

type WebsocketServer struct {
	upgrader websocket.Upgrader
}

func NewWebsocketServer() *WebsocketServer {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	return &WebsocketServer{
		upgrader: upgrader,
	}
}

func (s *WebsocketServer) Connect(c *gin.Context) (actors.ClientInterfacer, error) {
	userId := middlewares.GetUserId(c)
	if userId == "" {
		return nil, fmt.Errorf("User ID is Empty")
	}

	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return nil, err
	}

	websocketInterfacer := infra.NewWebsocketInterfacer(userId, conn)

	log.Printf("WebsocketServer: Created WebsocketInterfacer %q", websocketInterfacer.Id())

	return websocketInterfacer, nil
}
