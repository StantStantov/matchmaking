package strategies

import (
	"log"

	"github.com/lesta-battleship/matchmaking/internal/app/multiplayer/actors"
	"github.com/lesta-battleship/matchmaking/pkg/packets"
)

type InHub struct {
	Player actors.Actor
	Hub    actors.Actor
}

func (s *InHub) HandlePacket(senderId string, packet packets.Packet) {
	switch packet := packet.Body.(type) {
	case *packets.Disconnect:
		s.handleDisconnect(senderId, packet)
	default:
		log.Printf("Player %q: Received incorrect packet %T from %q", s.Player.Id(), packet, senderId)
	}
}

func (s *InHub) handleDisconnect(senderId string, packet *packets.Disconnect) {
	s.Hub.GetPacket(senderId, packets.NewDisconnect(senderId))
}

func (s *InHub) OnExit() {
	s.Hub = nil
}

func (s InHub) String() string {
	return "InHub"
}
