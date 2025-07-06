package websocket

import (
	"log"
	"net/http"

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

func (s *WebsocketServer) Connect(w http.ResponseWriter, r *http.Request) (actors.ClientInterfacer, error) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	connId := r.Header.Get("X-XSRF-TOKEN")
	if connId == "" {
		log.Println("WebsocketServer: Connection ID is empty")

		connId = infra.GenerateId()
	}

	websocketInterfacer := infra.NewWebsocketInterfacer(connId, conn)

	return websocketInterfacer, nil
}
