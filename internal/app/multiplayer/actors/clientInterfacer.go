package actors

import (
	"github.com/lesta-battleship/matchmaking/pkg/packets"
)

type ClientInterfacer interface {
	Id() string
	ConnectTo(Actor)
	GetPacket(senderId string, packet packets.Packet)
	ReadPump()
	WritePump()
	Stop()
}
